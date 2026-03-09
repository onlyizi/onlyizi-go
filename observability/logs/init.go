package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Init(cfg Config) error {
	var logger *zap.Logger
	var err error

	if cfg.Environment == Development || cfg.Environment == "" {
		logger, err = newDevelopmentLogger(cfg)
	} else {
		logger, err = newProductionLogger(cfg)
	}

	if err != nil {
		return err
	}

	global = logger

	return nil
}

func newDevelopmentLogger(cfg Config) (*zap.Logger, error) {
	config := zap.NewDevelopmentConfig()

	logger, err := config.Build(
		zap.Fields(
			zap.String("service", cfg.Service),
			zap.String("environment", string(cfg.Environment)),
			zap.String("version", cfg.Version),
		),
	)

	if err != nil {
		return nil, err
	}

	return logger, nil
}

func newProductionLogger(cfg Config) (*zap.Logger, error) {
	config := zap.NewProductionConfig()

	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	logger, err := config.Build(
		zap.Fields(
			zap.String("service", cfg.Service),
			zap.String("environment", string(cfg.Environment)),
			zap.String("version", cfg.Version),
		),
	)

	if err != nil {
		return nil, err
	}

	return logger, nil
}
