package api

import "github.com/gerarc/tireg/internal/time-registry/domain/model"

// FindTimeEntryByIDUseCase exposes the operation to fetch a single time entry owned by a user.
type FindTimeEntryByIDUseCase interface {
	// FindByID returns the time entry matching the given id, owned by the given user.
	FindByID(ownerID string, id string) (model.TimeEntry, error)
}
