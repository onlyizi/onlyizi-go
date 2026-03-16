package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/onlyizi/onlyizi-go/app"
	"github.com/onlyizi/onlyizi-go/config"
	"github.com/onlyizi/onlyizi-go/observability/logs"
)

type Service struct {
	cfg config.Redis
}

func New() app.Service {
	return &Service{
		cfg: config.RedisConfig(),
	}
}

func (s *Service) Name() string {
	return "redis"
}

func (s *Service) Start() error {

	addr := fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.Port)

	rdb := redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     s.cfg.Password,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		return err
	}

	client = rdb

	logs.L().Info(
		"redis connected",
		logs.Component("redis"),
		logs.Field("host", s.cfg.Host),
		logs.Field("port", s.cfg.Port),
	)

	return nil
}

func (s *Service) Shutdown(ctx context.Context) error {

	if client == nil {
		return nil
	}

	logs.L().Info(
		"redis shutting down",
		logs.Component("redis"),
	)

	return client.Close()
}
