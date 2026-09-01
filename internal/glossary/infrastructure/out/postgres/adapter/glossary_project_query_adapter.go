package adapter

import (
	"github.com/gerarc/tireg/internal/glossary/domain/model"
	"github.com/gerarc/tireg/internal/glossary/domain/spi"
	"github.com/gerarc/tireg/internal/glossary/infrastructure/out/postgres/mapper"
	"github.com/gerarc/tireg/internal/glossary/infrastructure/out/postgres/repository"
)

type GlossaryProjectQueryAdapter struct {
	glossaryProjectQueryRepository *repository.GlossaryProjectQueryRepository
}

func NewGlossaryProjectQueryAdapter(glossaryProjectQueryRepository *repository.GlossaryProjectQueryRepository) spi.GlossaryProjectQueryRepository {
	return &GlossaryProjectQueryAdapter{glossaryProjectQueryRepository: glossaryProjectQueryRepository}
}

func (glossaryProjectQueryAdapter *GlossaryProjectQueryAdapter) SelectAllByOwner(ownerID string) ([]model.GlossaryProject, error) {
	found, err := glossaryProjectQueryAdapter.glossaryProjectQueryRepository.SelectAllByOwner(ownerID)
	if err != nil {
		return nil, err
	}

	return mapper.ToGlossaryProjectModelList(found), nil
}
