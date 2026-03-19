package http

import (
	"context"
	stdHttp "net/http"

	"github.com/onlyizi/onlyizi-go/http/middlewares"
	"github.com/onlyizi/onlyizi-go/observability/logs"
)

type Server struct {
	name   string
	addr   string
	cors   middlewares.CORSConfig
	routes []RegisterRoutes
	server *stdHttp.Server
}

func NewServer(
	name string,
	addr string,
	cors middlewares.CORSConfig,
	routes ...RegisterRoutes,
) *Server {
	logs.L().Info(
		"http server created",
		logs.Component("http"),
		logs.Field("name", name),
		logs.Field("addr", addr),
	)

	return &Server{
		name:   name,
		addr:   addr,
		cors:   cors,
		routes: routes,
	}
}

func (s *Server) Name() string {
	return s.name
}

func (s *Server) Start() error {
	router := NewRouter(s.cors, s.routes...)

	s.server = &stdHttp.Server{
		Addr:    s.addr,
		Handler: router,
	}

	serverAddr := "localhost" + s.addr

	logs.L().Info(
		"http server starting",
		logs.Component("http"),
		logs.Field("name", s.name),
		logs.Field("addr", serverAddr),
	)

	err := s.server.ListenAndServe()
	if err != nil && err != stdHttp.ErrServerClosed {
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
	if s.server == nil {
		return nil
	}

	logs.L().Info(
		"http server shutting down",
		logs.Component("http"),
		logs.Field("name", s.name),
	)

	return s.server.Shutdown(ctx)
}
