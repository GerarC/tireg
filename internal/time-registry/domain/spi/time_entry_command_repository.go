package spi

import "github.com/gerarc/tireg/internal/time-registry/domain/model"

// TimeEntryCommandRepository defines the write persistence operations required by the time entry use cases.
type TimeEntryCommandRepository interface {
	// Insert persists a new time entry and returns it with its generated ID and audit timestamps.
	Insert(timeEntry model.TimeEntry) (model.TimeEntry, error)
	// UpdateByIDAndOwner updates the time entry matching the given id and owner, returning it updated.
	UpdateByIDAndOwner(id string, ownerID string, timeEntry model.TimeEntry) (model.TimeEntry, error)
	// DeleteByIDAndOwner deletes the time entry matching the given id and owner.
	DeleteByIDAndOwner(id string, ownerID string) error
}
