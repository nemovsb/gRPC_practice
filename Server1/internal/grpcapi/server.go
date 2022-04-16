package grpcapi

import (
	"context"
	"log"

	"google.golang.org/grpc"

	"server1/internal/grpcapi/protobuff/user"
)

type Server struct {
	user.UnimplementedUserStorageServer
	DB
}

type DB interface {
	GetUserByID(id int32) *user.UserResponse
	//PutUser() *user.UserResponse
}

func NewServer(db DB) *grpc.Server {

	s := grpc.NewServer()
	user.RegisterUserStorageServer(s, &Server{DB: db})

	return s
}

func (s *Server) GetByID(ctx context.Context, in *user.GetUserRequest) (*user.UserResponse, error) {
	log.Printf("Received: %v\n", in.GetId())

	response := s.DB.GetUserByID(in.GetId())

	return response, nil
}
