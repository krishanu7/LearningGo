package main

import (
	"context"
	"log"
	"time"

	pb "github.com/krishanu7/go-with-grpc/proto"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParameterRequest{})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Print("Response from server: ", res.Message);
}
