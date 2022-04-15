package main

import (
	"api/internal/configuration"
	"api/internal/transport/grpc_api"
	"api/internal/transport/http_api"
	"api/internal/transport/http_api/controllers"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	group "github.com/oklog/run"
)

var ErrOsSignal = errors.New("got os signal")

func main() {
	config, err := configuration.ViperConfigurationProvider(os.Getenv("GOLANG_ENVIRONMENT"), false)
	if err != nil {
		log.Fatal("Read config error: ", err)
	}

	log.Printf("Config: %+v\nconfig\n", config)

	grpcClient := grpc_api.NewGRPCClient(configuration.NewGRPCCOnfig(config))

	userHandler := controllers.NewUserHandler(grpcClient)
	set := controllers.NewHandlerSet(userHandler)
	handler := controllers.NewRouter(*set)
	server := http_api.NewServer(*configuration.NewHTTPServerConfig(config), handler)

	var (
		serviceGroup        group.Group
		interruptionChannel = make(chan os.Signal, 1)
	)

	serviceGroup.Add(func() error {
		signal.Notify(interruptionChannel, syscall.SIGINT, syscall.SIGTERM)
		osSignal := <-interruptionChannel

		return fmt.Errorf("%w: %s", ErrOsSignal, osSignal)
	}, func(error) {
		interruptionChannel <- syscall.SIGINT
	})

	serviceGroup.Add(func() error {
		log.Println("HTTP API started")

		return server.Run()
	}, func(error) {
		err = server.Shutdown()

		grpcClient.Close()

		log.Println("shutdown Http Server error")
	})

	err = serviceGroup.Run()
}
