// Copyright © 2018 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package unique_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/smartystreets/assertions"
	"go.thethings.network/lorawan-stack/pkg/ttnpb"
	. "go.thethings.network/lorawan-stack/pkg/unique"
	"go.thethings.network/lorawan-stack/pkg/util/test/assertions/should"
)

func TestRoundtrip(t *testing.T) {
	for _, tc := range []struct {
		ID       ttnpb.Identifiers
		Expected string
		Parser   func(string) (ttnpb.Identifiers, error)
	}{
		{
			nil,
			"",
			nil,
		},
		{
			(*ttnpb.ApplicationIdentifiers)(nil),
			"",
			nil,
		},
		{
			ttnpb.ApplicationIdentifiers{ApplicationID: "foo"},
			"foo",
			func(uid string) (ttnpb.Identifiers, error) { return ToApplicationID(uid) },
		},
		{
			&ttnpb.ApplicationIdentifiers{ApplicationID: "foo"},
			"foo",
			func(uid string) (ttnpb.Identifiers, error) { return ToApplicationID(uid) },
		},
		{
			(*ttnpb.ClientIdentifiers)(nil),
			"",
			nil,
		},
		{
			ttnpb.ClientIdentifiers{ClientID: "foo"},
			"foo",
			func(uid string) (ttnpb.Identifiers, error) { return ToClientID(uid) },
		},
		{
			&ttnpb.ClientIdentifiers{ClientID: "foo"},
			"foo",
			func(uid string) (ttnpb.Identifiers, error) { return ToClientID(uid) },
		},
		{
			(*ttnpb.EndDeviceIdentifiers)(nil),
			"",
			nil,
		},
		{
			ttnpb.EndDeviceIdentifiers{
				ApplicationIdentifiers: ttnpb.ApplicationIdentifiers{
					ApplicationID: "foo-app",
				},
				DeviceID: "foo-device",
			},
			"foo-app:foo-device",
			func(uid string) (ttnpb.Identifiers, error) { return ToDeviceID(uid) },
		},
		{
			&ttnpb.EndDeviceIdentifiers{
				ApplicationIdentifiers: ttnpb.ApplicationIdentifiers{
					ApplicationID: "foo-app",
				},
				DeviceID: "foo-device",
			},
			"foo-app:foo-device",
			func(uid string) (ttnpb.Identifiers, error) { return ToDeviceID(uid) },
		},
		{
			(*ttnpb.GatewayIdentifiers)(nil),
			"",
			nil,
		},
		{
			ttnpb.GatewayIdentifiers{GatewayID: "foo"},
			"foo",
			func(uid string) (ttnpb.Identifiers, error) { return ToGatewayID(uid) },
		},
		{
			&ttnpb.GatewayIdentifiers{GatewayID: "foo"},
			"foo",
			func(uid string) (ttnpb.Identifiers, error) { return ToGatewayID(uid) },
		},
		{
			(*ttnpb.OrganizationIdentifiers)(nil),
			"",
			nil,
		},
		{
			ttnpb.OrganizationIdentifiers{OrganizationID: "foo"},
			"foo",
			func(uid string) (ttnpb.Identifiers, error) { return ToOrganizationID(uid) },
		},
		{
			&ttnpb.OrganizationIdentifiers{OrganizationID: "foo"},
			"foo",
			func(uid string) (ttnpb.Identifiers, error) { return ToOrganizationID(uid) },
		},
		{
			(*ttnpb.UserIdentifiers)(nil),
			"",
			nil,
		},
		{
			ttnpb.UserIdentifiers{UserID: "foo"},
			"foo",
			func(uid string) (ttnpb.Identifiers, error) { return ToUserID(uid) },
		},
		{
			&ttnpb.UserIdentifiers{UserID: "foo"},
			"foo",
			func(uid string) (ttnpb.Identifiers, error) { return ToUserID(uid) },
		},
	} {
		t.Run(fmt.Sprintf("%T", tc.ID), func(t *testing.T) {
			a := assertions.New(t)
			a.So(ID(context.Background(), tc.ID), should.Equal, tc.Expected)
			if tc.Parser != nil {
				parsed, err := tc.Parser(tc.Expected)
				if a.So(err, should.BeNil) {
					a.So(parsed.CombinedIdentifiers(), should.Resemble, tc.ID.CombinedIdentifiers())
				}
			}
		})
	}
}
