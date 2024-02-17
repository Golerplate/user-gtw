package http

import (
	"context"
	"fmt"

	pkghttp "github.com/Golerplate/pkg/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
)

type httpServer struct {
	router *echo.Echo
	config pkghttp.Config
}

func NewServer(ctx context.Context, cfg pkghttp.Config) (Server, error) {
	return &httpServer{
		router: echo.New(),
		config: cfg,
	}, nil
}

func (s *httpServer) Setup(ctx context.Context) error {
	// setup middlewares
	s.router.Use(middleware.Logger())
	s.router.Use(middleware.Recover())
	s.router.Use(middleware.CORS())

	return nil
}

func (s *httpServer) Start(ctx context.Context) error {
	log.Info().
		Msg("handlers.http.httpServer.Start: Starting HTTP server...")

	return s.router.Start(fmt.Sprintf(":%d", s.config.Port))
}

func (s *httpServer) Stop(ctx context.Context) error {
	log.Info().
		Msg("handlers.http.httpServer.Stop: Stopping HTTP server...")

	return s.router.Shutdown(ctx)
}
