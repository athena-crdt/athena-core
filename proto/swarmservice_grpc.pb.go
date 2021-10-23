// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SwarmServiceClient is the client API for SwarmService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SwarmServiceClient interface {
	// Receive all peers info
	Init(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*InitResponse, error)
	// Watch on latest peers! - Start with one node
	// --- eventually it will form a network.
	Watch(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (SwarmService_WatchClient, error)
}

type swarmServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSwarmServiceClient(cc grpc.ClientConnInterface) SwarmServiceClient {
	return &swarmServiceClient{cc}
}

func (c *swarmServiceClient) Init(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*InitResponse, error) {
	out := new(InitResponse)
	err := c.cc.Invoke(ctx, "/swarm.SwarmService/Init", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *swarmServiceClient) Watch(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (SwarmService_WatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &SwarmService_ServiceDesc.Streams[0], "/swarm.SwarmService/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &swarmServiceWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SwarmService_WatchClient interface {
	Recv() (*PeerInfo, error)
	grpc.ClientStream
}

type swarmServiceWatchClient struct {
	grpc.ClientStream
}

func (x *swarmServiceWatchClient) Recv() (*PeerInfo, error) {
	m := new(PeerInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SwarmServiceServer is the server API for SwarmService service.
// All implementations must embed UnimplementedSwarmServiceServer
// for forward compatibility
type SwarmServiceServer interface {
	// Receive all peers info
	Init(context.Context, *emptypb.Empty) (*InitResponse, error)
	// Watch on latest peers! - Start with one node
	// --- eventually it will form a network.
	Watch(*emptypb.Empty, SwarmService_WatchServer) error
	mustEmbedUnimplementedSwarmServiceServer()
}

// UnimplementedSwarmServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSwarmServiceServer struct {
}

func (UnimplementedSwarmServiceServer) Init(context.Context, *emptypb.Empty) (*InitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Init not implemented")
}
func (UnimplementedSwarmServiceServer) Watch(*emptypb.Empty, SwarmService_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "method Watch not implemented")
}
func (UnimplementedSwarmServiceServer) mustEmbedUnimplementedSwarmServiceServer() {}

// UnsafeSwarmServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SwarmServiceServer will
// result in compilation errors.
type UnsafeSwarmServiceServer interface {
	mustEmbedUnimplementedSwarmServiceServer()
}

func RegisterSwarmServiceServer(s grpc.ServiceRegistrar, srv SwarmServiceServer) {
	s.RegisterService(&SwarmService_ServiceDesc, srv)
}

func _SwarmService_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwarmServiceServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/swarm.SwarmService/Init",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwarmServiceServer).Init(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SwarmService_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SwarmServiceServer).Watch(m, &swarmServiceWatchServer{stream})
}

type SwarmService_WatchServer interface {
	Send(*PeerInfo) error
	grpc.ServerStream
}

type swarmServiceWatchServer struct {
	grpc.ServerStream
}

func (x *swarmServiceWatchServer) Send(m *PeerInfo) error {
	return x.ServerStream.SendMsg(m)
}

// SwarmService_ServiceDesc is the grpc.ServiceDesc for SwarmService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SwarmService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "swarm.SwarmService",
	HandlerType: (*SwarmServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Init",
			Handler:    _SwarmService_Init_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _SwarmService_Watch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/swarmservice.proto",
}
