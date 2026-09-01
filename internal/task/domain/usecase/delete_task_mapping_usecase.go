package usecase

import (
	"github.com/gerarc/tireg/internal/task/domain/api"
	"github.com/gerarc/tireg/internal/task/domain/spi"
)

type DeleteTaskMappingUseCaseImplemented struct {
	taskMappingCommandRepository spi.TaskMappingCommandRepository
}

func NewDeleteTaskMappingUseCase(taskMappingCommandRepository spi.TaskMappingCommandRepository) api.DeleteTaskMappingUseCase {
	return &DeleteTaskMappingUseCaseImplemented{taskMappingCommandRepository: taskMappingCommandRepository}
}

func (deleteTaskMappingUseCase *DeleteTaskMappingUseCaseImplemented) Delete(ownerID string, id string) error {
	return deleteTaskMappingUseCase.taskMappingCommandRepository.DeleteByIDAndOwner(id, ownerID)
}
