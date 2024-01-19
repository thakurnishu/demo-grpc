package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/thakurnishu/demo-grpc/protobuf"
	"google.golang.org/grpc"
)

type Server struct {
	ListenAddr string
	Protocol   string
	lis        net.Listener
}

type helloServer struct {
	pb.GreetServiceServer
}

func NewServer(protobuf string, listenAddr string) *Server {
	return &Server{
		ListenAddr: fmt.Sprintf(":%s", listenAddr),
		Protocol:   protocol,
	}
}

func (s *Server) Run() {

	lis, err := net.Listen(s.Protocol, s.ListenAddr)
	if err != nil {
		log.Fatalf("failed to start the server %s", err.Error())
	}
	s.lis = lis

	grpcServer := grpc.NewServer()

	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})

	if err := grpcServer.Serve(s.lis); err != nil {
		log.Fatalf("failed to start the grpc server %s", err.Error())
	}

}
