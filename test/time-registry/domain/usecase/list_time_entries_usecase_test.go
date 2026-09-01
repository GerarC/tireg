package usecase_test

import (
	"testing"

	"github.com/gerarc/tireg/internal/time-registry/domain/model"
	"github.com/gerarc/tireg/internal/time-registry/domain/usecase"
)

func TestList_ReturnsEntriesOwnedByOwner(t *testing.T) {
	queryRepository := &stubTimeEntryQueryRepository{
		selectAllByOwnerFunc: func(ownerID string) ([]model.TimeEntry, error) {
			return []model.TimeEntry{{ID: "entry-id", OwnerID: ownerID}}, nil
		},
	}

	listTimeEntriesUseCase := usecase.NewListTimeEntriesUseCase(queryRepository)

	timeEntries, err := listTimeEntriesUseCase.List("owner-id")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(timeEntries) != 1 || timeEntries[0].ID != "entry-id" {
		t.Fatalf("unexpected time entries: %+v", timeEntries)
	}
}

func TestList_ReturnsEmptySlice_WhenOwnerHasNoEntries(t *testing.T) {
	queryRepository := &stubTimeEntryQueryRepository{
		selectAllByOwnerFunc: func(ownerID string) ([]model.TimeEntry, error) {
			return nil, nil
		},
	}

	listTimeEntriesUseCase := usecase.NewListTimeEntriesUseCase(queryRepository)

	timeEntries, err := listTimeEntriesUseCase.List("owner-id")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(timeEntries) != 0 {
		t.Fatalf("expected no entries, got %d", len(timeEntries))
	}
}
