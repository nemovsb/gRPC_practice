package grpc_api

import (
	"context"
	"fmt"
	"log"

	"api/internal/domain"
	"api/internal/transport/grpc_api/protobuff/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCConfig struct {
	Host string
	Port string
}

type GRPCClient struct {
	User user.UserStorageClient
	conn *grpc.ClientConn
}

func NewGRPCClient(cfg *GRPCConfig) *GRPCClient {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	fmt.Printf("COnnection: %+v\n", conn)

	user := user.NewUserStorageClient(conn)

	return &GRPCClient{
		User: user,
		conn: conn,
	}
}

func (g *GRPCClient) Close() error {
	return g.conn.Close()
}

func (g *GRPCClient) GetUserByID(id int32) (*domain.User, error) {

	u := g.User

	r, err := u.GetByID(context.Background(), &user.GetUserRequest{Id: id})
	if err != nil {
		return nil, fmt.Errorf("get user by ID error: %s", err)
	}
	return &domain.User{
		Id:   r.Id,
		Name: r.Name,
	}, nil
}
