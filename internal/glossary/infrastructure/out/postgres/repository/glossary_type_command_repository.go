package repository

import (
	"gorm.io/gorm"

	"github.com/gerarc/tireg/internal/glossary/infrastructure/out/postgres/entity"
)

type GlossaryTypeCommandRepository struct {
	db *gorm.DB
}

func NewGlossaryTypeCommandRepository(db *gorm.DB) *GlossaryTypeCommandRepository {
	return &GlossaryTypeCommandRepository{db: db}
}

func (glossaryTypeCommandRepository *GlossaryTypeCommandRepository) Insert(glossaryTypeEntity entity.GlossaryTypeEntity) (entity.GlossaryTypeEntity, error) {
	err := glossaryTypeCommandRepository.db.Create(&glossaryTypeEntity).Error
	return glossaryTypeEntity, err
}

func (glossaryTypeCommandRepository *GlossaryTypeCommandRepository) UpdateByIDAndOwner(id string, ownerID string, glossaryTypeEntity entity.GlossaryTypeEntity) (int64, error) {
	result := glossaryTypeCommandRepository.db.
		Model(&entity.GlossaryTypeEntity{}).
		Where("id = ? AND owner_id = ?", id, ownerID).
		Updates(map[string]any{
			"type_key":    glossaryTypeEntity.TypeKey,
			"label":       glossaryTypeEntity.Label,
			"description": glossaryTypeEntity.Description,
		})

	return result.RowsAffected, result.Error
}

func (glossaryTypeCommandRepository *GlossaryTypeCommandRepository) DeleteByIDAndOwner(id string, ownerID string) (int64, error) {
	result := glossaryTypeCommandRepository.db.
		Where("id = ? AND owner_id = ?", id, ownerID).
		Delete(&entity.GlossaryTypeEntity{})

	return result.RowsAffected, result.Error
}

func (glossaryTypeCommandRepository *GlossaryTypeCommandRepository) FindByIDAndOwner(id string, ownerID string) (entity.GlossaryTypeEntity, error) {
	var found entity.GlossaryTypeEntity
	err := glossaryTypeCommandRepository.db.Where("id = ? AND owner_id = ?", id, ownerID).First(&found).Error
	return found, err
}
