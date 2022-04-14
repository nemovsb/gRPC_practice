package configuration

import (
	"api/internal/transport/http_api"
)

func NewHTTPServerConfig(appConfig *AppConfig) *http_api.ServerConfig {
	return &http_api.ServerConfig{
		Port: appConfig.HTTP.port,
	}
}
