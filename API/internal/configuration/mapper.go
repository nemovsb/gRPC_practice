package configuration

import (
	"api/internal/transport/grpc_api"
	"api/internal/transport/http_api"
)

func NewHTTPServerConfig(appConfig *AppConfig) *http_api.ServerConfig {
	return &http_api.ServerConfig{
		Port: appConfig.HTTP.port,
	}
}

func NewGRPCCOnfig(appConfig *AppConfig) *grpc_api.GRPCConfig {
	return &grpc_api.GRPCConfig{
		Host: appConfig.GRPC.host,
		Port: appConfig.GRPC.port,
	}
}
