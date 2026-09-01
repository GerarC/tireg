package repository

import (
	"gorm.io/gorm"

	"github.com/gerarc/tireg/internal/time-registry/infrastructure/out/postgres/entity"
)

type TimeEntryCommandRepository struct {
	db *gorm.DB
}

func NewTimeEntryCommandRepository(db *gorm.DB) *TimeEntryCommandRepository {
	return &TimeEntryCommandRepository{db: db}
}

func (timeEntryCommandRepository *TimeEntryCommandRepository) Insert(timeEntryEntity entity.TimeEntryEntity) (entity.TimeEntryEntity, error) {
	err := timeEntryCommandRepository.db.Create(&timeEntryEntity).Error
	return timeEntryEntity, err
}

func (timeEntryCommandRepository *TimeEntryCommandRepository) UpdateByIDAndOwner(id string, ownerID string, timeEntryEntity entity.TimeEntryEntity) (int64, error) {
	result := timeEntryCommandRepository.db.
		Model(&entity.TimeEntryEntity{}).
		Where("id = ? AND owner_id = ?", id, ownerID).
		Updates(map[string]any{
			"date":            timeEntryEntity.Date,
			"project_label":   timeEntryEntity.ProjectLabel,
			"type_key":        timeEntryEntity.TypeKey,
			"issue_key":       timeEntryEntity.IssueKey,
			"start":           timeEntryEntity.Start,
			"end":             timeEntryEntity.End,
			"hours":           timeEntryEntity.Hours,
			"description":     timeEntryEntity.Description,
			"jira_worklog_id": timeEntryEntity.JiraWorklogID,
		})

	return result.RowsAffected, result.Error
}

func (timeEntryCommandRepository *TimeEntryCommandRepository) DeleteByIDAndOwner(id string, ownerID string) (int64, error) {
	result := timeEntryCommandRepository.db.
		Where("id = ? AND owner_id = ?", id, ownerID).
		Delete(&entity.TimeEntryEntity{})

	return result.RowsAffected, result.Error
}

func (timeEntryCommandRepository *TimeEntryCommandRepository) FindByIDAndOwner(id string, ownerID string) (entity.TimeEntryEntity, error) {
	var found entity.TimeEntryEntity
	err := timeEntryCommandRepository.db.Where("id = ? AND owner_id = ?", id, ownerID).First(&found).Error
	return found, err
}
