package usecase

import (
	"github.com/gerarc/tireg/internal/task/domain/api"
	"github.com/gerarc/tireg/internal/task/domain/exception"
	"github.com/gerarc/tireg/internal/task/domain/model"
	"github.com/gerarc/tireg/internal/task/domain/spi"
)

type UpdateTaskMappingUseCaseImplemented struct {
	taskMappingCommandRepository spi.TaskMappingCommandRepository
}

func NewUpdateTaskMappingUseCase(taskMappingCommandRepository spi.TaskMappingCommandRepository) api.UpdateTaskMappingUseCase {
	return &UpdateTaskMappingUseCaseImplemented{taskMappingCommandRepository: taskMappingCommandRepository}
}

func (updateTaskMappingUseCase *UpdateTaskMappingUseCaseImplemented) Update(ownerID string, id string, taskMapping model.TaskMapping) (model.TaskMapping, error) {
	if details := validateTaskMapping(taskMapping); len(details) > 0 {
		return model.TaskMapping{}, exception.NewValidationFailedError(details...)
	}

	return updateTaskMappingUseCase.taskMappingCommandRepository.UpdateByIDAndOwner(id, ownerID, taskMapping)
}
