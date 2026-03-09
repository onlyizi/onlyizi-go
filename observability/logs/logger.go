package logs

import "go.uber.org/zap"

var global *zap.Logger

func L() *zap.Logger {
	if global == nil {
		panic("logger not initialized")
	}
	return global
}
