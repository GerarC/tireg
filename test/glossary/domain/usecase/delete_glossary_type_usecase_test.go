package usecase_test

import (
	"errors"
	"testing"

	sharedException "github.com/gerarc/tireg/internal/shared/domain/exception"

	"github.com/gerarc/tireg/internal/glossary/domain/exception"
	"github.com/gerarc/tireg/internal/glossary/domain/usecase"
)

func TestDelete_Succeeds_WhenGlossaryTypeExists(t *testing.T) {
	var calledID, calledOwnerID string
	commandRepository := &stubGlossaryTypeCommandRepository{
		deleteByIDAndOwnerFunc: func(id string, ownerID string) error {
			calledID, calledOwnerID = id, ownerID
			return nil
		},
	}

	deleteGlossaryTypeUseCase := usecase.NewDeleteGlossaryTypeUseCase(commandRepository)

	if err := deleteGlossaryTypeUseCase.Delete("owner-id", "type-id"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if calledID != "type-id" || calledOwnerID != "owner-id" {
		t.Fatalf("expected delete to be scoped by id and owner, got id=%q ownerID=%q", calledID, calledOwnerID)
	}
}

func TestDelete_ReturnsNotFoundError_WhenRepositoryReportsNotFound(t *testing.T) {
	commandRepository := &stubGlossaryTypeCommandRepository{
		deleteByIDAndOwnerFunc: func(id string, ownerID string) error {
			return exception.NewGlossaryTypeNotFoundError()
		},
	}

	deleteGlossaryTypeUseCase := usecase.NewDeleteGlossaryTypeUseCase(commandRepository)

	err := deleteGlossaryTypeUseCase.Delete("owner-id", "someone-elses-id")

	var domainError *sharedException.DomainError
	if !errors.As(err, &domainError) || domainError.Code != 404 {
		t.Fatalf("expected a 404 DomainError, got %v", err)
	}
}
