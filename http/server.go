package http

import (
	"context"
	"net/http"

	"github.com/onlyizi/onlyizi-go/observability/logs"
)

type Server struct {
	name   string
	server *http.Server
}

func NewServer(name, addr string, routes ...RegisterRoutes) *Server {
	router := NewRouter(routes...)

	s := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	return &Server{
		name:   name,
		server: s,
	}
}

func (s *Server) Name() string {
	return s.name
}

func (s *Server) Start() error {
	serverAddr := "localhost" + s.server.Addr

	logs.L().Info("Http server starting: " + serverAddr)

	err := s.server.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	logs.L().Info("Http server shutting down")

	return s.server.Shutdown(ctx)
}
