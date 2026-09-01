package adapter

import (
	"errors"

	"gorm.io/gorm"

	"github.com/gerarc/tireg/internal/task/domain/exception"
	"github.com/gerarc/tireg/internal/task/domain/model"
	"github.com/gerarc/tireg/internal/task/domain/spi"
	"github.com/gerarc/tireg/internal/task/infrastructure/out/postgres/mapper"
	"github.com/gerarc/tireg/internal/task/infrastructure/out/postgres/repository"
)

type TaskMappingQueryAdapter struct {
	taskMappingQueryRepository *repository.TaskMappingQueryRepository
}

func NewTaskMappingQueryAdapter(taskMappingQueryRepository *repository.TaskMappingQueryRepository) spi.TaskMappingQueryRepository {
	return &TaskMappingQueryAdapter{taskMappingQueryRepository: taskMappingQueryRepository}
}

func (taskMappingQueryAdapter *TaskMappingQueryAdapter) SelectAllByOwner(ownerID string) ([]model.TaskMapping, error) {
	found, err := taskMappingQueryAdapter.taskMappingQueryRepository.SelectAllByOwner(ownerID)
	if err != nil {
		return nil, err
	}

	return mapper.ToTaskMappingModelList(found), nil
}

func (taskMappingQueryAdapter *TaskMappingQueryAdapter) SelectByIDAndOwner(id string, ownerID string) (model.TaskMapping, error) {
	found, err := taskMappingQueryAdapter.taskMappingQueryRepository.SelectByIDAndOwner(id, ownerID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.TaskMapping{}, exception.NewTaskMappingNotFoundError()
		}

		return model.TaskMapping{}, err
	}

	return mapper.ToTaskMappingModel(found), nil
}
