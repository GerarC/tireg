package usecase

import (
	"github.com/gerarc/tireg/internal/time-registry/domain/api"
	"github.com/gerarc/tireg/internal/time-registry/domain/spi"
)

type DeleteTimeEntryUseCaseImplemented struct {
	timeEntryCommandRepository spi.TimeEntryCommandRepository
}

func NewDeleteTimeEntryUseCase(timeEntryCommandRepository spi.TimeEntryCommandRepository) api.DeleteTimeEntryUseCase {
	return &DeleteTimeEntryUseCaseImplemented{timeEntryCommandRepository: timeEntryCommandRepository}
}

func (deleteTimeEntryUseCase *DeleteTimeEntryUseCaseImplemented) Delete(ownerID string, id string) error {
	return deleteTimeEntryUseCase.timeEntryCommandRepository.DeleteByIDAndOwner(id, ownerID)
}
