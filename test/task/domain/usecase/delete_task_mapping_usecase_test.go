package usecase_test

import (
	"errors"
	"testing"

	sharedException "github.com/gerarc/tireg/internal/shared/domain/exception"

	"github.com/gerarc/tireg/internal/task/domain/exception"
	"github.com/gerarc/tireg/internal/task/domain/usecase"
)

func TestDelete_Succeeds_WhenTaskMappingExists(t *testing.T) {
	var calledID, calledOwnerID string
	commandRepository := &stubTaskMappingCommandRepository{
		deleteByIDAndOwnerFunc: func(id string, ownerID string) error {
			calledID, calledOwnerID = id, ownerID
			return nil
		},
	}

	deleteTaskMappingUseCase := usecase.NewDeleteTaskMappingUseCase(commandRepository)

	if err := deleteTaskMappingUseCase.Delete("owner-id", "mapping-id"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if calledID != "mapping-id" || calledOwnerID != "owner-id" {
		t.Fatalf("expected delete to be scoped by id and owner, got id=%q ownerID=%q", calledID, calledOwnerID)
	}
}

func TestDelete_ReturnsNotFoundError_WhenNotOwnedByOwner(t *testing.T) {
	commandRepository := &stubTaskMappingCommandRepository{
		deleteByIDAndOwnerFunc: func(id string, ownerID string) error {
			return exception.NewTaskMappingNotFoundError()
		},
	}

	deleteTaskMappingUseCase := usecase.NewDeleteTaskMappingUseCase(commandRepository)

	err := deleteTaskMappingUseCase.Delete("owner-id", "someone-elses-id")

	var domainError *sharedException.DomainError
	if !errors.As(err, &domainError) || domainError.Code != 404 {
		t.Fatalf("expected a 404 DomainError, got %v", err)
	}
}
