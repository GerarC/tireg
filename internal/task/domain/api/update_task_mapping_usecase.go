package api

import "github.com/gerarc/tireg/internal/task/domain/model"

// UpdateTaskMappingUseCase exposes the operation to update a task mapping owned by a user.
type UpdateTaskMappingUseCase interface {
	// Update validates and updates the task mapping matching the given id, owned by the given user.
	Update(ownerID string, id string, taskMapping model.TaskMapping) (model.TaskMapping, error)
}
