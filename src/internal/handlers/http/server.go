package handlers_http

import (
	"context"
	"fmt"

	pkghttp "github.com/golerplate/pkg/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"

	"github.com/golerplate/user-gtw/internal/handlers"
	handlers_http_private_current_user_v1 "github.com/golerplate/user-gtw/internal/handlers/http/private/current-user/v1"
	handlers_http_private_users_v1 "github.com/golerplate/user-gtw/internal/handlers/http/private/users/v1"
	service_v1 "github.com/golerplate/user-gtw/internal/service/v1"
)

type httpServer struct {
	router  *echo.Echo
	config  pkghttp.HTTPServerConfig
	service *service_v1.Service
}

func NewServer(ctx context.Context, cfg pkghttp.HTTPServerConfig, service *service_v1.Service) (handlers.Server, error) {
	return &httpServer{
		router:  echo.New(),
		config:  cfg,
		service: service,
	}, nil
}

func (s *httpServer) Setup(ctx context.Context) error {
	log.Info().
		Msg("handlers.http.httpServer.Setup: Setting up HTTP server...")

	// setup handlers
	privateCurrentUserV1Handlers := handlers_http_private_current_user_v1.NewHandler(ctx, s.service)
	privateUsersV1Handlers := handlers_http_private_users_v1.NewHandler(ctx, s.service)

	// setup middlewares
	s.router.Use(middleware.Logger())
	s.router.Use(middleware.Recover())
	s.router.Use(middleware.CORS())

	// setup endpoints

	// private endpoints
	privateV1 := s.router.Group("/private/v1")

	// users related endpoints
	privateUsersV1 := privateV1.Group("/users")
	privateUsersV1.POST("/", privateUsersV1Handlers.CreateUser)
	privateUsersV1.GET("/:identifier", privateUsersV1Handlers.GetByIdentifier)

	// current-user related endpoints
	privateCurrentUserV1 := privateUsersV1.Group("/current-user")
	privateCurrentUserV1.GET("/", privateCurrentUserV1Handlers.Get)
	privateCurrentUserV1.PUT("/username", privateCurrentUserV1Handlers.UpdateUsername)

	return nil
}

func (s *httpServer) Start(ctx context.Context) error {
	log.Info().
		Uint16("port", s.config.Port).
		Msg("handlers.http.httpServer.Start: Starting HTTP server...")

	return s.router.Start(fmt.Sprintf(":%d", s.config.Port))
}

func (s *httpServer) Stop(ctx context.Context) error {
	log.Info().
		Msg("handlers.http.httpServer.Stop: Stopping HTTP server...")

	return s.router.Shutdown(ctx)
}
