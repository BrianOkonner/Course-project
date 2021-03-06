// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package game

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

// ChooseclassClient is the client API for Chooseclass service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChooseclassClient interface {
	Do(ctx context.Context, in *Requestcls, opts ...grpc.CallOption) (*Responsecls, error)
}

type chooseclassClient struct {
	cc grpc.ClientConnInterface
}

func NewChooseclassClient(cc grpc.ClientConnInterface) ChooseclassClient {
	return &chooseclassClient{cc}
}

func (c *chooseclassClient) Do(ctx context.Context, in *Requestcls, opts ...grpc.CallOption) (*Responsecls, error) {
	out := new(Responsecls)
	err := c.cc.Invoke(ctx, "/swords.Chooseclass/Do", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChooseclassServer is the server API for Chooseclass service.
// All implementations must embed UnimplementedChooseclassServer
// for forward compatibility
type ChooseclassServer interface {
	Do(context.Context, *Requestcls) (*Responsecls, error)
	mustEmbedUnimplementedChooseclassServer()
}

// UnimplementedChooseclassServer must be embedded to have forward compatible implementations.
type UnimplementedChooseclassServer struct {
}

func (UnimplementedChooseclassServer) Do(context.Context, *Requestcls) (*Responsecls, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Do not implemented")
}
func (UnimplementedChooseclassServer) mustEmbedUnimplementedChooseclassServer() {}

// UnsafeChooseclassServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChooseclassServer will
// result in compilation errors.
type UnsafeChooseclassServer interface {
	mustEmbedUnimplementedChooseclassServer()
}

func RegisterChooseclassServer(s grpc.ServiceRegistrar, srv ChooseclassServer) {
	s.RegisterService(&Chooseclass_ServiceDesc, srv)
}

func _Chooseclass_Do_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Requestcls)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChooseclassServer).Do(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/swords.Chooseclass/Do",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChooseclassServer).Do(ctx, req.(*Requestcls))
	}
	return interceptor(ctx, in, info, handler)
}

// Chooseclass_ServiceDesc is the grpc.ServiceDesc for Chooseclass service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Chooseclass_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "swords.Chooseclass",
	HandlerType: (*ChooseclassServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Do",
			Handler:    _Chooseclass_Do_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "game.proto",
}
