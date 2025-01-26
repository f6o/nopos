// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: hands/hands.proto

package hands

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Dealer_DealCard_FullMethodName = "/Dealer/DealCard"
)

// DealerClient is the client API for Dealer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DealerClient interface {
	DealCard(ctx context.Context, in *DealRequest, opts ...grpc.CallOption) (*Card, error)
}

type dealerClient struct {
	cc grpc.ClientConnInterface
}

func NewDealerClient(cc grpc.ClientConnInterface) DealerClient {
	return &dealerClient{cc}
}

func (c *dealerClient) DealCard(ctx context.Context, in *DealRequest, opts ...grpc.CallOption) (*Card, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Card)
	err := c.cc.Invoke(ctx, Dealer_DealCard_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DealerServer is the server API for Dealer service.
// All implementations must embed UnimplementedDealerServer
// for forward compatibility.
type DealerServer interface {
	DealCard(context.Context, *DealRequest) (*Card, error)
	mustEmbedUnimplementedDealerServer()
}

// UnimplementedDealerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDealerServer struct{}

func (UnimplementedDealerServer) DealCard(context.Context, *DealRequest) (*Card, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DealCard not implemented")
}
func (UnimplementedDealerServer) mustEmbedUnimplementedDealerServer() {}
func (UnimplementedDealerServer) testEmbeddedByValue()                {}

// UnsafeDealerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DealerServer will
// result in compilation errors.
type UnsafeDealerServer interface {
	mustEmbedUnimplementedDealerServer()
}

func RegisterDealerServer(s grpc.ServiceRegistrar, srv DealerServer) {
	// If the following call pancis, it indicates UnimplementedDealerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Dealer_ServiceDesc, srv)
}

func _Dealer_DealCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DealRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DealerServer).DealCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Dealer_DealCard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DealerServer).DealCard(ctx, req.(*DealRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Dealer_ServiceDesc is the grpc.ServiceDesc for Dealer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Dealer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Dealer",
	HandlerType: (*DealerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DealCard",
			Handler:    _Dealer_DealCard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hands/hands.proto",
}
