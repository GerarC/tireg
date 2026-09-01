package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	authException "github.com/gerarc/tireg/internal/auth/domain/exception"
	authModel "github.com/gerarc/tireg/internal/auth/domain/model"

	"github.com/gerarc/tireg/internal/shared/infrastructure/in/rest/middleware"
)

type stubVerifyTokenUseCase struct {
	verifyFunc func(string) (authModel.AuthenticatedUser, error)
}

func (stub *stubVerifyTokenUseCase) Verify(token string) (authModel.AuthenticatedUser, error) {
	return stub.verifyFunc(token)
}

func TestWrap_Returns401_WhenTokenInvalid(t *testing.T) {
	verifyTokenUseCase := &stubVerifyTokenUseCase{
		verifyFunc: func(token string) (authModel.AuthenticatedUser, error) {
			return authModel.AuthenticatedUser{}, authException.NewInvalidCredentialsError()
		},
	}
	requireAuthMiddleware := middleware.NewRequireAuthMiddleware(verifyTokenUseCase)

	called := false
	wrapped := requireAuthMiddleware.Wrap(func(w http.ResponseWriter, r *http.Request) {
		called = true
	})

	request := httptest.NewRequest(http.MethodGet, "/protected", nil)
	responseRecorder := httptest.NewRecorder()

	wrapped(responseRecorder, request)

	if responseRecorder.Code != http.StatusUnauthorized {
		t.Fatalf("expected status %d, got %d", http.StatusUnauthorized, responseRecorder.Code)
	}

	if called {
		t.Fatalf("expected the wrapped handler not to be called")
	}
}

func TestWrap_CallsNextWithAuthenticatedUserInContext_WhenTokenValid(t *testing.T) {
	verifyTokenUseCase := &stubVerifyTokenUseCase{
		verifyFunc: func(token string) (authModel.AuthenticatedUser, error) {
			if token != "valid-token" {
				t.Fatalf("expected token %q, got %q", "valid-token", token)
			}
			return authModel.AuthenticatedUser{ID: "user-id", Username: "ada"}, nil
		},
	}
	requireAuthMiddleware := middleware.NewRequireAuthMiddleware(verifyTokenUseCase)

	var contextUser authModel.AuthenticatedUser
	var contextUserFound bool
	wrapped := requireAuthMiddleware.Wrap(func(w http.ResponseWriter, r *http.Request) {
		contextUser, contextUserFound = middleware.AuthenticatedUserFromContext(r.Context())
		w.WriteHeader(http.StatusOK)
	})

	request := httptest.NewRequest(http.MethodGet, "/protected", nil)
	request.Header.Set("Authorization", "Bearer valid-token")
	responseRecorder := httptest.NewRecorder()

	wrapped(responseRecorder, request)

	if responseRecorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, responseRecorder.Code)
	}

	if !contextUserFound {
		t.Fatalf("expected authenticated user to be present in the request context")
	}

	if contextUser.ID != "user-id" || contextUser.Username != "ada" {
		t.Fatalf("unexpected authenticated user in context: %+v", contextUser)
	}
}
