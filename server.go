package main

import (
	"log"
	"net"

	"chat-service/chat"

	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp",":9000")

	if err != nil {
		log.Fatalf("Failed to launch on port 9000: %v", err)
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}
}