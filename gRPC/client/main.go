package main

import (
	pb "github.com/krishanu7/go-with-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const (
	Port = ":8080"
)

func main() {
	conn, err := grpc.NewClient("localhost"+Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList{
		Names: []string{"Akhil", "Alice", "Bob"},
	}

	// callSayHello(client)
	callSayHelloServerStream(client, names)
	//callSayHelloClientStream(client, names)
	// callSayHelloBidirectionalStream(client, names)
}
