package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/thakurnishu/demo-grpc/protobuf"
)

// Unaray
func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}

// Server Streaming
func (s *helloServer) SayHelloServerStreaming(req *pb.NameList, stream pb.GreetService_SayHelloServerStreamingServer) error {

	log.Printf("got request with names: %v", req.Names)

	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "Hello " + name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}

	return nil
}

// CLient Streaming
func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {

	var messages []string

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessageList{
				Messages: messages,
			})
		}
		if err != nil {
			return err
		}
		log.Printf("got request with name : %s", req.Name)
		messages = append(messages, "hello "+req.Name)
	}
}

// Bi-DirectionalStreaming
func (s *helloServer) SayHelloBiDirectionalStreaming(stream pb.GreetService_SayHelloBiDirectionalStreamingServer) error {

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("error reciving request: %s", err.Error())
		}
		log.Printf("got request with name : %s", req.Name)

		res := &pb.HelloResponse{
			Message: "Hello " + req.Name,
		}

		if err := stream.Send(res); err != nil {
			log.Fatalf("error sending response: %s", err.Error())
		}
	}
}
