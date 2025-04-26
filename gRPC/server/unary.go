package main

import (
	"context"
	"log"

	pb "github.com/krishanu7/go-with-grpc/proto"
)

func (s *Server) SayHello(ctx context.Context, req *pb.NoParameterRequest) (*pb.HelloResponse, error) {
	// Log the request
	log.Printf("Received request: %s", req.String())
	// Check if the request is nil
	if req == nil {
		log.Println("Received nil request")
	}

	// Create a response
	res := &pb.HelloResponse{
		Message: "Hello!",
	}
	// Return the response
	return res, nil
}
