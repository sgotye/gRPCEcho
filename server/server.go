package main

import (
	"context"
	"log"
	"net"
	"google.golang.org/grpc"
	"github.com/sgotye/gRPCEcho/pingpong"
)

const (
	port = ":8080"
)

type server struct {
	pingpong.UnimplementedPingPongServer
}

func (s *server) SendPing(ctx context.Context, in *pingpong.Ping) (*pingpong.Pong, error)  {
	log.Printf("Received ping.")
	return &pingpong.Pong{Message: "pong"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pingpong.RegisterPingPongServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}