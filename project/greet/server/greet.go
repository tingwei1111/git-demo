package main

import (
	"context"
	"log"

	pb "github.com/tingwei1111/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet was invoked with %v\n", in)
	return &pb.GreetResponse{
		Result: "Hello" + in.FirstName,
	}, nil

}
