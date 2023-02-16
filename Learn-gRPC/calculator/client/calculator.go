package main

import (
	"context"
	pb "fanny/learn-grpc/calculator/proto"
	"log"
)

func doMath(c pb.CalculatorServiceClient) {
	log.Println("doMath was invoked")

	res, err := c.Calculator(context.Background(), &pb.CalculatorRequest{
		Num1: 10,
		Num2: 3,
	})

	if err != nil {
		log.Fatalf("Could not calculate: %v\n", err)
	}

	log.Printf("Calculation: %v\n", res.Result)
}
