package api

import "github.com/gerarc/tireg/internal/time-registry/domain/model"

// UpdateTimeEntryUseCase exposes the operation to update a time entry owned by a user.
type UpdateTimeEntryUseCase interface {
	// Update validates and updates the time entry matching the given id, owned by the given user.
	Update(ownerID string, id string, timeEntry model.TimeEntry) (model.TimeEntry, error)
}
