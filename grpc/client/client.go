package grpcClient

import (
	"context"

	"github.com/onlyizi/onlyizi-go/grpc/interceptors"
	"github.com/onlyizi/onlyizi-go/observability/logs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	name    string
	address string
	conn    *grpc.ClientConn
}

func NewClient(name, address string) *Client {
	return &Client{
		name:    name,
		address: address,
	}
}

func (c *Client) Name() string {
	return c.name
}

func (c *Client) Start() error {
	conn, err := grpc.Dial(
		c.address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(
			interceptors.LoggingClientInterceptor(),
			interceptors.TracingClientInterceptor(),
			interceptors.MetricsClientInterceptor(),
		),
	)

	if err != nil {
		logs.L().Error("grpc dial failed", logs.Err(err))
		return nil
	}

	c.conn = conn

	logs.L().Info("grpc client ready",
		logs.Component("grpc"),
		logs.Field("name", c.name),
		logs.Field("address", c.address),
	)

	return nil
}

func (c *Client) Shutdown(ctx context.Context) error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

func (c *Client) Conn() *grpc.ClientConn {
	return c.conn
}
