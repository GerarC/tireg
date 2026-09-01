package usecase_test

import (
	"errors"
	"testing"

	sharedException "github.com/gerarc/tireg/internal/shared/domain/exception"

	"github.com/gerarc/tireg/internal/time-registry/domain/exception"
	"github.com/gerarc/tireg/internal/time-registry/domain/model"
	"github.com/gerarc/tireg/internal/time-registry/domain/usecase"
)

func TestFindByID_ReturnsEntry_WhenOwnedByOwner(t *testing.T) {
	queryRepository := &stubTimeEntryQueryRepository{
		selectByIDAndOwnerFunc: func(id string, ownerID string) (model.TimeEntry, error) {
			return model.TimeEntry{ID: id, OwnerID: ownerID}, nil
		},
	}

	findTimeEntryByIDUseCase := usecase.NewFindTimeEntryByIDUseCase(queryRepository)

	timeEntry, err := findTimeEntryByIDUseCase.FindByID("owner-id", "entry-id")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if timeEntry.ID != "entry-id" || timeEntry.OwnerID != "owner-id" {
		t.Fatalf("unexpected time entry: %+v", timeEntry)
	}
}

func TestFindByID_ReturnsNotFoundError_WhenNotOwnedByOwner(t *testing.T) {
	queryRepository := &stubTimeEntryQueryRepository{
		selectByIDAndOwnerFunc: func(id string, ownerID string) (model.TimeEntry, error) {
			return model.TimeEntry{}, exception.NewTimeEntryNotFoundError()
		},
	}

	findTimeEntryByIDUseCase := usecase.NewFindTimeEntryByIDUseCase(queryRepository)

	_, err := findTimeEntryByIDUseCase.FindByID("owner-id", "someone-elses-id")

	var domainError *sharedException.DomainError
	if !errors.As(err, &domainError) || domainError.Code != 404 {
		t.Fatalf("expected a 404 DomainError, got %v", err)
	}
}
