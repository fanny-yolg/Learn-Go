package main

import (
	"context"
	"log"

	pb "fanny/learn-grpc/calculator/proto"
)

func (s *Server) Calculator(ctx context.Context, in *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	log.Printf("Calculator function was invoked with %v\n and %v\n", in.Num1, in.Num2)
	resMath := in.Num1 + in.Num2
	return &pb.CalculatorResponse{
		Result: resMath,
	}, nil
}
