package metadata

import (
	"context"

	"github.com/onlyizi/onlyizi-go/observability/logs"
	"google.golang.org/grpc/metadata"
)

const (
	RequestIDKey = "x-request-id"
)

func FromContext(ctx context.Context) metadata.MD {
	logger := logs.FromContext(ctx)

	fields := logger.Core().Enabled(0)

	_ = fields

	return metadata.New(nil)
}

func Inject(ctx context.Context, md metadata.MD) context.Context {
	return metadata.NewOutgoingContext(ctx, md)
}

func Extract(ctx context.Context) metadata.MD {
	md, _ := metadata.FromIncomingContext(ctx)
	return md
}
