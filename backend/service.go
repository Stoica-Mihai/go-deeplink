package main

import (
	"go-deeplink/backend/router"
	"log"
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
	log.Printf("Listen on http://%s\n", s.IPAndPort())
	var router http.Handler = s.Router
	if logging && s.Logging != nil {
		router = s.Logging(s.Router)
	} else {
		log.Println("Logging could not be enabled due to the service not having a logging middleware")
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
