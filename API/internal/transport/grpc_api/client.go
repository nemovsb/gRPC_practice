package grpc_api

import (
	"context"
	"fmt"
	"log"

	"api/internal/domain"
	"api/internal/transport/grpc_api/protobuff/user"

	"google.golang.org/grpc"
)

type GRPCConfig struct {
	Host string
	Port string
}

type GRPCClient struct {
	User user.UserStorageClient
}

func NewGRPCClient(cfg *GRPCConfig) *GRPCClient {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	user := user.NewUserStorageClient(conn)

	return &GRPCClient{User: user}
}

func (g *GRPCClient) GetUserByID(id int32) (*domain.User, error) {
	r, err := g.User.GetByID(context.Background(), &user.GetUserRequest{Id: 123})
	if err != nil {
		return nil, fmt.Errorf("get user by ID error: %s", err)
	}
	return &domain.User{
		Id:   r.Id,
		Name: r.Name,
	}, nil
}
