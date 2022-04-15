package app

import (
	"api/internal/transport/grpc_api"
	"net/http"
)

type App struct {
	HTTPServer *http.Server
	GRPCClient *grpc_api.GRPCClient
}

func NewApp(apiServer *http.Server, grpcClient *grpc_api.GRPCClient) *App {
	return &App{
		HTTPServer: apiServer,
		GRPCClient: grpcClient,
	}
}
