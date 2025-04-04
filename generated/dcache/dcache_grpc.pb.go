// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: dcache.proto

package dcache

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	StripeService_Ping_FullMethodName         = "/dcache.StripeService/Ping"
	StripeService_GetStripe_FullMethodName    = "/dcache.StripeService/GetStripe"
	StripeService_PutStripe_FullMethodName    = "/dcache.StripeService/PutStripe"
	StripeService_RemoveStripe_FullMethodName = "/dcache.StripeService/RemoveStripe"
)

// StripeServiceClient is the client API for StripeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Define the service with RPC methods
type StripeServiceClient interface {
	Ping(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	// Fetch a stripe from the node
	GetStripe(ctx context.Context, in *GetStripeRequest, opts ...grpc.CallOption) (*Stripe, error)
	// Store a stripe on the node
	PutStripe(ctx context.Context, in *Stripe, opts ...grpc.CallOption) (*empty.Empty, error)
	// Delete a stripe from the node
	RemoveStripe(ctx context.Context, in *RemoveStripeRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type stripeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStripeServiceClient(cc grpc.ClientConnInterface) StripeServiceClient {
	return &stripeServiceClient{cc}
}

func (c *stripeServiceClient) Ping(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, StripeService_Ping_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stripeServiceClient) GetStripe(ctx context.Context, in *GetStripeRequest, opts ...grpc.CallOption) (*Stripe, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Stripe)
	err := c.cc.Invoke(ctx, StripeService_GetStripe_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stripeServiceClient) PutStripe(ctx context.Context, in *Stripe, opts ...grpc.CallOption) (*empty.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, StripeService_PutStripe_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stripeServiceClient) RemoveStripe(ctx context.Context, in *RemoveStripeRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, StripeService_RemoveStripe_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StripeServiceServer is the server API for StripeService service.
// All implementations must embed UnimplementedStripeServiceServer
// for forward compatibility.
//
// Define the service with RPC methods
type StripeServiceServer interface {
	Ping(context.Context, *empty.Empty) (*empty.Empty, error)
	// Fetch a stripe from the node
	GetStripe(context.Context, *GetStripeRequest) (*Stripe, error)
	// Store a stripe on the node
	PutStripe(context.Context, *Stripe) (*empty.Empty, error)
	// Delete a stripe from the node
	RemoveStripe(context.Context, *RemoveStripeRequest) (*empty.Empty, error)
	mustEmbedUnimplementedStripeServiceServer()
}

// UnimplementedStripeServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedStripeServiceServer struct{}

func (UnimplementedStripeServiceServer) Ping(context.Context, *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedStripeServiceServer) GetStripe(context.Context, *GetStripeRequest) (*Stripe, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStripe not implemented")
}
func (UnimplementedStripeServiceServer) PutStripe(context.Context, *Stripe) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutStripe not implemented")
}
func (UnimplementedStripeServiceServer) RemoveStripe(context.Context, *RemoveStripeRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveStripe not implemented")
}
func (UnimplementedStripeServiceServer) mustEmbedUnimplementedStripeServiceServer() {}
func (UnimplementedStripeServiceServer) testEmbeddedByValue()                       {}

// UnsafeStripeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StripeServiceServer will
// result in compilation errors.
type UnsafeStripeServiceServer interface {
	mustEmbedUnimplementedStripeServiceServer()
}

func RegisterStripeServiceServer(s grpc.ServiceRegistrar, srv StripeServiceServer) {
	// If the following call pancis, it indicates UnimplementedStripeServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&StripeService_ServiceDesc, srv)
}

func _StripeService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StripeServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StripeService_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StripeServiceServer).Ping(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _StripeService_GetStripe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStripeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StripeServiceServer).GetStripe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StripeService_GetStripe_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StripeServiceServer).GetStripe(ctx, req.(*GetStripeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StripeService_PutStripe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Stripe)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StripeServiceServer).PutStripe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StripeService_PutStripe_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StripeServiceServer).PutStripe(ctx, req.(*Stripe))
	}
	return interceptor(ctx, in, info, handler)
}

func _StripeService_RemoveStripe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveStripeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StripeServiceServer).RemoveStripe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StripeService_RemoveStripe_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StripeServiceServer).RemoveStripe(ctx, req.(*RemoveStripeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StripeService_ServiceDesc is the grpc.ServiceDesc for StripeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StripeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dcache.StripeService",
	HandlerType: (*StripeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _StripeService_Ping_Handler,
		},
		{
			MethodName: "GetStripe",
			Handler:    _StripeService_GetStripe_Handler,
		},
		{
			MethodName: "PutStripe",
			Handler:    _StripeService_PutStripe_Handler,
		},
		{
			MethodName: "RemoveStripe",
			Handler:    _StripeService_RemoveStripe_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dcache.proto",
}
