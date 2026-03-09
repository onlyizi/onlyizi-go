package metrics

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	metricapi "go.opentelemetry.io/otel/metric"
)

var (
	httpRequestsTotal metricapi.Int64Counter
	httpDuration      metricapi.Float64Histogram
)

func InitHTTP(serviceName string) error {
	meter := Meter(serviceName)

	var err error

	httpRequestsTotal, err = meter.Int64Counter(
		"http_requests_total",
		metricapi.WithDescription("Total number of HTTP requests"),
	)
	if err != nil {
		return err
	}

	httpDuration, err = meter.Float64Histogram(
		"http_request_duration_ms",
		metricapi.WithDescription("Duration of HTTP requests in milliseconds"),
	)
	if err != nil {
		return err
	}

	return nil
}

func RecordHTTPRequest(
	ctx context.Context,
	method string,
	path string,
	status int,
	duration float64,
) {

	attrs := []attribute.KeyValue{
		attribute.String("method", method),
		attribute.String("path", path),
		attribute.Int("status", status),
	}

	httpRequestsTotal.Add(ctx, 1, metricapi.WithAttributes(attrs...))

	httpDuration.Record(ctx, duration, metricapi.WithAttributes(attrs...))
}
