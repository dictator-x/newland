package main

import (
	"context"
	"fmt"
	pb2 "github.com/dictator-x/newland/grpc/demo/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	PORT = ":50001"
)

type server struct{}
func (s *server) SayHello(ctx context.Context, in *pb2.HelloRequest) (*pb2.HelloReply, error) {
	fmt.Println("request: ", in.Name)
	return &pb2.HelloReply{Message: "Hello " + in.Name}, nil
}
func main() {
	lis, err := net.Listen("tcp", PORT)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb2.RegisterGreeterServer(s, &server{})
	log.Println("Demo grpc server start")
	s.Serve(lis)
}