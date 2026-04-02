package interceptors

import (
	"context"

	obsctx "github.com/onlyizi/onlyizi-go/observability/context"
	"github.com/onlyizi/onlyizi-go/observability/logs"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func TracingInterceptor() grpc.UnaryServerInterceptor {
	tracer := otel.Tracer("grpc")

	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		ctx, span := tracer.Start(ctx, info.FullMethod)
		defer span.End()

		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			if requestIDs := md.Get("x-request-id"); len(requestIDs) > 0 {

				requestID := requestIDs[0]

				ctx = obsctx.WithRequestID(ctx, requestID)

				logger := logs.L().With(
					logs.RequestID(requestID),
				)

				ctx = logs.WithLogger(ctx, logger)
			}
		}

		return handler(ctx, req)
	}
}

func TracingClientInterceptor() grpc.UnaryClientInterceptor {
	tracer := otel.Tracer("grpc-client")

	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		ctx, span := tracer.Start(ctx, method)
		defer span.End()

		md := metadata.New(nil)

		if requestID, ok := obsctx.GetRequestID(ctx); ok {
			md.Set("x-request-id", requestID)
		}

		ctx = metadata.NewOutgoingContext(ctx, md)

		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
