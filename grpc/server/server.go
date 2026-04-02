package grpcServer

import (
	"context"
	"net"

	"github.com/onlyizi/onlyizi-go/grpc/interceptors"
	"github.com/onlyizi/onlyizi-go/observability/logs"
	"google.golang.org/grpc"
)

type RegisterService func(*grpc.Server)

type Server struct {
	name     string
	addr     string
	server   *grpc.Server
	services []RegisterService
}

func NewServer(name, addr string, services ...RegisterService) *Server {
	return &Server{
		name:     name,
		addr:     addr,
		services: services,
	}
}

func (s *Server) Name() string {
	return s.name
}

func (s *Server) Start() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}

	s.server = grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptors.LoggingInterceptor(),
			interceptors.TracingInterceptor(),
			interceptors.MetricsInterceptor(),
		),
	)

	for _, register := range s.services {
		register(s.server)
	}

	logs.L().Info("grpc server starting",
		logs.Component("grpc"),
		logs.Field("name", s.name),
		logs.Field("addr", s.addr),
	)

	go func() {
		if err := s.server.Serve(lis); err != nil {
			logs.L().Error("grpc server stopped", logs.Err(err))
		}
	}()

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	if s.server == nil {
		return nil
	}

	logs.L().Info("grpc server shutting down",
		logs.Component("grpc"),
		logs.Field("name", s.name),
	)

	s.server.GracefulStop()
	return nil
}
