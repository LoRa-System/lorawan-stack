// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/packetbrokeragent.proto

package ttnpb

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("lorawan-stack/api/packetbrokeragent.proto", fileDescriptor_1a44242dc5cd678e)
}
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/packetbrokeragent.proto", fileDescriptor_1a44242dc5cd678e)
}

var fileDescriptor_1a44242dc5cd678e = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0xd0, 0xa1, 0x6f, 0xea, 0x50,
	0x14, 0x06, 0xf0, 0x73, 0x04, 0x4f, 0x34, 0x79, 0xef, 0x25, 0x88, 0x09, 0x96, 0x9c, 0x4d, 0xcc,
	0x2c, 0xd9, 0x6e, 0x13, 0xf8, 0x0f, 0x96, 0x2d, 0x24, 0x4b, 0x58, 0x48, 0x96, 0x4d, 0xcc, 0xdd,
	0x92, 0xbb, 0xd2, 0xb4, 0xdc, 0xdb, 0xb4, 0x97, 0x35, 0x38, 0x24, 0x72, 0x72, 0x72, 0x99, 0x42,
	0x22, 0x91, 0x48, 0x24, 0x12, 0x49, 0xef, 0x35, 0x48, 0x24, 0x72, 0x59, 0x0b, 0x82, 0x4d, 0x4c,
	0x7f, 0x5f, 0x7e, 0xdf, 0xc9, 0x71, 0xce, 0x23, 0x95, 0xf0, 0x8c, 0xcb, 0xcb, 0x54, 0xf3, 0x4e,
	0xe8, 0xf2, 0x38, 0x70, 0x63, 0xde, 0x09, 0x85, 0xf6, 0x12, 0x15, 0x8a, 0x84, 0xfb, 0x42, 0x6a,
	0x16, 0x27, 0x4a, 0xab, 0xea, 0x3f, 0xad, 0x25, 0xdb, 0xd5, 0xd9, 0x4b, 0xa3, 0x76, 0xec, 0x2b,
	0xe5, 0x47, 0xc2, 0x2d, 0x52, 0xaf, 0xff, 0xec, 0x8a, 0x5e, 0xac, 0x07, 0x65, 0xb9, 0x76, 0xfa,
	0xd3, 0xed, 0x89, 0x34, 0xe5, 0xbe, 0x48, 0xcb, 0x46, 0xfd, 0xd1, 0xa9, 0x34, 0xd3, 0xb6, 0xc7,
	0xab, 0x2d, 0xe7, 0x6f, 0xbb, 0xef, 0x45, 0x41, 0xda, 0x7d, 0x88, 0xa3, 0x40, 0x86, 0xd5, 0x33,
	0x76, 0xb8, 0xc4, 0x9a, 0x5c, 0x8b, 0x8c, 0x0f, 0xca, 0xb8, 0x55, 0x32, 0xb5, 0x23, 0x56, 0xee,
	0xb3, 0xfd, 0x3e, 0xbb, 0xf9, 0xda, 0xaf, 0xdf, 0x3b, 0x95, 0xbb, 0xc2, 0xbd, 0x75, 0xfe, 0xef,
	0xdc, 0x6b, 0x95, 0xc9, 0x42, 0x3e, 0xf9, 0x2e, 0xef, 0x93, 0x5f, 0xd0, 0xab, 0x0f, 0x9c, 0xe7,
	0x84, 0x8b, 0x9c, 0x70, 0x99, 0x13, 0xac, 0x72, 0x82, 0x75, 0x4e, 0xb0, 0xc9, 0x09, 0xb6, 0x39,
	0xe1, 0xd0, 0x10, 0x8e, 0x0c, 0xc1, 0xd8, 0x10, 0x4e, 0x0c, 0xc1, 0xd4, 0x10, 0xcc, 0x0c, 0xc1,
	0xdc, 0x10, 0x2e, 0x0c, 0xe1, 0xd2, 0x10, 0xac, 0x0c, 0xe1, 0xda, 0x10, 0x6c, 0x0c, 0xe1, 0xd6,
	0x10, 0x0c, 0x2d, 0xc1, 0xc8, 0x12, 0xbe, 0x5a, 0x82, 0x37, 0x4b, 0xf8, 0x6e, 0x09, 0xc6, 0x96,
	0x60, 0x62, 0x09, 0xa7, 0x96, 0x70, 0x66, 0x09, 0x9f, 0x2e, 0x7c, 0xc5, 0x74, 0x57, 0xe8, 0x6e,
	0x20, 0xfd, 0x94, 0x49, 0xa1, 0x33, 0x95, 0x84, 0xee, 0xe1, 0x6b, 0xe3, 0xd0, 0x77, 0xb5, 0x96,
	0xb1, 0xe7, 0xfd, 0x29, 0x8e, 0x6e, 0x7c, 0x06, 0x00, 0x00, 0xff, 0xff, 0x1a, 0x20, 0xdb, 0xf7,
	0xd4, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GsPbaClient is the client API for GsPba service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GsPbaClient interface {
	PublishUplink(ctx context.Context, in *GatewayUplinkMessage, opts ...grpc.CallOption) (*types.Empty, error)
}

