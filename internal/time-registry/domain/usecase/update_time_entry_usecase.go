package usecase

import (
	"github.com/gerarc/tireg/internal/time-registry/domain/api"
	"github.com/gerarc/tireg/internal/time-registry/domain/exception"
	"github.com/gerarc/tireg/internal/time-registry/domain/model"
	"github.com/gerarc/tireg/internal/time-registry/domain/spi"
)

type UpdateTimeEntryUseCaseImplemented struct {
	timeEntryCommandRepository spi.TimeEntryCommandRepository
}

func NewUpdateTimeEntryUseCase(timeEntryCommandRepository spi.TimeEntryCommandRepository) api.UpdateTimeEntryUseCase {
	return &UpdateTimeEntryUseCaseImplemented{timeEntryCommandRepository: timeEntryCommandRepository}
}

func (updateTimeEntryUseCase *UpdateTimeEntryUseCaseImplemented) Update(ownerID string, id string, timeEntry model.TimeEntry) (model.TimeEntry, error) {
	if details := validateTimeEntry(timeEntry); len(details) > 0 {
		return model.TimeEntry{}, exception.NewValidationFailedError(details...)
	}

	return updateTimeEntryUseCase.timeEntryCommandRepository.UpdateByIDAndOwner(id, ownerID, timeEntry)
}
