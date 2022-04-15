package grpcapi

import (
	"context"
	"log"

	"google.golang.org/grpc"

	"server1/internal/grpcapi/protobuff/user"
)

type Server struct {
	user.UnimplementedUserStorageServer
}

func NewServer() *grpc.Server {

	s := grpc.NewServer()
	user.RegisterUserStorageServer(s, &Server{})

	return s
}

func (s *Server) GetByID(ctx context.Context, in *user.GetUserRequest) (*user.UserResponse, error) {
	log.Printf("Received: %v\n", in.GetId())
	return &user.UserResponse{
		Id:   in.GetId(),
		Name: "Serge",
	}, nil
}
