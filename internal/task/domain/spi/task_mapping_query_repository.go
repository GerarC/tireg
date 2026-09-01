package spi

import "github.com/gerarc/tireg/internal/task/domain/model"

// TaskMappingQueryRepository defines the read persistence operations required by the task mapping use cases.
type TaskMappingQueryRepository interface {
	// SelectAllByOwner returns every task mapping owned by the given user.
	SelectAllByOwner(ownerID string) ([]model.TaskMapping, error)
	// SelectByIDAndOwner returns the task mapping matching the given id and owner.
	SelectByIDAndOwner(id string, ownerID string) (model.TaskMapping, error)
}
