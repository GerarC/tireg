package spi

import "github.com/gerarc/tireg/internal/task/domain/model"

// TaskMappingCommandRepository defines the write persistence operations required by the task mapping use cases.
type TaskMappingCommandRepository interface {
	// Insert persists a new task mapping and returns it with its generated ID and audit timestamps.
	Insert(taskMapping model.TaskMapping) (model.TaskMapping, error)
	// UpdateByIDAndOwner updates the task mapping matching the given id and owner, returning it updated.
	UpdateByIDAndOwner(id string, ownerID string, taskMapping model.TaskMapping) (model.TaskMapping, error)
	// DeleteByIDAndOwner deletes the task mapping matching the given id and owner.
	DeleteByIDAndOwner(id string, ownerID string) error
}
