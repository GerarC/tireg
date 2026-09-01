package usecase_test

import (
	"errors"
	"testing"

	sharedException "github.com/gerarc/tireg/internal/shared/domain/exception"

	"github.com/gerarc/tireg/internal/user/domain/model"
	"github.com/gerarc/tireg/internal/user/domain/usecase"
)

type stubUserCommandRepository struct {
	saveFunc func(model.User) (model.User, error)
}

func (stub *stubUserCommandRepository) Save(user model.User) (model.User, error) {
	return stub.saveFunc(user)
}

type stubUserQueryRepository struct {
	existsByUsername    bool
	existsByUsernameErr error
	existsByEmail       bool
	existsByEmailErr    error
	findByIdentifierErr error
}

func (stub *stubUserQueryRepository) ExistsByUsername(username string) (bool, error) {
	return stub.existsByUsername, stub.existsByUsernameErr
}

func (stub *stubUserQueryRepository) ExistsByEmail(email string) (bool, error) {
	return stub.existsByEmail, stub.existsByEmailErr
}

func (stub *stubUserQueryRepository) FindByIdentifier(identifier string) (model.User, error) {
	if stub.findByIdentifierErr != nil {
		return model.User{}, stub.findByIdentifierErr
	}

	return model.User{ID: "user-id", Username: identifier}, nil
}

type stubPasswordHasher struct {
	hashFunc func(string) (string, error)
}

func (stub *stubPasswordHasher) Hash(password string) (string, error) {
	return stub.hashFunc(password)
}

func (stub *stubPasswordHasher) Verify(hash string, password string) bool {
	return hash == password
}

func validRegistration() model.UserRegistration {
	return model.UserRegistration{
		FirstName: "Ada",
		LastName:  "Lovelace",
		Username:  "ada",
		Email:     "ada@example.com",
		Password:  "super-secret",
	}
}

func TestRegister_Succeeds_WhenUsernameAndEmailAreFree(t *testing.T) {
	commandRepository := &stubUserCommandRepository{
		saveFunc: func(user model.User) (model.User, error) {
			user.ID = "generated-id"
			return user, nil
		},
	}
	queryRepository := &stubUserQueryRepository{}
	hasher := &stubPasswordHasher{hashFunc: func(password string) (string, error) { return "hashed-" + password, nil }}

	registerUserUseCase := usecase.NewRegisterUserUseCase(commandRepository, queryRepository, hasher)

	user, err := registerUserUseCase.Register(validRegistration())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if user.ID != "generated-id" {
		t.Fatalf("expected generated id, got %q", user.ID)
	}

	if user.PasswordHash != "hashed-super-secret" {
		t.Fatalf("expected hashed password, got %q", user.PasswordHash)
	}
}

func TestRegister_ReturnsValidationError_WithAllAccumulatedDetails(t *testing.T) {
	commandRepository := &stubUserCommandRepository{}
	queryRepository := &stubUserQueryRepository{}
	hasher := &stubPasswordHasher{hashFunc: func(password string) (string, error) { return password, nil }}

	registerUserUseCase := usecase.NewRegisterUserUseCase(commandRepository, queryRepository, hasher)

	registration := validRegistration()
	registration.Username = "ab"
	registration.Password = "short"

	_, err := registerUserUseCase.Register(registration)

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

func TestRegister_ReturnsAlreadyTakenError_WithBothConflictsAccumulated(t *testing.T) {
	commandRepository := &stubUserCommandRepository{}
	queryRepository := &stubUserQueryRepository{existsByUsername: true, existsByEmail: true}
	hasher := &stubPasswordHasher{hashFunc: func(password string) (string, error) { return password, nil }}

	registerUserUseCase := usecase.NewRegisterUserUseCase(commandRepository, queryRepository, hasher)

	_, err := registerUserUseCase.Register(validRegistration())

	var domainError *sharedException.DomainError
	if !errors.As(err, &domainError) {
		t.Fatalf("expected a DomainError, got %v", err)
	}

	if domainError.Code != 409 {
		t.Fatalf("expected code 409, got %d", domainError.Code)
	}

	if len(domainError.Details) != 2 {
		t.Fatalf("expected 2 conflict details, got %d: %v", len(domainError.Details), domainError.Details)
	}
}

func TestRegister_ReturnsAlreadyTakenError_WithOnlyUsernameConflict(t *testing.T) {
	commandRepository := &stubUserCommandRepository{}
	queryRepository := &stubUserQueryRepository{existsByUsername: true, existsByEmail: false}
	hasher := &stubPasswordHasher{hashFunc: func(password string) (string, error) { return password, nil }}

	registerUserUseCase := usecase.NewRegisterUserUseCase(commandRepository, queryRepository, hasher)

	_, err := registerUserUseCase.Register(validRegistration())

	var domainError *sharedException.DomainError
	if !errors.As(err, &domainError) {
		t.Fatalf("expected a DomainError, got %v", err)
	}

	if len(domainError.Details) != 1 {
		t.Fatalf("expected 1 conflict detail, got %d: %v", len(domainError.Details), domainError.Details)
	}
}
