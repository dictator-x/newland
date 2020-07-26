package main

import (
	"context"
	pb2 "github.com/dictator-x/newland/grpc/demo/pb"
	"google.golang.org/grpc"
	"log"
)

const (
	address = "localhost:50001"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := pb2.NewGreeterClient(conn)

	name := "messi"
	r, err := c.SayHello(context.Background(), &pb2.HelloRequest{Name: name})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Println(r.Message)
}