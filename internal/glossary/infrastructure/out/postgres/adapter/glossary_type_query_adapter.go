package adapter

import (
	"github.com/gerarc/tireg/internal/glossary/domain/model"
	"github.com/gerarc/tireg/internal/glossary/domain/spi"
	"github.com/gerarc/tireg/internal/glossary/infrastructure/out/postgres/mapper"
	"github.com/gerarc/tireg/internal/glossary/infrastructure/out/postgres/repository"
)

type GlossaryTypeQueryAdapter struct {
	glossaryTypeQueryRepository *repository.GlossaryTypeQueryRepository
}

func NewGlossaryTypeQueryAdapter(glossaryTypeQueryRepository *repository.GlossaryTypeQueryRepository) spi.GlossaryTypeQueryRepository {
	return &GlossaryTypeQueryAdapter{glossaryTypeQueryRepository: glossaryTypeQueryRepository}
}

func (glossaryTypeQueryAdapter *GlossaryTypeQueryAdapter) SelectAllByOwner(ownerID string) ([]model.GlossaryType, error) {
	found, err := glossaryTypeQueryAdapter.glossaryTypeQueryRepository.SelectAllByOwner(ownerID)
	if err != nil {
		return nil, err
	}

	return mapper.ToGlossaryTypeModelList(found), nil
}
