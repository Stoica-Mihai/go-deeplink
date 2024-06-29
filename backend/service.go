package main

import (
	"go-deeplink/backend/router"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	IP   = "0.0.0.0"
	PORT = "9000"
)

type Service struct {
	IP      string
	Port    string
	Router  *echo.Echo
	Logging echo.MiddlewareFunc
}

func NewService() *Service {
	return &Service{
		IP:      IP,
		Port:    PORT,
		Router:  router.GetRouter(),
		Logging: LoggingMiddleware(),
	}
}

func (s *Service) IPAndPort() string {
	return s.IP + ":" + s.Port
}

func (s *Service) Run(logging bool) {
	log.Printf("Listen on http://%s\n", s.IPAndPort())
	if logging {
		s.Router.Use(s.Logging)
	}
	s.Router.Logger.Fatal(s.Router.Start(s.IPAndPort()))
}

func LoggingMiddleware() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} [${method}] ${uri} ${status} ${error}\n",
	})
}
