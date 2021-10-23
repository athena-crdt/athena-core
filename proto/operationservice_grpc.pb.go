// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// OperationServiceClient is the client API for OperationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OperationServiceClient interface {
	// Unary RPC requesting committed operation of counter_id + 1
	Receive(ctx context.Context, in *ReceiveRequest, opts ...grpc.CallOption) (*Operation, error)
	// Serverside streaming until both of their lamport counters matches.
	Watch(ctx context.Context, in *ReceiveRequest, opts ...grpc.CallOption) (OperationService_WatchClient, error)
}

type operationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOperationServiceClient(cc grpc.ClientConnInterface) OperationServiceClient {
	return &operationServiceClient{cc}
}

func (c *operationServiceClient) Receive(ctx context.Context, in *ReceiveRequest, opts ...grpc.CallOption) (*Operation, error) {
	out := new(Operation)
	err := c.cc.Invoke(ctx, "/operations.OperationService/Receive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operationServiceClient) Watch(ctx context.Context, in *ReceiveRequest, opts ...grpc.CallOption) (OperationService_WatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &OperationService_ServiceDesc.Streams[0], "/operations.OperationService/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &operationServiceWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OperationService_WatchClient interface {
	Recv() (*Operation, error)
	grpc.ClientStream
}

type operationServiceWatchClient struct {
	grpc.ClientStream
}

func (x *operationServiceWatchClient) Recv() (*Operation, error) {
	m := new(Operation)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OperationServiceServer is the server API for OperationService service.
// All implementations must embed UnimplementedOperationServiceServer
// for forward compatibility
type OperationServiceServer interface {
	// Unary RPC requesting committed operation of counter_id + 1
	Receive(context.Context, *ReceiveRequest) (*Operation, error)
	// Serverside streaming until both of their lamport counters matches.
	Watch(*ReceiveRequest, OperationService_WatchServer) error
	mustEmbedUnimplementedOperationServiceServer()
}

// UnimplementedOperationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOperationServiceServer struct {
}

func (UnimplementedOperationServiceServer) Receive(context.Context, *ReceiveRequest) (*Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Receive not implemented")
}
func (UnimplementedOperationServiceServer) Watch(*ReceiveRequest, OperationService_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "method Watch not implemented")
}
func (UnimplementedOperationServiceServer) mustEmbedUnimplementedOperationServiceServer() {}

// UnsafeOperationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OperationServiceServer will
// result in compilation errors.
type UnsafeOperationServiceServer interface {
	mustEmbedUnimplementedOperationServiceServer()
}

func RegisterOperationServiceServer(s grpc.ServiceRegistrar, srv OperationServiceServer) {
	s.RegisterService(&OperationService_ServiceDesc, srv)
}

func _OperationService_Receive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReceiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationServiceServer).Receive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/operations.OperationService/Receive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationServiceServer).Receive(ctx, req.(*ReceiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperationService_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReceiveRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OperationServiceServer).Watch(m, &operationServiceWatchServer{stream})
}

type OperationService_WatchServer interface {
	Send(*Operation) error
	grpc.ServerStream
}

type operationServiceWatchServer struct {
	grpc.ServerStream
}

func (x *operationServiceWatchServer) Send(m *Operation) error {
	return x.ServerStream.SendMsg(m)
}

// OperationService_ServiceDesc is the grpc.ServiceDesc for OperationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OperationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "operations.OperationService",
	HandlerType: (*OperationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Receive",
			Handler:    _OperationService_Receive_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _OperationService_Watch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/operationservice.proto",
}