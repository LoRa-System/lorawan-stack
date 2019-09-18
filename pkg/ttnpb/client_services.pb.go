// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/client_services.proto

package ttnpb

import (
	context "context"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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
	proto.RegisterFile("lorawan-stack/api/client_services.proto", fileDescriptor_80815ba053239a77)
}
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/client_services.proto", fileDescriptor_80815ba053239a77)
}

var fileDescriptor_80815ba053239a77 = []byte{
	// 731 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x3f, 0x4c, 0x14, 0x4d,
	0x14, 0x9f, 0xe1, 0xfb, 0xbe, 0xcb, 0x97, 0x8d, 0x81, 0x38, 0x31, 0x98, 0x1c, 0xf0, 0x94, 0xc5,
	0x84, 0x84, 0xc0, 0xac, 0x81, 0x98, 0x18, 0x3b, 0x45, 0x45, 0xa3, 0x8d, 0x10, 0x9b, 0x6b, 0xc8,
	0xde, 0x31, 0xec, 0x4d, 0xee, 0xd8, 0x39, 0x76, 0xe6, 0x20, 0x48, 0x48, 0x88, 0x85, 0xd2, 0x98,
	0x48, 0x6c, 0xb4, 0x33, 0x56, 0x94, 0x94, 0x94, 0x94, 0x74, 0x92, 0xd0, 0x50, 0x72, 0xb3, 0x16,
	0x94, 0x74, 0x52, 0x9a, 0x9b, 0xdd, 0xc5, 0xbd, 0x7f, 0xc2, 0x25, 0x76, 0xb3, 0xf3, 0x7e, 0xf3,
	0x7b, 0xf3, 0x7e, 0xef, 0xf7, 0x66, 0xad, 0xd1, 0xb2, 0x08, 0xdc, 0x55, 0xd7, 0x9f, 0x90, 0xca,
	0x2d, 0x94, 0x1c, 0xb7, 0xc2, 0x9d, 0x42, 0x99, 0x33, 0x5f, 0xcd, 0x4b, 0x16, 0xac, 0xf0, 0x02,
	0x93, 0xb4, 0x12, 0x08, 0x25, 0x48, 0xaf, 0x52, 0x3e, 0x8d, 0xc1, 0x74, 0x65, 0x2a, 0x3b, 0xe1,
	0x71, 0x55, 0xac, 0xe6, 0x69, 0x41, 0x2c, 0x39, 0x9e, 0xf0, 0x84, 0x63, 0x60, 0xf9, 0xea, 0xa2,
	0xf9, 0x32, 0x1f, 0x66, 0x15, 0x1d, 0xcf, 0x0e, 0x7a, 0x42, 0x78, 0x65, 0x66, 0x12, 0xb8, 0xbe,
	0x2f, 0x94, 0xab, 0xb8, 0xf0, 0x63, 0xf2, 0xec, 0x40, 0x1c, 0xbd, 0xe0, 0x60, 0x4b, 0x15, 0xb5,
	0x16, 0x07, 0xa1, 0xd3, 0x15, 0xe3, 0xf8, 0x48, 0x6b, 0x9c, 0x2f, 0x30, 0x5f, 0xf1, 0x45, 0xce,
	0x02, 0xd9, 0x99, 0x24, 0xe0, 0x5e, 0x51, 0xc5, 0xf1, 0xc9, 0xf3, 0xff, 0xac, 0xde, 0x69, 0xc3,
	0x3a, 0xcb, 0x3c, 0x2e, 0x55, 0xb0, 0x46, 0xbe, 0x63, 0x2b, 0x33, 0x1d, 0x30, 0x57, 0x31, 0x32,
	0x42, 0x1b, 0xab, 0xa7, 0xd1, 0x7e, 0x72, 0x60, 0xb9, 0xca, 0xa4, 0xca, 0xf6, 0xb7, 0x80, 0x4c,
	0xd8, 0x7e, 0x8f, 0xdf, 0x1e, 0xfd, 0xf8, 0xd4, 0xb3, 0x89, 0x6d, 0xea, 0x54, 0x25, 0x0b, 0xa4,
	0xb3, 0x5e, 0x10, 0xe5, 0xb2, 0x9b, 0x17, 0x81, 0xab, 0x44, 0x40, 0xeb, 0x7b, 0xf3, 0x7c, 0x41,
	0x26, 0x8b, 0x8d, 0xb8, 0x3c, 0xf9, 0x00, 0x8f, 0xe5, 0x5e, 0xd8, 0x4f, 0x1d, 0x11, 0x78, 0xae,
	0xcf, 0xdf, 0x44, 0x8a, 0x35, 0x1d, 0x4e, 0xc7, 0x0c, 0x49, 0xd3, 0x46, 0x9a, 0x8c, 0x14, 0xad,
	0x7f, 0x66, 0x98, 0x22, 0xb7, 0x9b, 0x2f, 0x3a, 0xc3, 0xd4, 0xd5, 0x4a, 0x19, 0x35, 0x95, 0x0c,
	0x93, 0x5b, 0x09, 0xab, 0xb3, 0x1e, 0xbb, 0xa5, 0x9e, 0xfa, 0x62, 0xb9, 0x41, 0x8e, 0xb0, 0xf5,
	0xef, 0x4b, 0x2e, 0x15, 0xb1, 0x9b, 0x99, 0xea, 0xbb, 0x11, 0x9b, 0x4c, 0xb2, 0xdd, 0x6c, 0x9f,
	0x4d, 0xda, 0x1f, 0x22, 0xe5, 0xde, 0x61, 0xf2, 0x7f, 0x92, 0x30, 0x77, 0x97, 0x74, 0xa9, 0x62,
	0xee, 0x19, 0xf9, 0x4b, 0x12, 0x92, 0x65, 0x2b, 0xf3, 0xba, 0xb2, 0xd0, 0xd6, 0x10, 0xd1, 0xfe,
	0xd5, 0x54, 0x1c, 0x33, 0x55, 0xdd, 0xc9, 0xb6, 0xa8, 0x48, 0x1b, 0x55, 0xac, 0xb7, 0xcc, 0xb5,
	0x32, 0x8f, 0x59, 0x99, 0x29, 0x46, 0x86, 0xdb, 0xb3, 0x3d, 0xff, 0x6d, 0xf5, 0x6c, 0x3f, 0x8d,
	0xe6, 0x88, 0x26, 0x73, 0x44, 0x9f, 0xd4, 0xe7, 0xc8, 0x1e, 0x34, 0x09, 0xfb, 0xc7, 0x6e, 0xb4,
	0x69, 0xdb, 0xc6, 0xe4, 0x76, 0xc6, 0xba, 0x16, 0x71, 0x3d, 0x2c, 0x14, 0x98, 0x94, 0xa4, 0x6c,
	0x59, 0xf5, 0x2e, 0xcd, 0x9a, 0xf9, 0xb8, 0x5a, 0xde, 0x26, 0x48, 0x74, 0xd4, 0x1e, 0x31, 0x79,
	0x87, 0xc8, 0x40, 0xbb, 0xbc, 0xf1, 0xfc, 0x11, 0xdd, 0x63, 0xf5, 0xd5, 0x0d, 0x98, 0x6a, 0x09,
	0x19, 0xef, 0xe8, 0xd0, 0x34, 0x2c, 0xd1, 0x79, 0xb4, 0x1d, 0xba, 0x01, 0x27, 0x2b, 0xc2, 0x97,
	0xcc, 0xfe, 0x19, 0xf9, 0xe9, 0x0c, 0x93, 0xf1, 0x4b, 0x0c, 0xec, 0xa4, 0x1d, 0x92, 0x9b, 0x23,
	0xaf, 0xba, 0xc1, 0x1b, 0x7f, 0x5e, 0x66, 0xcf, 0x5c, 0x89, 0xf0, 0xae, 0x48, 0xd3, 0xae, 0xec,
	0xd6, 0xc1, 0x64, 0x1b, 0x5b, 0x7d, 0x73, 0x97, 0x89, 0x3c, 0xf7, 0x27, 0x91, 0x3b, 0x79, 0xeb,
	0xbe, 0x91, 0x74, 0x32, 0x3b, 0xd1, 0x4d, 0x31, 0xe6, 0x35, 0xfa, 0x82, 0xad, 0xeb, 0xe6, 0x35,
	0x48, 0x07, 0x08, 0xed, 0xfc, 0x60, 0x34, 0x00, 0x93, 0x7b, 0x0d, 0xb5, 0xd8, 0x33, 0x8d, 0xb2,
	0xef, 0x99, 0xeb, 0x39, 0xa4, 0xbb, 0xeb, 0x3d, 0xfa, 0x86, 0x0f, 0x6a, 0x80, 0x0f, 0x6b, 0x80,
	0x8f, 0x6b, 0x80, 0x4e, 0x6a, 0x80, 0x4e, 0x6b, 0x80, 0xce, 0x6a, 0x80, 0xce, 0x6b, 0x80, 0x37,
	0x35, 0xe0, 0x2d, 0x0d, 0x68, 0x47, 0x03, 0xde, 0xd5, 0x80, 0xf6, 0x34, 0xa0, 0x7d, 0x0d, 0xe8,
	0x40, 0x03, 0x3e, 0xd4, 0x80, 0x8f, 0x35, 0xa0, 0x13, 0x0d, 0xf8, 0x54, 0x03, 0x3a, 0xd3, 0x80,
	0xcf, 0x35, 0xa0, 0xcd, 0x10, 0xd0, 0x56, 0x08, 0xf8, 0x63, 0x08, 0xe8, 0x73, 0x08, 0xf8, 0x6b,
	0x08, 0x68, 0x27, 0x04, 0xb4, 0x1b, 0x02, 0xde, 0x0b, 0x01, 0xef, 0x87, 0x80, 0x73, 0xe3, 0x9e,
	0xa0, 0xaa, 0xc8, 0x54, 0x91, 0xfb, 0x9e, 0xa4, 0x3e, 0x53, 0xab, 0x22, 0x28, 0x39, 0x8d, 0xbf,
	0xae, 0x4a, 0xc9, 0x73, 0x94, 0xf2, 0x2b, 0xf9, 0x7c, 0xc6, 0xb4, 0x62, 0xea, 0x57, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xc8, 0x1c, 0x22, 0x73, 0xc4, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ClientRegistryClient is the client API for ClientRegistry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClientRegistryClient interface {
	// Create a new OAuth client. This also sets the given organization or user as
	// first collaborator with all possible rights.
	Create(ctx context.Context, in *CreateClientRequest, opts ...grpc.CallOption) (*Client, error)
	// Get the OAuth client with the given identifiers, selecting the fields given
	// by the field mask. The method may return more or less fields, depending on
	// the rights of the caller.
	Get(ctx context.Context, in *GetClientRequest, opts ...grpc.CallOption) (*Client, error)
	// List OAuth clients. See request message for details.
	List(ctx context.Context, in *ListClientsRequest, opts ...grpc.CallOption) (*Clients, error)
	Update(ctx context.Context, in *UpdateClientRequest, opts ...grpc.CallOption) (*Client, error)
	Delete(ctx context.Context, in *ClientIdentifiers, opts ...grpc.CallOption) (*types.Empty, error)
}

type clientRegistryClient struct {
	cc *grpc.ClientConn
}

func NewClientRegistryClient(cc *grpc.ClientConn) ClientRegistryClient {
	return &clientRegistryClient{cc}
}

func (c *clientRegistryClient) Create(ctx context.Context, in *CreateClientRequest, opts ...grpc.CallOption) (*Client, error) {
	out := new(Client)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ClientRegistry/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientRegistryClient) Get(ctx context.Context, in *GetClientRequest, opts ...grpc.CallOption) (*Client, error) {
	out := new(Client)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ClientRegistry/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientRegistryClient) List(ctx context.Context, in *ListClientsRequest, opts ...grpc.CallOption) (*Clients, error) {
	out := new(Clients)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ClientRegistry/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientRegistryClient) Update(ctx context.Context, in *UpdateClientRequest, opts ...grpc.CallOption) (*Client, error) {
	out := new(Client)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ClientRegistry/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientRegistryClient) Delete(ctx context.Context, in *ClientIdentifiers, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ClientRegistry/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientRegistryServer is the server API for ClientRegistry service.
type ClientRegistryServer interface {
	// Create a new OAuth client. This also sets the given organization or user as
	// first collaborator with all possible rights.
	Create(context.Context, *CreateClientRequest) (*Client, error)
	// Get the OAuth client with the given identifiers, selecting the fields given
	// by the field mask. The method may return more or less fields, depending on
	// the rights of the caller.
	Get(context.Context, *GetClientRequest) (*Client, error)
	// List OAuth clients. See request message for details.
	List(context.Context, *ListClientsRequest) (*Clients, error)
	Update(context.Context, *UpdateClientRequest) (*Client, error)
	Delete(context.Context, *ClientIdentifiers) (*types.Empty, error)
}

// UnimplementedClientRegistryServer can be embedded to have forward compatible implementations.
type UnimplementedClientRegistryServer struct {
}

func (*UnimplementedClientRegistryServer) Create(ctx context.Context, req *CreateClientRequest) (*Client, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedClientRegistryServer) Get(ctx context.Context, req *GetClientRequest) (*Client, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedClientRegistryServer) List(ctx context.Context, req *ListClientsRequest) (*Clients, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedClientRegistryServer) Update(ctx context.Context, req *UpdateClientRequest) (*Client, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedClientRegistryServer) Delete(ctx context.Context, req *ClientIdentifiers) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterClientRegistryServer(s *grpc.Server, srv ClientRegistryServer) {
	s.RegisterService(&_ClientRegistry_serviceDesc, srv)
}

func _ClientRegistry_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientRegistryServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ClientRegistry/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientRegistryServer).Create(ctx, req.(*CreateClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientRegistry_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientRegistryServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ClientRegistry/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientRegistryServer).Get(ctx, req.(*GetClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientRegistry_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClientsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientRegistryServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ClientRegistry/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientRegistryServer).List(ctx, req.(*ListClientsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientRegistry_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientRegistryServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ClientRegistry/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientRegistryServer).Update(ctx, req.(*UpdateClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientRegistry_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientRegistryServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ClientRegistry/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientRegistryServer).Delete(ctx, req.(*ClientIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClientRegistry_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.ClientRegistry",
	HandlerType: (*ClientRegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ClientRegistry_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ClientRegistry_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ClientRegistry_List_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ClientRegistry_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ClientRegistry_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/client_services.proto",
}

// ClientAccessClient is the client API for ClientAccess service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClientAccessClient interface {
	ListRights(ctx context.Context, in *ClientIdentifiers, opts ...grpc.CallOption) (*Rights, error)
	// Get the rights of a collaborator (member) of the client.
	// Pseudo-rights in the response (such as the "_ALL" right) are not expanded.
	GetCollaborator(ctx context.Context, in *GetClientCollaboratorRequest, opts ...grpc.CallOption) (*GetCollaboratorResponse, error)
	// Set the rights of a collaborator (member) on the client. It is required for the caller to
	// have all assigned or/and removed rights.
	// Setting a collaborator without rights, removes them.
	SetCollaborator(ctx context.Context, in *SetClientCollaboratorRequest, opts ...grpc.CallOption) (*types.Empty, error)
	ListCollaborators(ctx context.Context, in *ListClientCollaboratorsRequest, opts ...grpc.CallOption) (*Collaborators, error)
}

type clientAccessClient struct {
	cc *grpc.ClientConn
}

func NewClientAccessClient(cc *grpc.ClientConn) ClientAccessClient {
	return &clientAccessClient{cc}
}

func (c *clientAccessClient) ListRights(ctx context.Context, in *ClientIdentifiers, opts ...grpc.CallOption) (*Rights, error) {
	out := new(Rights)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ClientAccess/ListRights", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientAccessClient) GetCollaborator(ctx context.Context, in *GetClientCollaboratorRequest, opts ...grpc.CallOption) (*GetCollaboratorResponse, error) {
	out := new(GetCollaboratorResponse)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ClientAccess/GetCollaborator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientAccessClient) SetCollaborator(ctx context.Context, in *SetClientCollaboratorRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ClientAccess/SetCollaborator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientAccessClient) ListCollaborators(ctx context.Context, in *ListClientCollaboratorsRequest, opts ...grpc.CallOption) (*Collaborators, error) {
	out := new(Collaborators)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ClientAccess/ListCollaborators", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientAccessServer is the server API for ClientAccess service.
type ClientAccessServer interface {
	ListRights(context.Context, *ClientIdentifiers) (*Rights, error)
	// Get the rights of a collaborator (member) of the client.
	// Pseudo-rights in the response (such as the "_ALL" right) are not expanded.
	GetCollaborator(context.Context, *GetClientCollaboratorRequest) (*GetCollaboratorResponse, error)
	// Set the rights of a collaborator (member) on the client. It is required for the caller to
	// have all assigned or/and removed rights.
	// Setting a collaborator without rights, removes them.
	SetCollaborator(context.Context, *SetClientCollaboratorRequest) (*types.Empty, error)
	ListCollaborators(context.Context, *ListClientCollaboratorsRequest) (*Collaborators, error)
}

// UnimplementedClientAccessServer can be embedded to have forward compatible implementations.
type UnimplementedClientAccessServer struct {
}

func (*UnimplementedClientAccessServer) ListRights(ctx context.Context, req *ClientIdentifiers) (*Rights, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRights not implemented")
}
func (*UnimplementedClientAccessServer) GetCollaborator(ctx context.Context, req *GetClientCollaboratorRequest) (*GetCollaboratorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCollaborator not implemented")
}
func (*UnimplementedClientAccessServer) SetCollaborator(ctx context.Context, req *SetClientCollaboratorRequest) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetCollaborator not implemented")
}
func (*UnimplementedClientAccessServer) ListCollaborators(ctx context.Context, req *ListClientCollaboratorsRequest) (*Collaborators, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCollaborators not implemented")
}

func RegisterClientAccessServer(s *grpc.Server, srv ClientAccessServer) {
	s.RegisterService(&_ClientAccess_serviceDesc, srv)
}

func _ClientAccess_ListRights_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientAccessServer).ListRights(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ClientAccess/ListRights",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientAccessServer).ListRights(ctx, req.(*ClientIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientAccess_GetCollaborator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClientCollaboratorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientAccessServer).GetCollaborator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ClientAccess/GetCollaborator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientAccessServer).GetCollaborator(ctx, req.(*GetClientCollaboratorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientAccess_SetCollaborator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetClientCollaboratorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientAccessServer).SetCollaborator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ClientAccess/SetCollaborator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientAccessServer).SetCollaborator(ctx, req.(*SetClientCollaboratorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientAccess_ListCollaborators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClientCollaboratorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientAccessServer).ListCollaborators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ClientAccess/ListCollaborators",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientAccessServer).ListCollaborators(ctx, req.(*ListClientCollaboratorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClientAccess_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.ClientAccess",
	HandlerType: (*ClientAccessServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListRights",
			Handler:    _ClientAccess_ListRights_Handler,
		},
		{
			MethodName: "GetCollaborator",
			Handler:    _ClientAccess_GetCollaborator_Handler,
		},
		{
			MethodName: "SetCollaborator",
			Handler:    _ClientAccess_SetCollaborator_Handler,
		},
		{
			MethodName: "ListCollaborators",
			Handler:    _ClientAccess_ListCollaborators_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/client_services.proto",
}
