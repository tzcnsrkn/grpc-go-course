package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"example.com/greetings/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("\nGreet function was invoked with %v", req)
	fName := req.GetGreeting().GetFirstName()
	result := "Hello " + fName
	resp := &greetpb.GreetResponse{
		Result: result,
	}
	return resp, nil
}

func main() {
	fmt.Println("Hello world")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
