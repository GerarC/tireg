package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gerarc/tireg/internal/shared/application/utils/logger"
	"github.com/gerarc/tireg/internal/shared/infrastructure/in/rest/middleware"
)

type stubLogger struct {
	infoFunc func(string, ...any)
}

func (stub *stubLogger) Info(msg string, args ...any) {
	if stub.infoFunc != nil {
		stub.infoFunc(msg, args...)
	}
}

func (stub *stubLogger) Warn(msg string, args ...any) {}

func (stub *stubLogger) Error(msg string, err error, args ...any) {}

func (stub *stubLogger) Debug(msg string, args ...any) {}

func (stub *stubLogger) With(args ...any) logger.Logger { return stub }

func TestWrap_LogsMethodPathStatusAndDuration(t *testing.T) {
	var loggedArgs []any
	stub := &stubLogger{
		infoFunc: func(msg string, args ...any) {
			loggedArgs = args
		},
	}
	requestLoggingMiddleware := middleware.NewRequestLoggingMiddleware(stub)

	wrapped := requestLoggingMiddleware.Wrap(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
	}))

	request := httptest.NewRequest(http.MethodPost, "/api/v1/users", nil)
	responseRecorder := httptest.NewRecorder()

	wrapped.ServeHTTP(responseRecorder, request)

	if responseRecorder.Code != http.StatusCreated {
		t.Fatalf("expected wrapped handler's status to pass through, got %d", responseRecorder.Code)
	}

	found := map[string]any{}
	for i := 0; i+1 < len(loggedArgs); i += 2 {
		key, ok := loggedArgs[i].(string)
		if !ok {
			continue
		}
		found[key] = loggedArgs[i+1]
	}

	if found["method"] != http.MethodPost {
		t.Fatalf("expected logged method %q, got %v", http.MethodPost, found["method"])
	}

	if found["path"] != "/api/v1/users" {
		t.Fatalf("expected logged path %q, got %v", "/api/v1/users", found["path"])
	}

	if found["status"] != http.StatusCreated {
		t.Fatalf("expected logged status %d, got %v", http.StatusCreated, found["status"])
	}
}

func TestWrap_DefaultsStatusTo200_WhenHandlerNeverCallsWriteHeader(t *testing.T) {
	var loggedArgs []any
	stub := &stubLogger{
		infoFunc: func(msg string, args ...any) {
			loggedArgs = args
		},
	}
	requestLoggingMiddleware := middleware.NewRequestLoggingMiddleware(stub)

	wrapped := requestLoggingMiddleware.Wrap(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	}))

	request := httptest.NewRequest(http.MethodGet, "/api/v1/health", nil)
	responseRecorder := httptest.NewRecorder()

	wrapped.ServeHTTP(responseRecorder, request)

	found := map[string]any{}
	for i := 0; i+1 < len(loggedArgs); i += 2 {
		key, ok := loggedArgs[i].(string)
		if !ok {
			continue
		}
		found[key] = loggedArgs[i+1]
	}

	if found["status"] != http.StatusOK {
		t.Fatalf("expected default logged status %d, got %v", http.StatusOK, found["status"])
	}
}
