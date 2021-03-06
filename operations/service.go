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
