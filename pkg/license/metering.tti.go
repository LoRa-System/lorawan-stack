// Copyright © 2019 The Things Industries B.V.

package license

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"go.thethings.network/lorawan-stack/pkg/ttipb"
	"go.thethings.network/lorawan-stack/pkg/ttnpb"
	"google.golang.org/grpc"
)

// Cluster is the interface used for getting metrics.
type Cluster interface {
	WithClusterAuth() grpc.CallOption
	GetPeerConn(ctx context.Context, role ttnpb.ClusterRole) (*grpc.ClientConn, error)
}

// MeteringReporter is the interface used for reporting metrics.
type MeteringReporter interface {
	Report(ctx context.Context, data *ttipb.MeteringData) error
}

type meteringSetup struct {
	config   *ttipb.MeteringConfiguration
	cluster  Cluster
	reporter MeteringReporter

	mu    sync.RWMutex
	apply func(ttipb.License) ttipb.License
}

// Apply updates the license according to the update rules.
func (s *meteringSetup) Apply(license ttipb.License) ttipb.License {
	if s == nil {
		return license
	}
	s.mu.RLock()
	if s.apply != nil {
		license = s.apply(license)
	}
	s.mu.RUnlock()
	return license
}

// CollectAndReport collects metrics from the cluster and reports them to the MeteringReporter.
func (s *meteringSetup) CollectAndReport(ctx context.Context) error {
	var meteringData ttipb.MeteringData

	cc, err := s.cluster.GetPeerConn(ctx, ttnpb.ClusterRole_ENTITY_REGISTRY)
	if err != nil {
		return err
	}

	reg := ttipb.NewTenantRegistryClient(cc)
	creds := s.cluster.WithClusterAuth()

	// TODO: List all tenants and range over those.
	// tenants, err := reg.List(ctx, &ttipb.ListTenantsRequest{
	// 	FieldMask: pbtypes.FieldMask{Paths: []string{"ids"}},
	// 	Limit:     0,
	// 	Page:      1, // TODO: Pagination
	// }, creds)
	// if err != nil {
	// 	return err
	// }

	totals, err := reg.GetRegistryTotals(ctx, &ttipb.GetTenantRegistryTotalsRequest{
		// TODO: TenantIdentifiers
	}, creds)
	if err != nil {
		return err
	}
	meteringData.Tenants = append(meteringData.Tenants, &ttipb.MeteringData_TenantMeteringData{
		// TODO: TenantIdentifiers
		Totals: totals,
	})

	if err = s.reporter.Report(ctx, &meteringData); err != nil {
		return err
	}

	if s.config.OnSuccess != nil {
		now := time.Now()
		s.mu.Lock()
		s.apply = func(license ttipb.License) ttipb.License {
			if s.config.OnSuccess.ExtendValidUntil != nil {
				license.ValidUntil = now.Add(*s.config.OnSuccess.ExtendValidUntil)
			}
			return license
		}
		s.mu.Unlock()
	}

	return nil
}

// Run the periodic metrics reporting.
func (s *meteringSetup) Run(ctx context.Context) error {
	interval := time.Hour
	if s.config.Interval != nil {
		interval = *s.config.Interval
	}
	reportTicker := time.NewTicker(interval)
	defer reportTicker.Stop()
	for {
		select {
		case <-ctx.Done():
			return s.CollectAndReport(ctx)
		case <-reportTicker.C:
			s.CollectAndReport(ctx)
		}
	}
}

var globalMetering *meteringSetup

// SetupMetering sets up metering on cluster.
func SetupMetering(ctx context.Context, config *ttipb.MeteringConfiguration, cluster Cluster) error {
	if globalMetering != nil {
		return errors.New("only one metering configuration can be set up")
	}
	globalMetering = &meteringSetup{
		config:  config,
		cluster: cluster,
	}
	switch reporterConfig := config.Metering.(type) {
	case *ttipb.MeteringConfiguration_Aws:
		// TODO: Set up AWS metering reporter.
		_ = reporterConfig
		globalMetering.reporter = nil
	default:
		panic(fmt.Errorf("unsupported metering reporter config type: %T", config.Metering))
	}

	if err := globalMetering.CollectAndReport(ctx); err != nil {
		return err
	}

	go globalMetering.Run(ctx)

	return nil
}
