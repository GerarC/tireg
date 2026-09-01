package usecase_test

import (
	"errors"
	"testing"

	sharedException "github.com/gerarc/tireg/internal/shared/domain/exception"

	userModel "github.com/gerarc/tireg/internal/user/domain/model"

	"github.com/gerarc/tireg/internal/auth/domain/model"
	"github.com/gerarc/tireg/internal/auth/domain/usecase"
)

type stubFindUserByIdentifierUseCase struct {
	findByIdentifierFunc func(string) (userModel.User, error)
}

func (stub *stubFindUserByIdentifierUseCase) FindByIdentifier(identifier string) (userModel.User, error) {
	return stub.findByIdentifierFunc(identifier)
}

type stubPasswordHasher struct {
	verifyFunc func(hash string, password string) bool
}

func (stub *stubPasswordHasher) Hash(password string) (string, error) {
	return password, nil
}

func (stub *stubPasswordHasher) Verify(hash string, password string) bool {
	return stub.verifyFunc(hash, password)
}

type stubTokenIssuer struct {
	issueFunc func(userID string, username string) (model.AccessToken, error)
}

func (stub *stubTokenIssuer) Issue(userID string, username string) (model.AccessToken, error) {
	return stub.issueFunc(userID, username)
}

func TestLogin_Succeeds_WhenCredentialsValid(t *testing.T) {
	findUserByIdentifierUseCase := &stubFindUserByIdentifierUseCase{
		findByIdentifierFunc: func(identifier string) (userModel.User, error) {
			return userModel.User{ID: "user-id", Username: "ada", PasswordHash: "hashed"}, nil
		},
	}
	passwordHasher := &stubPasswordHasher{verifyFunc: func(hash, password string) bool { return hash == "hashed" && password == "secret" }}
	tokenIssuer := &stubTokenIssuer{issueFunc: func(userID, username string) (model.AccessToken, error) {
		return model.AccessToken{Token: "jwt-token"}, nil
	}}

	loginUseCase := usecase.NewLoginUseCase(findUserByIdentifierUseCase, passwordHasher, tokenIssuer)

	accessToken, err := loginUseCase.Login(model.Credentials{Identifier: "ada", Password: "secret"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if accessToken.Token != "jwt-token" {
		t.Fatalf("expected token %q, got %q", "jwt-token", accessToken.Token)
	}
}

func TestLogin_ReturnsInvalidCredentials_WhenUserNotFound(t *testing.T) {
	findUserByIdentifierUseCase := &stubFindUserByIdentifierUseCase{
		findByIdentifierFunc: func(identifier string) (userModel.User, error) {
			return userModel.User{}, errors.New("not found")
		},
	}
	passwordHasher := &stubPasswordHasher{verifyFunc: func(hash, password string) bool { return true }}
	tokenIssuer := &stubTokenIssuer{}

	loginUseCase := usecase.NewLoginUseCase(findUserByIdentifierUseCase, passwordHasher, tokenIssuer)

	_, err := loginUseCase.Login(model.Credentials{Identifier: "ghost", Password: "secret"})

	var domainError *sharedException.DomainError
	if !errors.As(err, &domainError) || domainError.Code != 401 {
		t.Fatalf("expected a 401 DomainError, got %v", err)
	}
}

func TestLogin_ReturnsInvalidCredentials_WhenPasswordWrong(t *testing.T) {
	findUserByIdentifierUseCase := &stubFindUserByIdentifierUseCase{
		findByIdentifierFunc: func(identifier string) (userModel.User, error) {
			return userModel.User{ID: "user-id", Username: "ada", PasswordHash: "hashed"}, nil
		},
	}
	passwordHasher := &stubPasswordHasher{verifyFunc: func(hash, password string) bool { return false }}
	tokenIssuer := &stubTokenIssuer{}

	loginUseCase := usecase.NewLoginUseCase(findUserByIdentifierUseCase, passwordHasher, tokenIssuer)

	_, err := loginUseCase.Login(model.Credentials{Identifier: "ada", Password: "wrong"})

	var domainError *sharedException.DomainError
	if !errors.As(err, &domainError) || domainError.Code != 401 {
		t.Fatalf("expected a 401 DomainError, got %v", err)
	}
}

func TestLogin_PropagatesError_WhenTokenIssuerFails(t *testing.T) {
	tokenIssuerErr := errors.New("signing failed")
	findUserByIdentifierUseCase := &stubFindUserByIdentifierUseCase{
		findByIdentifierFunc: func(identifier string) (userModel.User, error) {
			return userModel.User{ID: "user-id", Username: "ada", PasswordHash: "hashed"}, nil
		},
	}
	passwordHasher := &stubPasswordHasher{verifyFunc: func(hash, password string) bool { return true }}
	tokenIssuer := &stubTokenIssuer{issueFunc: func(userID, username string) (model.AccessToken, error) {
		return model.AccessToken{}, tokenIssuerErr
	}}

	loginUseCase := usecase.NewLoginUseCase(findUserByIdentifierUseCase, passwordHasher, tokenIssuer)

	_, err := loginUseCase.Login(model.Credentials{Identifier: "ada", Password: "secret"})
	if !errors.Is(err, tokenIssuerErr) {
		t.Fatalf("expected token issuer error to propagate, got %v", err)
	}
}

func TestLogin_ReturnsValidationError_WithAllAccumulatedDetails_WhenFieldsEmpty(t *testing.T) {
	loginUseCase := usecase.NewLoginUseCase(&stubFindUserByIdentifierUseCase{}, &stubPasswordHasher{}, &stubTokenIssuer{})

	_, err := loginUseCase.Login(model.Credentials{Identifier: "", Password: ""})

	var domainError *sharedException.DomainError
	if !errors.As(err, &domainError) {
		t.Fatalf("expected a DomainError, got %v", err)
	}

	if domainError.Code != 400 {
		t.Fatalf("expected code 400, got %d", domainError.Code)
	}

	if len(domainError.Details) != 2 {
		t.Fatalf("expected 2 validation details, got %d: %v", len(domainError.Details), domainError.Details)
	}
}
