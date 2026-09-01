package repository

import (
	"gorm.io/gorm"

	"github.com/gerarc/tireg/internal/task/infrastructure/out/postgres/entity"
)

type TaskMappingCommandRepository struct {
	db *gorm.DB
}

func NewTaskMappingCommandRepository(db *gorm.DB) *TaskMappingCommandRepository {
	return &TaskMappingCommandRepository{db: db}
}

func (taskMappingCommandRepository *TaskMappingCommandRepository) Insert(taskMappingEntity entity.TaskMappingEntity) (entity.TaskMappingEntity, error) {
	err := taskMappingCommandRepository.db.Create(&taskMappingEntity).Error
	return taskMappingEntity, err
}

func (taskMappingCommandRepository *TaskMappingCommandRepository) UpdateByIDAndOwner(id string, ownerID string, taskMappingEntity entity.TaskMappingEntity) (int64, error) {
	result := taskMappingCommandRepository.db.
		Model(&entity.TaskMappingEntity{}).
		Where("id = ? AND owner_id = ?", id, ownerID).
		Updates(map[string]any{
			"project_label":          taskMappingEntity.ProjectLabel,
			"pattern":                taskMappingEntity.Pattern,
			"match_keywords":         taskMappingEntity.MatchKeywords,
			"match_organizer_domain": taskMappingEntity.MatchOrganizerDomain,
			"issue_key":              taskMappingEntity.IssueKey,
			"type_key":               taskMappingEntity.TypeKey,
			"notes":                  taskMappingEntity.Notes,
		})

	return result.RowsAffected, result.Error
}

func (taskMappingCommandRepository *TaskMappingCommandRepository) DeleteByIDAndOwner(id string, ownerID string) (int64, error) {
	result := taskMappingCommandRepository.db.
		Where("id = ? AND owner_id = ?", id, ownerID).
		Delete(&entity.TaskMappingEntity{})

	return result.RowsAffected, result.Error
}

func (taskMappingCommandRepository *TaskMappingCommandRepository) FindByIDAndOwner(id string, ownerID string) (entity.TaskMappingEntity, error) {
	var found entity.TaskMappingEntity
	err := taskMappingCommandRepository.db.Where("id = ? AND owner_id = ?", id, ownerID).First(&found).Error
	return found, err
}
