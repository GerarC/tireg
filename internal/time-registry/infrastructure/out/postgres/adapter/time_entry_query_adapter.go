package adapter

import (
	"errors"

	"gorm.io/gorm"

	"github.com/gerarc/tireg/internal/time-registry/domain/exception"
	"github.com/gerarc/tireg/internal/time-registry/domain/model"
	"github.com/gerarc/tireg/internal/time-registry/domain/spi"
	"github.com/gerarc/tireg/internal/time-registry/infrastructure/out/postgres/mapper"
	"github.com/gerarc/tireg/internal/time-registry/infrastructure/out/postgres/repository"
)

type TimeEntryQueryAdapter struct {
	timeEntryQueryRepository *repository.TimeEntryQueryRepository
}

func NewTimeEntryQueryAdapter(timeEntryQueryRepository *repository.TimeEntryQueryRepository) spi.TimeEntryQueryRepository {
	return &TimeEntryQueryAdapter{timeEntryQueryRepository: timeEntryQueryRepository}
}

func (timeEntryQueryAdapter *TimeEntryQueryAdapter) SelectAllByOwner(ownerID string) ([]model.TimeEntry, error) {
	found, err := timeEntryQueryAdapter.timeEntryQueryRepository.SelectAllByOwner(ownerID)
	if err != nil {
		return nil, err
	}

	return mapper.ToTimeEntryModelList(found), nil
}

func (timeEntryQueryAdapter *TimeEntryQueryAdapter) SelectByIDAndOwner(id string, ownerID string) (model.TimeEntry, error) {
	found, err := timeEntryQueryAdapter.timeEntryQueryRepository.SelectByIDAndOwner(id, ownerID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.TimeEntry{}, exception.NewTimeEntryNotFoundError()
		}

		return model.TimeEntry{}, err
	}

	return mapper.ToTimeEntryModel(found), nil
}
