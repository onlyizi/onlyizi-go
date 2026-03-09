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

	if err := metrics.Init(); err != nil {
		return err
	}

	if err := metrics.InitHTTP(cfg.ServiceName); err != nil {
		return err
	}

	if err := tracing.Init(cfg.ServiceName); err != nil {
		return err
	}

	return nil
}

func Shutdown(ctx context.Context) error {
	return tracing.Shutdown(ctx)
}
