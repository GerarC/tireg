package api

import "github.com/gerarc/tireg/internal/time-registry/domain/model"

// ListTimeEntriesUseCase exposes the operation to list every time entry owned by a user.
type ListTimeEntriesUseCase interface {
	// List returns every time entry owned by the given user.
	List(ownerID string) ([]model.TimeEntry, error)
}
