package observability

import (
	"context"

	"github.com/onlyizi/onlyizi-go/observability/logs"
	"github.com/onlyizi/onlyizi-go/observability/metrics"
	"github.com/onlyizi/onlyizi-go/observability/tracing"
)

func Init(cfg Config) error {

	if err := logs.Init(logs.Config{
		Service:     cfg.ServiceName,
		Environment: logs.Environment(cfg.Environment),
		Version:     cfg.Version,
	}); err != nil {
		return err
	}

	logs.L().Info(
		"observability initializing",
		logs.Field("service", cfg.ServiceName),
		logs.Field("environment", cfg.Environment),
		logs.Field("version", cfg.Version),
	)

	if err := metrics.Init(); err != nil {
		logs.L().Error("metrics initialization failed", logs.Err(err))
		return err
	}

	logs.L().Info("metrics initialized")

	if err := metrics.InitHTTP(cfg.ServiceName); err != nil {
		logs.L().Error("http metrics initialization failed", logs.Err(err))
		return err
	}

	logs.L().Info("http metrics initialized")

	if err := tracing.Init(cfg.ServiceName); err != nil {
		logs.L().Error("tracing initialization failed", logs.Err(err))
		return err
	}

	logs.L().Info("tracing initialized")

	logs.L().Info("observability initialized successfully")

	return nil
}

func Shutdown(ctx context.Context) error {

	logs.L().Info("observability shutting down")

	return tracing.Shutdown(ctx)
}
