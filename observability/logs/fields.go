package logs

import (
	"time"

	"go.uber.org/zap"
)

func RequestID(id string) zap.Field {
	return zap.String("request_id", id)
}

func TraceID(id string) zap.Field {
	return zap.String("trace_id", id)
}

func UserId(id string) zap.Field {
	return zap.String("user_id", id)
}

func Component(comp string) zap.Field {
	return zap.String("Component", comp)
}

func Field(key string, value any) zap.Field {
	return zap.Any(key, value)
}

func Err(err error) zap.Field {
	return zap.Error(err)
}

func Method(method string) zap.Field {
	return zap.String("method", method)
}

func Path(path string) zap.Field {
	return zap.String("path", path)
}

func Status(status int) zap.Field {
	return zap.Int("status", status)
}

func Duration(d time.Duration) zap.Field {
	return zap.Int64("duration_ms", d.Milliseconds())
}

func Bytes(n int) zap.Field {
	return zap.Int("bytes", n)
}

func ClientIP(ip string) zap.Field {
	return zap.String("client_ip", ip)
}

func UserAgent(agent string) zap.Field {
	return zap.String("user_agent", agent)
}

func ContentLength(length int64) zap.Field {
	return zap.Int64("content_length", length)
}

func ErrorCode(code string) zap.Field {
	return zap.String("error_code", code)
}
