package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/onlyizi/onlyizi-go/observability/logs"
)

func Run(bootstrap []Service, runtime []Service) error {
	started := make([]Service, 0, len(bootstrap)+len(runtime))

	for _, s := range bootstrap {
		logs.L().Info("starting bootstrap service", logs.Component(s.Name()))

		if err := s.Start(); err != nil {
			logs.L().Error(
				"bootstrap service failed",
				logs.Component(s.Name()),
				logs.Err(err),
			)

			shutdownStarted(started)
			return err
		}

		started = append(started, s)
	}

	for _, s := range runtime {
		started = append(started, s)

		go func(s Service) {
			logs.L().Info("starting runtime service", logs.Component(s.Name()))

			if err := s.Start(); err != nil {
				logs.L().Error(
					"runtime service failed",
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

	for i := len(started) - 1; i >= 0; i-- {
		s := started[i]

		logs.L().Info("shutting down service", logs.Component(s.Name()))

		if err := s.Shutdown(ctx); err != nil {
			logs.L().Error(
				"shutdown failed",
				logs.Component(s.Name()),
				logs.Err(err),
			)
		}
	}

	return nil
}

func shutdownStarted(services []Service) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for i := len(services) - 1; i >= 0; i-- {
		_ = services[i].Shutdown(ctx)
	}
}
