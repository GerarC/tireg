package adapter

import (
	"github.com/gerarc/tireg/internal/task/domain/exception"
	"github.com/gerarc/tireg/internal/task/domain/model"
	"github.com/gerarc/tireg/internal/task/domain/spi"
	"github.com/gerarc/tireg/internal/task/infrastructure/out/postgres/mapper"
	"github.com/gerarc/tireg/internal/task/infrastructure/out/postgres/repository"
)

type TaskMappingCommandAdapter struct {
	taskMappingCommandRepository *repository.TaskMappingCommandRepository
}

func NewTaskMappingCommandAdapter(taskMappingCommandRepository *repository.TaskMappingCommandRepository) spi.TaskMappingCommandRepository {
	return &TaskMappingCommandAdapter{taskMappingCommandRepository: taskMappingCommandRepository}
}

func (taskMappingCommandAdapter *TaskMappingCommandAdapter) Insert(taskMapping model.TaskMapping) (model.TaskMapping, error) {
	inserted, err := taskMappingCommandAdapter.taskMappingCommandRepository.Insert(mapper.ToTaskMappingEntity(taskMapping))
	if err != nil {
		return model.TaskMapping{}, err
	}

	return mapper.ToTaskMappingModel(inserted), nil
}

func (taskMappingCommandAdapter *TaskMappingCommandAdapter) UpdateByIDAndOwner(id string, ownerID string, taskMapping model.TaskMapping) (model.TaskMapping, error) {
	rowsAffected, err := taskMappingCommandAdapter.taskMappingCommandRepository.UpdateByIDAndOwner(id, ownerID, mapper.ToTaskMappingEntity(taskMapping))
	if err != nil {
		return model.TaskMapping{}, err
	}

	if rowsAffected == 0 {
		return model.TaskMapping{}, exception.NewTaskMappingNotFoundError()
	}

	updated, err := taskMappingCommandAdapter.taskMappingCommandRepository.FindByIDAndOwner(id, ownerID)
	if err != nil {
		return model.TaskMapping{}, err
	}

	return mapper.ToTaskMappingModel(updated), nil
}

func (taskMappingCommandAdapter *TaskMappingCommandAdapter) DeleteByIDAndOwner(id string, ownerID string) error {
	rowsAffected, err := taskMappingCommandAdapter.taskMappingCommandRepository.DeleteByIDAndOwner(id, ownerID)
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return exception.NewTaskMappingNotFoundError()
	}

	return nil
}
