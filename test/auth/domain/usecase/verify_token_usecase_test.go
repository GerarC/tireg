package usecase_test

import (
	"errors"
	"testing"

	sharedException "github.com/gerarc/tireg/internal/shared/domain/exception"

	"github.com/gerarc/tireg/internal/auth/domain/model"
	"github.com/gerarc/tireg/internal/auth/domain/usecase"
)

type stubTokenVerifier struct {
	verifyFunc func(string) (model.AuthenticatedUser, error)
}

func (stub *stubTokenVerifier) Verify(token string) (model.AuthenticatedUser, error) {
	return stub.verifyFunc(token)
}

func TestVerify_ReturnsAuthenticatedUser_WhenTokenValid(t *testing.T) {
	tokenVerifier := &stubTokenVerifier{
		verifyFunc: func(token string) (model.AuthenticatedUser, error) {
			return model.AuthenticatedUser{ID: "user-id", Username: "ada"}, nil
		},
	}

	verifyTokenUseCase := usecase.NewVerifyTokenUseCase(tokenVerifier)

	authenticatedUser, err := verifyTokenUseCase.Verify("valid-token")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if authenticatedUser.ID != "user-id" || authenticatedUser.Username != "ada" {
		t.Fatalf("unexpected authenticated user: %+v", authenticatedUser)
	}
}

func TestVerify_ReturnsInvalidCredentials_WhenTokenVerifierFails(t *testing.T) {
	tokenVerifier := &stubTokenVerifier{
		verifyFunc: func(token string) (model.AuthenticatedUser, error) {
			return model.AuthenticatedUser{}, errors.New("malformed token")
		},
	}

	verifyTokenUseCase := usecase.NewVerifyTokenUseCase(tokenVerifier)

	_, err := verifyTokenUseCase.Verify("bad-token")

	var domainError *sharedException.DomainError
	if !errors.As(err, &domainError) || domainError.Code != 401 {
		t.Fatalf("expected a 401 DomainError, got %v", err)
	}
}
