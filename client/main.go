package main

import (
	"fmt"
	"log"

	pb "github.com/thakurnishu/demo-grpc/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = "3000"
)

func main() {
	conn, err := grpc.Dial("localhost"+":"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect %s", err.Error())
	}

	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NameList{
		Names: []string{"Nishu", "Viku", "Bob"},
	}

	// Unaray
	fmt.Println("Unaray Stream")
	callSayHello(client)
	fmt.Println()

	// ServerStreaming
	fmt.Println("Server Streaming")
	callSayHelloServerStream(names, client)
	fmt.Println()

	// ClientStreaming
	fmt.Println("Client Streaming")
	callSayHelloClientStream(names, client)
	fmt.Println()

	// Bi-DirectionalStreaming
	fmt.Println("Client Streaming")
	callSayHelloBiDirectionalStream(names, client)
	fmt.Println()
}
