package middlewares

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/onlyizi/onlyizi-go/observability/logs"
	"github.com/onlyizi/onlyizi-go/observability/metrics"
)

func ObservabilityMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		reqID := uuid.NewString()

		logger := logs.L().With(
			logs.RequestID(reqID),
		)

		ctx := logs.WithLogger(r.Context(), logger)

		w.Header().Set("X-Request-ID", reqID)

		rw := newResponseWriter(w)

		clientIP := clientIP(r)
		userAgent := r.UserAgent()
		contentLength := r.ContentLength

		logger.Info("request started",
			logs.Method(r.Method),
			logs.Path(r.URL.Path),
			logs.ClientIP(clientIP),
			logs.UserAgent(userAgent),
			logs.ContentLength(contentLength),
		)

		next.ServeHTTP(rw, r.WithContext(ctx))

		duration := time.Since(start)

		// 🔹 MÉTRICAS
		metrics.RecordHTTPRequest(
			ctx,
			r.Method,
			r.URL.Path,
			rw.status,
			float64(duration.Milliseconds()),
		)

		logger.Info("request completed",
			logs.Method(r.Method),
			logs.Path(r.URL.Path),
			logs.Status(rw.status),
			logs.Duration(duration),
			logs.Bytes(rw.bytes),
			logs.ClientIP(clientIP),
		)
	})
}

func clientIP(r *http.Request) string {

	ip := r.Header.Get("X-Forwarded-For")
	if ip != "" {
		return ip
	}

	ip = r.Header.Get("X-Real-IP")
	if ip != "" {
		return ip
	}

	return r.RemoteAddr
}
