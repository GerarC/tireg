package repository

import (
	"gorm.io/gorm"

	"github.com/gerarc/tireg/internal/task/infrastructure/out/postgres/entity"
)

type TaskMappingQueryRepository struct {
	db *gorm.DB
}

func NewTaskMappingQueryRepository(db *gorm.DB) *TaskMappingQueryRepository {
	return &TaskMappingQueryRepository{db: db}
}

func (taskMappingQueryRepository *TaskMappingQueryRepository) SelectAllByOwner(ownerID string) ([]entity.TaskMappingEntity, error) {
	var found []entity.TaskMappingEntity
	err := taskMappingQueryRepository.db.Where("owner_id = ?", ownerID).Find(&found).Error
	return found, err
}

func (taskMappingQueryRepository *TaskMappingQueryRepository) SelectByIDAndOwner(id string, ownerID string) (entity.TaskMappingEntity, error) {
	var found entity.TaskMappingEntity
	err := taskMappingQueryRepository.db.Where("id = ? AND owner_id = ?", id, ownerID).First(&found).Error
	return found, err
}
