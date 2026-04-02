package interceptors

import (
	"context"
	"time"

	"github.com/onlyizi/onlyizi-go/observability/metrics"
	"go.opentelemetry.io/otel/attribute"
	"google.golang.org/grpc"
)

func MetricsInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {

		start := time.Now()

		resp, err := handler(ctx, req)

		duration := time.Since(start)

		attrs := []attribute.KeyValue{
			attribute.String("method", info.FullMethod),
		}

		if err != nil {
			attrs = append(attrs, attribute.String("status", "error"))
		} else {
			attrs = append(attrs, attribute.String("status", "ok"))
		}

		metrics.RecordRPC(ctx, duration, attrs...)

		return resp, err
	}
}

func MetricsClientInterceptor() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {

		start := time.Now()

		err := invoker(ctx, method, req, reply, cc, opts...)

		duration := time.Since(start)

		attrs := []attribute.KeyValue{
			attribute.String("method", method),
		}

		if err != nil {
			attrs = append(attrs, attribute.String("status", "error"))
		} else {
			attrs = append(attrs, attribute.String("status", "ok"))
		}

		metrics.RecordRPC(ctx, duration, attrs...)

		return err
	}
}
