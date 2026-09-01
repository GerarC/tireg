package usecase_test

import (
	"errors"
	"testing"

	"github.com/gerarc/tireg/internal/user/domain/usecase"
)

func TestFindByIdentifier_DelegatesToRepository(t *testing.T) {
	queryRepository := &stubUserQueryRepository{}

	findUserByIdentifierUseCase := usecase.NewFindUserByIdentifierUseCase(queryRepository)

	user, err := findUserByIdentifierUseCase.FindByIdentifier("ada")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if user.Username != "ada" {
		t.Fatalf("expected username %q, got %q", "ada", user.Username)
	}
}

func TestFindByIdentifier_PropagatesRepositoryError(t *testing.T) {
	repositoryErr := errors.New("not found")
	queryRepository := &stubUserQueryRepository{findByIdentifierErr: repositoryErr}

	findUserByIdentifierUseCase := usecase.NewFindUserByIdentifierUseCase(queryRepository)

	_, err := findUserByIdentifierUseCase.FindByIdentifier("missing")
	if !errors.Is(err, repositoryErr) {
		t.Fatalf("expected repository error to propagate, got %v", err)
	}
}
