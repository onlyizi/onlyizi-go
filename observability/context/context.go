package context

import "context"

type ctxKey string

const (
	RequestIDKey ctxKey = "request_id"
	UserIDKey    ctxKey = "user_id"
)

func WithRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, RequestIDKey, requestID)
}

func GetRequestID(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(RequestIDKey).(string)
	return id, ok
}