type gsPbaClient struct {
	cc *grpc.ClientConn
}

func NewGsPbaClient(cc *grpc.ClientConn) GsPbaClient {
	return &gsPbaClient{cc}
}

func (c *gsPbaClient) PublishUplink(ctx context.Context, in *GatewayUplinkMessage, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.GsPba/PublishUplink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GsPbaServer is the server API for GsPba service.
type GsPbaServer interface {
	PublishUplink(context.Context, *GatewayUplinkMessage) (*types.Empty, error)
}

// UnimplementedGsPbaServer can be embedded to have forward compatible implementations.
type UnimplementedGsPbaServer struct {
}

func (*UnimplementedGsPbaServer) PublishUplink(ctx context.Context, req *GatewayUplinkMessage) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishUplink not implemented")
}

func RegisterGsPbaServer(s *grpc.Server, srv GsPbaServer) {
	s.RegisterService(&_GsPba_serviceDesc, srv)
}

func _GsPba_PublishUplink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GatewayUplinkMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GsPbaServer).PublishUplink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.GsPba/PublishUplink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GsPbaServer).PublishUplink(ctx, req.(*GatewayUplinkMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _GsPba_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.GsPba",
	HandlerType: (*GsPbaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PublishUplink",
			Handler:    _GsPba_PublishUplink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/packetbrokeragent.proto",
}

// NsPbaClient is the client API for NsPba service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NsPbaClient interface {
	// PublishDownlink instructs the Packet Broker Agent to publish a downlink
	// message to Packet Broker Router.
	PublishDownlink(ctx context.Context, in *DownlinkMessage, opts ...grpc.CallOption) (*types.Empty, error)
}

type nsPbaClient struct {
	cc *grpc.ClientConn
}

func NewNsPbaClient(cc *grpc.ClientConn) NsPbaClient {
	return &nsPbaClient{cc}
}

func (c *nsPbaClient) PublishDownlink(ctx context.Context, in *DownlinkMessage, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.NsPba/PublishDownlink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NsPbaServer is the server API for NsPba service.
type NsPbaServer interface {
	// PublishDownlink instructs the Packet Broker Agent to publish a downlink
	// message to Packet Broker Router.
	PublishDownlink(context.Context, *DownlinkMessage) (*types.Empty, error)
}

// UnimplementedNsPbaServer can be embedded to have forward compatible implementations.
type UnimplementedNsPbaServer struct {
}

func (*UnimplementedNsPbaServer) PublishDownlink(ctx context.Context, req *DownlinkMessage) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishDownlink not implemented")
}

func RegisterNsPbaServer(s *grpc.Server, srv NsPbaServer) {
	s.RegisterService(&_NsPba_serviceDesc, srv)
}

func _NsPba_PublishDownlink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownlinkMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NsPbaServer).PublishDownlink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.NsPba/PublishDownlink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NsPbaServer).PublishDownlink(ctx, req.(*DownlinkMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _NsPba_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.NsPba",
	HandlerType: (*NsPbaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PublishDownlink",
			Handler:    _NsPba_PublishDownlink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/packetbrokeragent.proto",
}
