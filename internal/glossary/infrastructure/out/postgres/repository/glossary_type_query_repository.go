package repository

import (
	"gorm.io/gorm"

	"github.com/gerarc/tireg/internal/glossary/infrastructure/out/postgres/entity"
)

type GlossaryTypeQueryRepository struct {
	db *gorm.DB
}

func NewGlossaryTypeQueryRepository(db *gorm.DB) *GlossaryTypeQueryRepository {
	return &GlossaryTypeQueryRepository{db: db}
}

func (glossaryTypeQueryRepository *GlossaryTypeQueryRepository) SelectAllByOwner(ownerID string) ([]entity.GlossaryTypeEntity, error) {
	var found []entity.GlossaryTypeEntity
	err := glossaryTypeQueryRepository.db.Where("owner_id = ?", ownerID).Find(&found).Error
	return found, err
}
