package main

import (
	"calculator.com/calculator/calculatorpb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	fmt.Println("Hello I'm a calculator client")

	// User input
	fmt.Println("Please input a number in order to decompose prime numbers of: ")
	var num int32
	fmt.Scan(&num)

	clientConn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer clientConn.Close()
	client := calculatorpb.NewPrimeNumberCalculatorServiceClient(clientConn)

	doServerStreaming(client, num)
}

func doServerStreaming(oClient calculatorpb.PrimeNumberCalculatorServiceClient, num int32) {
	fmt.Println("Starting a Server Streaming API...")
	req := &calculatorpb.PrimeNumberCalculatorRequest{
		Number: num,
	}
	resultStream, err := oClient.PrimeNumberCalculator(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling server streaming API: %v", err)
	}

	for {
		msg, err := resultStream.Recv()
		if err == io.EOF {
			// end of stream
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream message: %v", err)
		}
		log.Printf("Response from streaming API: %v", msg.GetResult())
	}
}
