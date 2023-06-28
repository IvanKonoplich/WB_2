package controller

import (
	"net/http"
	"time"
)

// Server - это структура http сервера
type Server struct {
	Server http.Server
}

// InitServer запускает сервер на переданном порту
func (s *Server) InitServer(port string) error {
	s.Server = http.Server{
		Addr:           ":" + port,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.Server.ListenAndServe()
}
