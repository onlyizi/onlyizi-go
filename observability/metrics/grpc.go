package metrics

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/attribute"
	metricapi "go.opentelemetry.io/otel/metric"
)

var (
	rpcDuration metricapi.Float64Histogram
)

func InitRPC(serviceName string) error {
	meter := Meter(serviceName)

	var err error

	rpcDuration, err = meter.Float64Histogram(
		"grpc_request_duration_ms",
	)
	return err
}

func RecordRPC(
	ctx context.Context,
	duration time.Duration,
	attrs ...attribute.KeyValue,
) {
	rpcDuration.Record(
		ctx,
		float64(duration.Milliseconds()),
		metricapi.WithAttributes(attrs...),
	)
}
