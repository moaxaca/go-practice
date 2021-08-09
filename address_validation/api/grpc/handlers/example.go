package handlers

import (
	"context"
	pb "io.parcely.address_validation/api/grpc/proto"
	"log"
)

type exampleHandle struct {
	pb.UnimplementedExampleServer
}

func (s *exampleHandle) Ping(_ context.Context, req *pb.ExampleRequest) (*pb.ExampleResponse, error) {
	log.Printf("Received: %v", req.GetMessage())
	return &pb.ExampleResponse{Relay: "Message " + req.GetMessage()}, nil
}

func CreateExampleHandler() *exampleHandle {
	return &exampleHandle{}
}
