package api

import "github.com/gerarc/tireg/internal/time-registry/domain/model"

// CreateTimeEntryUseCase exposes the operation to register a new time entry owned by a user.
type CreateTimeEntryUseCase interface {
	// Create auto-fills classification fields from the owner's task mappings when left blank, validates, and persists a new time entry for the given owner.
	Create(ownerID string, timeEntry model.TimeEntry) (model.TimeEntry, error)
}
