package usecase_test

import (
	"errors"
	"testing"

	sharedException "github.com/gerarc/tireg/internal/shared/domain/exception"

	"github.com/gerarc/tireg/internal/time-registry/domain/exception"
	"github.com/gerarc/tireg/internal/time-registry/domain/model"
	"github.com/gerarc/tireg/internal/time-registry/domain/usecase"
)

func TestUpdate_Succeeds_WhenTimeEntryValid(t *testing.T) {
	commandRepository := &stubTimeEntryCommandRepository{
		updateByIDAndOwnerFunc: func(id string, ownerID string, timeEntry model.TimeEntry) (model.TimeEntry, error) {
			timeEntry.ID = id
			timeEntry.OwnerID = ownerID
			return timeEntry, nil
		},
	}

	updateTimeEntryUseCase := usecase.NewUpdateTimeEntryUseCase(commandRepository)

	timeEntry, err := updateTimeEntryUseCase.Update("owner-id", "entry-id", validTimeEntry())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if timeEntry.ID != "entry-id" || timeEntry.OwnerID != "owner-id" {
		t.Fatalf("unexpected time entry: %+v", timeEntry)
	}
}

func TestUpdate_ReturnsValidationError_WhenRequiredFieldsMissing(t *testing.T) {
	updateTimeEntryUseCase := usecase.NewUpdateTimeEntryUseCase(&stubTimeEntryCommandRepository{})

	_, err := updateTimeEntryUseCase.Update("owner-id", "entry-id", model.TimeEntry{})

	var domainError *sharedException.DomainError
	if !errors.As(err, &domainError) || domainError.Code != 400 {
		t.Fatalf("expected a 400 DomainError, got %v", err)
	}
}

func TestUpdate_ReturnsNotFoundError_WhenNotOwnedByOwner(t *testing.T) {
	commandRepository := &stubTimeEntryCommandRepository{
		updateByIDAndOwnerFunc: func(id string, ownerID string, timeEntry model.TimeEntry) (model.TimeEntry, error) {
			return model.TimeEntry{}, exception.NewTimeEntryNotFoundError()
		},
	}

	updateTimeEntryUseCase := usecase.NewUpdateTimeEntryUseCase(commandRepository)

	_, err := updateTimeEntryUseCase.Update("owner-id", "someone-elses-id", validTimeEntry())

	var domainError *sharedException.DomainError
	if !errors.As(err, &domainError) || domainError.Code != 404 {
		t.Fatalf("expected a 404 DomainError, got %v", err)
	}
}
