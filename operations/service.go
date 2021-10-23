package operations

import (
	"context"

	pb "github.com/athena-crdt/athena-core/proto"
)

type Service struct {
	pb.UnimplementedOperationServiceServer
}

func (o *Service) Receive(ctx context.Context, stream *pb.ReceiveRequest) (*pb.Operation, error) {
	panic("implement me")
}

func (o *Service) Watch(request *pb.ReceiveRequest, stream pb.OperationService_WatchServer) error {
	panic("implement me")
}
