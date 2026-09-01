package api

import "github.com/gerarc/tireg/internal/task/domain/model"

// FindTaskMappingByIDUseCase exposes the operation to fetch a single task mapping owned by a user.
type FindTaskMappingByIDUseCase interface {
	// FindByID returns the task mapping matching the given id, owned by the given user.
	FindByID(ownerID string, id string) (model.TaskMapping, error)
}
