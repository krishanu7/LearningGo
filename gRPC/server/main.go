package main

import (
	"fmt"
	"log"
	"net"
	pb "github.com/krishanu7/go-with-grpc/proto"
	"google.golang.org/grpc"
)

// Define the port
const (
	Port = ":8080"
)
// Struct to hold the server api
type Server struct {
	pb.GreetServiceServer
}

func main() {
	// Create a new Server
	s := &Server{}
	// Create a new gRPC server
	grpcServer := grpc.NewServer()
	// Register the server with the gRPC server
	pb.RegisterGreetServiceServer(grpcServer, s)
	// Create a listener on the port
	lis, err := net.Listen("tcp", Port)
	if err != nil {
		log.Fatalf("[INFO] Failed to listen: %v", err)
	}
	// Start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("[INFO] Failed to serve: %v", err)
	}
	// Print the server address
	fmt.Printf("Server is running on port %s\n", Port)
}