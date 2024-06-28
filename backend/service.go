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
	IP     string
	Port   string
	Router http.Handler
}

func (s *Service) IPAndPort() string {
	return s.IP + ":" + s.Port
}

func (s *Service) Run() {
	fmt.Printf("Listen on http://%s\n", s.IPAndPort())
	http.ListenAndServe(s.IPAndPort(), s.Router)
}

func NewService() *Service {
	return &Service{
		IP:     IP,
		Port:   PORT,
		Router: router.GetRouter(),
	}
}
