package spi

import "github.com/gerarc/tireg/internal/time-registry/domain/model"

// TimeEntryQueryRepository defines the read persistence operations required by the time entry use cases.
type TimeEntryQueryRepository interface {
	// SelectAllByOwner returns every time entry owned by the given user.
	SelectAllByOwner(ownerID string) ([]model.TimeEntry, error)
	// SelectByIDAndOwner returns the time entry matching the given id and owner.
	SelectByIDAndOwner(id string, ownerID string) (model.TimeEntry, error)
}
