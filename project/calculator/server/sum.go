package main

import (
	"context"
	"log"

	pb "github.com/tingwei1111/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum function was invoked with %v\n", in)
	return &pb.SumResponse{
		Result: in.FirstNumber + in.SecondNumber,
	},nil 
}
