package repository

import (
	"gorm.io/gorm"

	"github.com/gerarc/tireg/internal/glossary/infrastructure/out/postgres/entity"
)

type GlossaryProjectCommandRepository struct {
	db *gorm.DB
}

func NewGlossaryProjectCommandRepository(db *gorm.DB) *GlossaryProjectCommandRepository {
	return &GlossaryProjectCommandRepository{db: db}
}

func (glossaryProjectCommandRepository *GlossaryProjectCommandRepository) Insert(glossaryProjectEntity entity.GlossaryProjectEntity) (entity.GlossaryProjectEntity, error) {
	err := glossaryProjectCommandRepository.db.Create(&glossaryProjectEntity).Error
	return glossaryProjectEntity, err
}

func (glossaryProjectCommandRepository *GlossaryProjectCommandRepository) UpdateByIDAndOwner(id string, ownerID string, glossaryProjectEntity entity.GlossaryProjectEntity) (int64, error) {
	result := glossaryProjectCommandRepository.db.
		Model(&entity.GlossaryProjectEntity{}).
		Where("id = ? AND owner_id = ?", id, ownerID).
		Updates(map[string]any{
			"project_label":    glossaryProjectEntity.ProjectLabel,
			"client":           glossaryProjectEntity.Client,
			"jira_project_key": glossaryProjectEntity.JiraProjectKey,
			"board_url":        glossaryProjectEntity.BoardURL,
			"notes":            glossaryProjectEntity.Notes,
		})

	return result.RowsAffected, result.Error
}

func (glossaryProjectCommandRepository *GlossaryProjectCommandRepository) DeleteByIDAndOwner(id string, ownerID string) (int64, error) {
	result := glossaryProjectCommandRepository.db.
		Where("id = ? AND owner_id = ?", id, ownerID).
		Delete(&entity.GlossaryProjectEntity{})

	return result.RowsAffected, result.Error
}

func (glossaryProjectCommandRepository *GlossaryProjectCommandRepository) FindByIDAndOwner(id string, ownerID string) (entity.GlossaryProjectEntity, error) {
	var found entity.GlossaryProjectEntity
	err := glossaryProjectCommandRepository.db.Where("id = ? AND owner_id = ?", id, ownerID).First(&found).Error
	return found, err
}
