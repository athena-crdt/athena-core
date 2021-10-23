package healtz

import (
	"context"

	pb "github.com/athena-crdt/athena-core/proto"
)

type Service struct {
	// Embed the unimplemented server
	pb.UnimplementedHealthServiceServer
}

func (h *Service) Check(ctx context.Context, req *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	return &pb.HealthCheckResponse{Status: pb.HealthCheckResponse_SERVING}, nil
}

func (h *Service) Watch(req *pb.HealthCheckRequest, stream pb.HealthService_WatchServer) error {
	panic("implement me")
}
