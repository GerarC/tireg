package usecase

import (
	"strings"

	taskApi "github.com/gerarc/tireg/internal/task/domain/api"
	taskModel "github.com/gerarc/tireg/internal/task/domain/model"

	"github.com/gerarc/tireg/internal/time-registry/domain/api"
	"github.com/gerarc/tireg/internal/time-registry/domain/exception"
	"github.com/gerarc/tireg/internal/time-registry/domain/model"
	"github.com/gerarc/tireg/internal/time-registry/domain/spi"
	"github.com/gerarc/tireg/internal/time-registry/domain/util/constant"
)

type CreateTimeEntryUseCaseImplemented struct {
	timeEntryCommandRepository spi.TimeEntryCommandRepository
	listTaskMappingsUseCase    taskApi.ListTaskMappingsUseCase
}

func NewCreateTimeEntryUseCase(timeEntryCommandRepository spi.TimeEntryCommandRepository, listTaskMappingsUseCase taskApi.ListTaskMappingsUseCase) api.CreateTimeEntryUseCase {
	return &CreateTimeEntryUseCaseImplemented{
		timeEntryCommandRepository: timeEntryCommandRepository,
		listTaskMappingsUseCase:    listTaskMappingsUseCase,
	}
}

func (createTimeEntryUseCase *CreateTimeEntryUseCaseImplemented) Create(ownerID string, timeEntry model.TimeEntry) (model.TimeEntry, error) {
	if timeEntry.ProjectLabel == "" || timeEntry.TypeKey == "" || timeEntry.IssueKey == "" {
		taskMappings, err := createTimeEntryUseCase.listTaskMappingsUseCase.List(ownerID)
		if err != nil {
			return model.TimeEntry{}, err
		}

		if matched, ok := matchTaskMapping(taskMappings, timeEntry.Description); ok {
			if timeEntry.ProjectLabel == "" {
				timeEntry.ProjectLabel = matched.ProjectLabel
			}
			if timeEntry.TypeKey == "" {
				timeEntry.TypeKey = matched.TypeKey
			}
			if timeEntry.IssueKey == "" {
				timeEntry.IssueKey = matched.IssueKey
			}
		}
	}

	if details := validateTimeEntry(timeEntry); len(details) > 0 {
		return model.TimeEntry{}, exception.NewValidationFailedError(details...)
	}

	timeEntry.OwnerID = ownerID

	return createTimeEntryUseCase.timeEntryCommandRepository.Insert(timeEntry)
}

func matchTaskMapping(taskMappings []taskModel.TaskMapping, description string) (taskModel.TaskMapping, bool) {
	loweredDescription := strings.ToLower(description)

	for _, taskMapping := range taskMappings {
		for _, keyword := range taskMapping.MatchKeywords {
			if keyword != "" && strings.Contains(loweredDescription, strings.ToLower(keyword)) {
				return taskMapping, true
			}
		}
	}

	return taskModel.TaskMapping{}, false
}

func validateTimeEntry(timeEntry model.TimeEntry) []string {
	var details []string

	if strings.TrimSpace(timeEntry.Date) == "" {
		details = append(details, constant.DetailDateRequired)
	}

	if strings.TrimSpace(timeEntry.ProjectLabel) == "" {
		details = append(details, constant.DetailProjectLabelRequired)
	}

	if strings.TrimSpace(timeEntry.Start) == "" {
		details = append(details, constant.DetailStartRequired)
	}

	if strings.TrimSpace(timeEntry.End) == "" {
		details = append(details, constant.DetailEndRequired)
	}

	if timeEntry.Hours <= 0 {
		details = append(details, constant.DetailHoursMustBePositive)
	}

	return details
}
