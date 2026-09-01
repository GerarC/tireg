package repository

import (
	"gorm.io/gorm"

	"github.com/gerarc/tireg/internal/glossary/infrastructure/out/postgres/entity"
)

type GlossaryProjectQueryRepository struct {
	db *gorm.DB
}

func NewGlossaryProjectQueryRepository(db *gorm.DB) *GlossaryProjectQueryRepository {
	return &GlossaryProjectQueryRepository{db: db}
}

func (glossaryProjectQueryRepository *GlossaryProjectQueryRepository) SelectAllByOwner(ownerID string) ([]entity.GlossaryProjectEntity, error) {
	var found []entity.GlossaryProjectEntity
	err := glossaryProjectQueryRepository.db.Where("owner_id = ?", ownerID).Find(&found).Error
	return found, err
}
