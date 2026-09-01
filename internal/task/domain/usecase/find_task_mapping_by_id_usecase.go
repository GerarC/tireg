package usecase

import (
	"github.com/gerarc/tireg/internal/task/domain/api"
	"github.com/gerarc/tireg/internal/task/domain/model"
	"github.com/gerarc/tireg/internal/task/domain/spi"
)

type FindTaskMappingByIDUseCaseImplemented struct {
	taskMappingQueryRepository spi.TaskMappingQueryRepository
}

func NewFindTaskMappingByIDUseCase(taskMappingQueryRepository spi.TaskMappingQueryRepository) api.FindTaskMappingByIDUseCase {
	return &FindTaskMappingByIDUseCaseImplemented{taskMappingQueryRepository: taskMappingQueryRepository}
}

func (findTaskMappingByIDUseCase *FindTaskMappingByIDUseCaseImplemented) FindByID(ownerID string, id string) (model.TaskMapping, error) {
	return findTaskMappingByIDUseCase.taskMappingQueryRepository.SelectByIDAndOwner(id, ownerID)
}
