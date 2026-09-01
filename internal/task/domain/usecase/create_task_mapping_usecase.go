package usecase

import (
	"strings"

	"github.com/gerarc/tireg/internal/task/domain/api"
	"github.com/gerarc/tireg/internal/task/domain/exception"
	"github.com/gerarc/tireg/internal/task/domain/model"
	"github.com/gerarc/tireg/internal/task/domain/spi"
	"github.com/gerarc/tireg/internal/task/domain/util/constant"
)

type CreateTaskMappingUseCaseImplemented struct {
	taskMappingCommandRepository spi.TaskMappingCommandRepository
}

func NewCreateTaskMappingUseCase(taskMappingCommandRepository spi.TaskMappingCommandRepository) api.CreateTaskMappingUseCase {
	return &CreateTaskMappingUseCaseImplemented{taskMappingCommandRepository: taskMappingCommandRepository}
}

func (createTaskMappingUseCase *CreateTaskMappingUseCaseImplemented) Create(ownerID string, taskMapping model.TaskMapping) (model.TaskMapping, error) {
	if details := validateTaskMapping(taskMapping); len(details) > 0 {
		return model.TaskMapping{}, exception.NewValidationFailedError(details...)
	}

	taskMapping.OwnerID = ownerID

	return createTaskMappingUseCase.taskMappingCommandRepository.Insert(taskMapping)
}

func validateTaskMapping(taskMapping model.TaskMapping) []string {
	var details []string

	if strings.TrimSpace(taskMapping.ProjectLabel) == "" {
		details = append(details, constant.DetailProjectLabelRequired)
	}

	if strings.TrimSpace(taskMapping.Pattern) == "" {
		details = append(details, constant.DetailPatternRequired)
	}

	return details
}
