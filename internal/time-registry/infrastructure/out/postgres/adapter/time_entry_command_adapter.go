package adapter

import (
	"github.com/gerarc/tireg/internal/time-registry/domain/exception"
	"github.com/gerarc/tireg/internal/time-registry/domain/model"
	"github.com/gerarc/tireg/internal/time-registry/domain/spi"
	"github.com/gerarc/tireg/internal/time-registry/infrastructure/out/postgres/mapper"
	"github.com/gerarc/tireg/internal/time-registry/infrastructure/out/postgres/repository"
)

type TimeEntryCommandAdapter struct {
	timeEntryCommandRepository *repository.TimeEntryCommandRepository
}

func NewTimeEntryCommandAdapter(timeEntryCommandRepository *repository.TimeEntryCommandRepository) spi.TimeEntryCommandRepository {
	return &TimeEntryCommandAdapter{timeEntryCommandRepository: timeEntryCommandRepository}
}

func (timeEntryCommandAdapter *TimeEntryCommandAdapter) Insert(timeEntry model.TimeEntry) (model.TimeEntry, error) {
	inserted, err := timeEntryCommandAdapter.timeEntryCommandRepository.Insert(mapper.ToTimeEntryEntity(timeEntry))
	if err != nil {
		return model.TimeEntry{}, err
	}

	return mapper.ToTimeEntryModel(inserted), nil
}

func (timeEntryCommandAdapter *TimeEntryCommandAdapter) UpdateByIDAndOwner(id string, ownerID string, timeEntry model.TimeEntry) (model.TimeEntry, error) {
	rowsAffected, err := timeEntryCommandAdapter.timeEntryCommandRepository.UpdateByIDAndOwner(id, ownerID, mapper.ToTimeEntryEntity(timeEntry))
	if err != nil {
		return model.TimeEntry{}, err
	}

	if rowsAffected == 0 {
		return model.TimeEntry{}, exception.NewTimeEntryNotFoundError()
	}

	updated, err := timeEntryCommandAdapter.timeEntryCommandRepository.FindByIDAndOwner(id, ownerID)
	if err != nil {
		return model.TimeEntry{}, err
	}

	return mapper.ToTimeEntryModel(updated), nil
}

func (timeEntryCommandAdapter *TimeEntryCommandAdapter) DeleteByIDAndOwner(id string, ownerID string) error {
	rowsAffected, err := timeEntryCommandAdapter.timeEntryCommandRepository.DeleteByIDAndOwner(id, ownerID)
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return exception.NewTimeEntryNotFoundError()
	}

	return nil
}
