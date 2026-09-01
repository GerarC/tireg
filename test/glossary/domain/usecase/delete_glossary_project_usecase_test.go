package usecase_test

import (
	"errors"
	"testing"

	sharedException "github.com/gerarc/tireg/internal/shared/domain/exception"

	"github.com/gerarc/tireg/internal/glossary/domain/exception"
	"github.com/gerarc/tireg/internal/glossary/domain/usecase"
)

func TestDelete_Succeeds_WhenGlossaryProjectExists(t *testing.T) {
	var calledID, calledOwnerID string
	commandRepository := &stubGlossaryProjectCommandRepository{
		deleteByIDAndOwnerFunc: func(id string, ownerID string) error {
			calledID, calledOwnerID = id, ownerID
			return nil
		},
	}

	deleteGlossaryProjectUseCase := usecase.NewDeleteGlossaryProjectUseCase(commandRepository)

	if err := deleteGlossaryProjectUseCase.Delete("owner-id", "project-id"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if calledID != "project-id" || calledOwnerID != "owner-id" {
		t.Fatalf("expected delete to be scoped by id and owner, got id=%q ownerID=%q", calledID, calledOwnerID)
	}
}

func TestDeleteProject_ReturnsNotFoundError_WhenRepositoryReportsNotFound(t *testing.T) {
	commandRepository := &stubGlossaryProjectCommandRepository{
		deleteByIDAndOwnerFunc: func(id string, ownerID string) error {
			return exception.NewGlossaryProjectNotFoundError()
		},
	}

	deleteGlossaryProjectUseCase := usecase.NewDeleteGlossaryProjectUseCase(commandRepository)

	err := deleteGlossaryProjectUseCase.Delete("owner-id", "someone-elses-id")

	var domainError *sharedException.DomainError
	if !errors.As(err, &domainError) || domainError.Code != 404 {
		t.Fatalf("expected a 404 DomainError, got %v", err)
	}
}
