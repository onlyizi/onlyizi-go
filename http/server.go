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

	logs.L().Info(
		"http server created",
		logs.Component("http"),
		logs.Field("name", name),
		logs.Field("addr", addr),
	)

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

	logs.L().Info(
		"http server starting",
		logs.Component("http"),
		logs.Field("name", s.name),
		logs.Field("addr", serverAddr),
	)

	err := s.server.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {

		logs.L().Error(
			"http server failed",
			logs.Component("http"),
			logs.Err(err),
		)

		return err
	}

	logs.L().Info(
		"http server stopped",
		logs.Component("http"),
	)

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {

	logs.L().Info(
		"http server shutting down",
		logs.Component("http"),
		logs.Field("name", s.name),
	)

	return s.server.Shutdown(ctx)
}
