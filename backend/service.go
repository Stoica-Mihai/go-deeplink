package main

import (
	"fmt"
	"go-deeplink/backend/router"
	"net/http"
)

const (
	IP   = "0.0.0.0"
	PORT = "9000"
)

type Service struct {
	IP      string
	Port    string
	Router  *http.ServeMux
	Logging func(http.Handler) http.Handler
}

func (s *Service) IPAndPort() string {
	return s.IP + ":" + s.Port
}

func (s *Service) Run(logging bool) {
	fmt.Printf("Listen on http://%s\n", s.IPAndPort())
	var router http.Handler = s.Router
	if logging {
		router = s.Logging(s.Router)
	}
	http.ListenAndServe(s.IPAndPort(), router)
}

func NewService() *Service {
	return &Service{
		IP:      IP,
		Port:    PORT,
		Router:  router.GetRouter(),
		Logging: router.LoggingMiddleware,
	}
}
