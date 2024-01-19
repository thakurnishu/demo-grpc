package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/thakurnishu/demo-grpc/protobuf"
)

// Unaray
func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("could not greet %s", err.Error())
	}
	log.Printf("%s", res.Message)
}

// Server Streaming
func callSayHelloServerStream(names *pb.NameList, client pb.GreetServiceClient) {
	log.Println("streaming started...")
	ctx := context.Background()

	stream, err := client.SayHelloServerStreaming(ctx, names)
	if err != nil {
		log.Fatalf("could not stream %s", err.Error())
	}

free:
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break free
		} else if err != nil {
			log.Fatalf("error while streaming.. %s", err.Error())
		}
		log.Println(resp.Message)
	}
	log.Println("stream finished ")
}

// ClientStreaming
func callSayHelloClientStream(names *pb.NameList, client pb.GreetServiceClient) {

	log.Println("starting client streaming...")
	ctx := context.Background()

	stream, err := client.SayHelloClientStreaming(ctx)
	if err != nil {
		log.Fatalf("could not send names %s", err.Error())
	}

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("error while sending %s", err.Error())
		}
		log.Printf("sent request with name : %s", req.Name)

		time.Sleep(2 * time.Second)
	}

	log.Println("stream finished")

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error getting response %s", err.Error())
	}
	log.Printf("response from server : %v\n", res.Messages)
}

// Bi-DirectionalStreaming
func callSayHelloBiDirectionalStream(names *pb.NameList, client pb.GreetServiceClient) {
	log.Println("bi-directional streaming started ...")
	ctx := context.Background()

	stream, err := client.SayHelloBiDirectionalStreaming(ctx)
	if err != nil {
		log.Fatalf("error streaming : %s", err.Error())
	}

	waitch := make(chan struct{})

	go func() {
	free:
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				break free
			}
			if err != nil {
				log.Fatalf("error reciveing response : %s", err.Error())
			}
			log.Printf("message recivied : %s", resp.Message)
		}
		close(waitch)
	}()

	for _, name := range names.Names {

		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("error sending request : %s", err.Error())
		}
		time.Sleep(2 * time.Second)
	}
	stream.CloseSend()
	<-waitch
	log.Println("bi-directional streaming finished")

}
