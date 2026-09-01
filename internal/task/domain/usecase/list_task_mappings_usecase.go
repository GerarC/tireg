package usecase

import (
	"github.com/gerarc/tireg/internal/task/domain/api"
	"github.com/gerarc/tireg/internal/task/domain/model"
	"github.com/gerarc/tireg/internal/task/domain/spi"
)

type ListTaskMappingsUseCaseImplemented struct {
	taskMappingQueryRepository spi.TaskMappingQueryRepository
}

func NewListTaskMappingsUseCase(taskMappingQueryRepository spi.TaskMappingQueryRepository) api.ListTaskMappingsUseCase {
	return &ListTaskMappingsUseCaseImplemented{taskMappingQueryRepository: taskMappingQueryRepository}
}

func (listTaskMappingsUseCase *ListTaskMappingsUseCaseImplemented) List(ownerID string) ([]model.TaskMapping, error) {
	return listTaskMappingsUseCase.taskMappingQueryRepository.SelectAllByOwner(ownerID)
}
