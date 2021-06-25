package main

import (
	"context"
	"github.com/sgotye/gRPCEcho/pingpong"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address = "localhost:8080"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()
	c := pingpong.NewPingPongClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SendPing(ctx, &pingpong.Ping{Message: "ping"})
	if err != nil {
		log.Fatalf("could not recieve response: %v", err)
	}
	log.Printf(r.GetMessage())
}
