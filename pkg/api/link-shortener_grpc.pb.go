// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// LinkShortenerClient is the client API for LinkShortener service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LinkShortenerClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type linkShortenerClient struct {
	cc grpc.ClientConnInterface
}

func NewLinkShortenerClient(cc grpc.ClientConnInterface) LinkShortenerClient {
	return &linkShortenerClient{cc}
}

func (c *linkShortenerClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/api.linkShortener/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkShortenerClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/api.linkShortener/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LinkShortenerServer is the server API for LinkShortener service.
// All implementations must embed UnimplementedLinkShortenerServer
// for forward compatibility
type LinkShortenerServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	mustEmbedUnimplementedLinkShortenerServer()
}

// UnimplementedLinkShortenerServer must be embedded to have forward compatible implementations.
type UnimplementedLinkShortenerServer struct {
}

func (UnimplementedLinkShortenerServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedLinkShortenerServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedLinkShortenerServer) mustEmbedUnimplementedLinkShortenerServer() {}

// UnsafeLinkShortenerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LinkShortenerServer will
// result in compilation errors.
type UnsafeLinkShortenerServer interface {
	mustEmbedUnimplementedLinkShortenerServer()
}

func RegisterLinkShortenerServer(s grpc.ServiceRegistrar, srv LinkShortenerServer) {
	s.RegisterService(&LinkShortener_ServiceDesc, srv)
}

func _LinkShortener_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkShortenerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.linkShortener/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkShortenerServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinkShortener_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkShortenerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.linkShortener/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkShortenerServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LinkShortener_ServiceDesc is the grpc.ServiceDesc for LinkShortener service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LinkShortener_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.linkShortener",
	HandlerType: (*LinkShortenerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _LinkShortener_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _LinkShortener_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "link-shortener/proto/link-shortener.proto",
}
