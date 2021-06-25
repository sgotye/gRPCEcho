package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"google.golang.org/grpc"
	"github.com/sgotye/gRPCEcho/pingpong"
)

const (
	port = ":8080"
)

type server struct {
	mu 			 sync.RWMutex
	countClients uint32
	clientId     uint32
	idChan		 chan uint32
	pingpong.UnimplementedPingPongServer
}

func (s *server) addClient() {
	var currentId uint32
	s.mu.Lock()
	s.countClients += 1
	s.clientId += 1
	currentId = s.clientId
	s.mu.Unlock()
	s.idChan <- currentId
}

func (s *server) delClient() {
	s.mu.Lock()
	s.countClients -= 1
	s.mu.Unlock()
}

func (s *server) SendPing(ctx context.Context, in *pingpong.Ping) (*pingpong.Pong, error)  {
	var response string
	s.addClient()
	defer s.delClient()
	log.Printf("Received ping.")

	currentId := <- s.idChan
	s.mu.RLock()
	currentClentsNum := s.countClients
	s.mu.RUnlock()

	if currentClentsNum > 1 {
		response = fmt.Sprintf("pong%v", currentId)
	} else {
		response = "pong"
	}
	return &pingpong.Pong{Message: response}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pingpong.RegisterPingPongServer(s, &server{
		idChan: make(chan uint32, 10),
	})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}