package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	user_store_grpc_v1 "github.com/golerplate/contracts/clients/user-store-svc/v1/grpc"
	pkgconfig "github.com/golerplate/pkg/config"
	"github.com/rs/zerolog/log"

	"github.com/golerplate/user-gtw/internal/config"
	handlers_http "github.com/golerplate/user-gtw/internal/handlers/http"
	service_v1 "github.com/golerplate/user-gtw/internal/service/v1"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	// listen to signals
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM)

	// parse configuration
	config := &config.Config{}
	if err := pkgconfig.ParseConfig(config); err != nil {
		log.Fatal().
			Err(err).
			Msg("main: unable to parse config")
	}

	fmt.Print("coucou", config)

	userStoreClient := user_store_grpc_v1.NewUserStoreSvcClient(ctx, config.UserStoreSvcConfig.Addr)

	service, err := service_v1.NewService(ctx, userStoreClient)
	if err != nil {
		log.Fatal().Err(err).
			Msg("main: unable to create service")
	}

	// create http server
	httpServer, err := handlers_http.NewServer(ctx, config.HTTPServerConfig, service)
	if err != nil {
		log.Fatal().Err(err).
			Msg("main: unable to create http server")
	}

	// setup http server
	if err := httpServer.Setup(ctx); err != nil {
		log.Fatal().Err(err).
			Msg("main: unable to setup http server")
	}

	// start http server
	if err := httpServer.Start(ctx); err != nil {
		log.Fatal().Err(err).
			Msg("main: unable to start http server")
	}

	<-signals
	cancel()

	// stop http server
	if err := httpServer.Stop(ctx); err != nil {
		log.Fatal().Err(err).
			Msg("main: unable to stop http server")
	}

	os.Exit(0)
}
