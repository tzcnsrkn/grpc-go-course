package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"calculator.com/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) PrimeNumberCalculator(req *calculatorpb.PrimeNumberCalculatorRequest, stream calculatorpb.PrimeNumberCalculatorService_PrimeNumberCalculatorServer) error {
	fmt.Printf("PrimeNumberCalculator is invoked with: %v", req)
	input := req.GetNumber()
	k := int32(2)
	for input > 1 {
		if input%k == 0 {
			res := &calculatorpb.PrimeNumberCalculatorResponse{
				Result: fmt.Sprint(k) + " ",
			}
			input /= k
			stream.Send(res)
			time.Sleep(1000 * time.Millisecond)
		} else {
			k++
		}
	}
	return nil
}

func main() {
	fmt.Println("Hello world! I'm a calculator")
	lis, err := net.Listen("tcp", "0.0.0.0:50052")

	if err != nil {
		log.Fatalf("Error while listening port: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterPrimeNumberCalculatorServiceServer(s, &server{})

	err2 := s.Serve(lis)
	if err2 != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
