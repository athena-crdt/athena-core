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

	"github.com/athena-crdt/athena-core/healtz"
	"github.com/athena-crdt/athena-core/operations"
	pb "github.com/athena-crdt/athena-core/proto"
	"github.com/athena-crdt/athena-core/swarm"
)

var port = flag.String("port", "50051", "grpc server port")

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", net.JoinHostPort("0.0.0.0", *port))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	pb.RegisterHealthServiceServer(s, &healtz.Service{})
	pb.RegisterOperationServiceServer(s, &operations.Service{})
	pb.RegisterSwarmServiceServer(s, &swarm.Service{})

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
