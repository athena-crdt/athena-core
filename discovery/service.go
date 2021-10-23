//  Copyright 2021, athena-crdt authors.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package discovery

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
