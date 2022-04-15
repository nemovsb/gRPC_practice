package app

import (
	"net/http"
)

type App struct {
	HTTPServer http.Server
	GRPCClient grpc.
}
