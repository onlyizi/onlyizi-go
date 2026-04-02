package grpcClient

import (
	"context"
	"sync"

	"github.com/onlyizi/onlyizi-go/grpc/interceptors"
	"github.com/onlyizi/onlyizi-go/observability/logs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	name    string
	address string
	conn    *grpc.ClientConn
	mu      sync.Mutex
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
	logs.L().Info("grpc client initialized",
		logs.Field("name", c.name),
		logs.Field("address", c.address),
	)

	return nil
}

func (c *Client) getConn() (*grpc.ClientConn, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.conn != nil {
		return c.conn, nil
	}

	logs.L().Info("grpc dialing (lazy)...",
		logs.Field("name", c.name),
		logs.Field("address", c.address),
	)

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
		logs.L().Error("grpc dial failed",
			logs.Field("name", c.name),
			logs.Err(err),
		)
		return nil, err
	}

	c.conn = conn
	return conn, nil
}

func (c *Client) Conn() (*grpc.ClientConn, error) {
	return c.getConn()
}

func (c *Client) Shutdown(ctx context.Context) error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}
