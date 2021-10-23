package swarm

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/athena-crdt/athena-core/proto"
)

type Service struct {
	pb.UnimplementedSwarmServiceServer
}

func (s *Service) Init(ctx context.Context, _ *emptypb.Empty) (*pb.InitResponse, error) {
	panic("implement me")
}

func (s *Service) Watch(_ *emptypb.Empty, stream pb.SwarmService_WatchServer) error {
	panic("implement me")
}
