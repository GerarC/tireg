package repository

import (
	"gorm.io/gorm"

	"github.com/gerarc/tireg/internal/time-registry/infrastructure/out/postgres/entity"
)

type TimeEntryQueryRepository struct {
	db *gorm.DB
}

func NewTimeEntryQueryRepository(db *gorm.DB) *TimeEntryQueryRepository {
	return &TimeEntryQueryRepository{db: db}
}

func (timeEntryQueryRepository *TimeEntryQueryRepository) SelectAllByOwner(ownerID string) ([]entity.TimeEntryEntity, error) {
	var found []entity.TimeEntryEntity
	err := timeEntryQueryRepository.db.Where("owner_id = ?", ownerID).Find(&found).Error
	return found, err
}

func (timeEntryQueryRepository *TimeEntryQueryRepository) SelectByIDAndOwner(id string, ownerID string) (entity.TimeEntryEntity, error) {
	var found entity.TimeEntryEntity
	err := timeEntryQueryRepository.db.Where("id = ? AND owner_id = ?", id, ownerID).First(&found).Error
	return found, err
}
