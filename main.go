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

package main

import (
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/athena-crdt/athena-core/discovery"
	"github.com/athena-crdt/athena-core/healthz"
	"github.com/athena-crdt/athena-core/operations"
	pb "github.com/athena-crdt/athena-core/proto"
)

var port = flag.String("port", "50051", "grpc server port")

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", net.JoinHostPort("0.0.0.0", *port))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	pb.RegisterHealthServiceServer(s, &healthz.Service{})
	pb.RegisterOperationServiceServer(s, &operations.Service{})
	pb.RegisterDiscoveryServiceServer(s, &discovery.Service{})

	reflection.Register(s)
	log.Println("Service Registration Done!!")

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGKILL, syscall.SIGINT)
	go func() {
		defer close(sig)
		if err := s.Serve(lis); err != nil {
			// TODO: Change logger
			log.Println("Failed to run gRPC service server")
		}
	}()

	log.Println("Server is running on port: ", *port)

	<-sig
	s.Stop()
	log.Println("Server Stopped Successfully")
}
