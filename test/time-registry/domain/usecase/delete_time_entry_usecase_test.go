package usecase_test

import (
	"errors"
	"testing"

	sharedException "github.com/gerarc/tireg/internal/shared/domain/exception"

	"github.com/gerarc/tireg/internal/time-registry/domain/exception"
	"github.com/gerarc/tireg/internal/time-registry/domain/usecase"
)

func TestDelete_Succeeds_WhenTimeEntryExists(t *testing.T) {
	var calledID, calledOwnerID string
	commandRepository := &stubTimeEntryCommandRepository{
		deleteByIDAndOwnerFunc: func(id string, ownerID string) error {
			calledID, calledOwnerID = id, ownerID
			return nil
		},
	}

	deleteTimeEntryUseCase := usecase.NewDeleteTimeEntryUseCase(commandRepository)

	if err := deleteTimeEntryUseCase.Delete("owner-id", "entry-id"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if calledID != "entry-id" || calledOwnerID != "owner-id" {
		t.Fatalf("expected delete to be scoped by id and owner, got id=%q ownerID=%q", calledID, calledOwnerID)
	}
}

func TestDelete_ReturnsNotFoundError_WhenNotOwnedByOwner(t *testing.T) {
	commandRepository := &stubTimeEntryCommandRepository{
		deleteByIDAndOwnerFunc: func(id string, ownerID string) error {
			return exception.NewTimeEntryNotFoundError()
		},
	}

	deleteTimeEntryUseCase := usecase.NewDeleteTimeEntryUseCase(commandRepository)

	err := deleteTimeEntryUseCase.Delete("owner-id", "someone-elses-id")

	var domainError *sharedException.DomainError
	if !errors.As(err, &domainError) || domainError.Code != 404 {
		t.Fatalf("expected a 404 DomainError, got %v", err)
	}
}
