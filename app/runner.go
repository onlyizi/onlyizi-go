package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/onlyizi/onlyizi-go/observability/logs"
)

func Run(services ...Service) error {

	for _, s := range services {
		go func(s Service) {
			logs.L().Info("starting service", logs.Component(s.Name()))

			if err := s.Start(); err != nil {
				logs.L().Error("service failed",
					logs.Component(s.Name()),
					logs.Err(err),
				)
			}
		}(s)
	}

	quit := make(chan os.Signal, 1)

	signal.Notify(
		quit,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	<-quit

	logs.L().Info("shutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for _, s := range services {
		logs.L().Info("shutting down service", logs.Component(s.Name()))

		if err := s.Shutdown(ctx); err != nil {
			logs.L().Error("shutdown failed",
				logs.Component(s.Name()),
				logs.Err(err),
			)
		}
	}

	return nil
}
