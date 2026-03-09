package logs

import "go.uber.org/zap"

func RequestID(id string) zap.Field {
	return zap.String("request_id", id)
}

func TraceID(id string) zap.Field {
	return zap.String("trace_id", id)
}

func UserId(id string) zap.Field {
	return zap.String("user_id", id)
}

func Err(err error) zap.Field {
	return zap.Error(err)
}
