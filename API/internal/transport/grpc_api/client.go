package grpc_api

import (
	"fmt"
	"log"

	"google.golang.org/grpc"
)

type GRPCConfig struct {
	Host string
	Port string
}

func NewConnection(cfg *GRPCConfig) *grpc.ClientConn {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c:= pb.
	defer conn.Close()
}
