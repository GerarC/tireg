package middleware

import (
	"net/http"
	"time"

	"github.com/gerarc/tireg/internal/shared/application/utils/logger"
)

type RequestLoggingMiddleware struct {
	logger logger.Logger
}

func NewRequestLoggingMiddleware(appLogger logger.Logger) *RequestLoggingMiddleware {
	return &RequestLoggingMiddleware{logger: appLogger}
}

func (requestLoggingMiddleware *RequestLoggingMiddleware) Wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		start := time.Now()
		recorder := &statusRecordingResponseWriter{ResponseWriter: responseWriter, statusCode: http.StatusOK}

		next.ServeHTTP(recorder, request)

		requestLoggingMiddleware.logger.Info(
			"request handled",
			"method", request.Method,
			"path", request.URL.Path,
			"status", recorder.statusCode,
			"duration_ms", time.Since(start).Milliseconds(),
		)
	})
}

type statusRecordingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (statusRecordingResponseWriter *statusRecordingResponseWriter) WriteHeader(statusCode int) {
	statusRecordingResponseWriter.statusCode = statusCode
	statusRecordingResponseWriter.ResponseWriter.WriteHeader(statusCode)
}
