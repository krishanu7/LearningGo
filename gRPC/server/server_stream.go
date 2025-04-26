package main

import (
	"log"
	"time"

	pb "github.com/krishanu7/go-with-grpc/proto"
)

func (s *Server) SayHelloServerStream (req *pb.NamesList, stream pb.GreetService_SayHelloServerStreamServer) error {
	// Log the request
	log.Printf("Received request: %s", req.String());
	// Check if the request is nil
	if req.Names == nil {
		log.Println("Received nil request data")
	}
	// Create a response
	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "Hello " + name,
		}
		// Send the response
		if err := stream.Send(res); err != nil {
			log.Printf("Error sending response: %v", err)
			return err
		}
		time.Sleep(time.Second*2);
	}
	return nil
}