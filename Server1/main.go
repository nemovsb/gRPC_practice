package main

import (
	"log"
	"net"
	"server1/internal/grpcapi"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpcapi.NewServer()

	log.Printf("server listening at %v", listener.Addr())

	err = server.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
