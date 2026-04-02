package interceptors

import (
	"context"
	"time"

	"github.com/onlyizi/onlyizi-go/observability/logs"
	"google.golang.org/grpc"
)

func LoggingInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {

		start := time.Now()

		logger := logs.FromContext(ctx)

		resp, err := handler(ctx, req)

		duration := time.Since(start)

		if err != nil {
			logger.Error("grpc request failed",
				logs.Field("method", info.FullMethod),
				logs.Duration(duration),
				logs.Err(err),
			)
		} else {
			logger.Info("grpc request completed",
				logs.Field("method", info.FullMethod),
				logs.Duration(duration),
			)
		}

		return resp, err
	}
}

func LoggingClientInterceptor() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {

		start := time.Now()

		logger := logs.FromContext(ctx)

		err := invoker(ctx, method, req, reply, cc, opts...)

		duration := time.Since(start)

		if err != nil {
			logger.Error("grpc client request failed",
				logs.Field("method", method),
				logs.Duration(duration),
				logs.Err(err),
			)
		} else {
			logger.Info("grpc client request completed",
				logs.Field("method", method),
				logs.Duration(duration),
			)
		}

		return err
	}
}
