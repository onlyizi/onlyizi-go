package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/onlyizi/onlyizi-go/app"
	"github.com/onlyizi/onlyizi-go/config"
	"github.com/onlyizi/onlyizi-go/observability/logs"
)

type Service struct{}

func New() app.Service {
	return &Service{}
}

func (s *Service) Name() string {
	return "postgres"
}

func (s *Service) Start() error {
	cfg := config.PostgresConfig()

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.DB,
	)

	conn, err := sql.Open("pgx", dsn)
	if err != nil {
		return err
	}

	conn.SetMaxOpenConns(25)
	conn.SetMaxIdleConns(10)
	conn.SetConnMaxLifetime(time.Hour)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := conn.PingContext(ctx); err != nil {
		return err
	}

	db = conn

	logs.L().Info(
		"postgres connected",
		logs.Component("postgres"),
		logs.Field("host", cfg.Host),
		logs.Field("port", cfg.Port),
		logs.Field("database", cfg.DB),
	)

	return nil
}

func (s *Service) Shutdown(ctx context.Context) error {
	if db == nil {
		return nil
	}

	logs.L().Info(
		"postgres shutting down",
		logs.Component("postgres"),
	)

	return db.Close()
}