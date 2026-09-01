package usecase

import (
	"github.com/gerarc/tireg/internal/time-registry/domain/api"
	"github.com/gerarc/tireg/internal/time-registry/domain/model"
	"github.com/gerarc/tireg/internal/time-registry/domain/spi"
)

type FindTimeEntryByIDUseCaseImplemented struct {
	timeEntryQueryRepository spi.TimeEntryQueryRepository
}

func NewFindTimeEntryByIDUseCase(timeEntryQueryRepository spi.TimeEntryQueryRepository) api.FindTimeEntryByIDUseCase {
	return &FindTimeEntryByIDUseCaseImplemented{timeEntryQueryRepository: timeEntryQueryRepository}
}

func (findTimeEntryByIDUseCase *FindTimeEntryByIDUseCaseImplemented) FindByID(ownerID string, id string) (model.TimeEntry, error) {
	return findTimeEntryByIDUseCase.timeEntryQueryRepository.SelectByIDAndOwner(id, ownerID)
}
