package usecase

import (
	"github.com/gerarc/tireg/internal/time-registry/domain/api"
	"github.com/gerarc/tireg/internal/time-registry/domain/model"
	"github.com/gerarc/tireg/internal/time-registry/domain/spi"
)

type ListTimeEntriesUseCaseImplemented struct {
	timeEntryQueryRepository spi.TimeEntryQueryRepository
}

func NewListTimeEntriesUseCase(timeEntryQueryRepository spi.TimeEntryQueryRepository) api.ListTimeEntriesUseCase {
	return &ListTimeEntriesUseCaseImplemented{timeEntryQueryRepository: timeEntryQueryRepository}
}

func (listTimeEntriesUseCase *ListTimeEntriesUseCaseImplemented) List(ownerID string) ([]model.TimeEntry, error) {
	return listTimeEntriesUseCase.timeEntryQueryRepository.SelectAllByOwner(ownerID)
}
