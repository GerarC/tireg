package api

import "github.com/gerarc/tireg/internal/task/domain/model"

// CreateTaskMappingUseCase exposes the operation to create a new task mapping owned by a user.
type CreateTaskMappingUseCase interface {
	// Create validates and persists a new task mapping for the given owner.
	Create(ownerID string, taskMapping model.TaskMapping) (model.TaskMapping, error)
}
