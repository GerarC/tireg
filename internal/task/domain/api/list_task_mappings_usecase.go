package api

import "github.com/gerarc/tireg/internal/task/domain/model"

// ListTaskMappingsUseCase exposes the operation to list every task mapping owned by a user.
type ListTaskMappingsUseCase interface {
	// List returns every task mapping owned by the given user.
	List(ownerID string) ([]model.TaskMapping, error)
}
