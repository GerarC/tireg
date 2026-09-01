package adapter

import (
	"github.com/gerarc/tireg/internal/glossary/domain/exception"
	"github.com/gerarc/tireg/internal/glossary/domain/model"
	"github.com/gerarc/tireg/internal/glossary/domain/spi"
	"github.com/gerarc/tireg/internal/glossary/infrastructure/out/postgres/mapper"
	"github.com/gerarc/tireg/internal/glossary/infrastructure/out/postgres/repository"
)

type GlossaryTypeCommandAdapter struct {
	glossaryTypeCommandRepository *repository.GlossaryTypeCommandRepository
}

func NewGlossaryTypeCommandAdapter(glossaryTypeCommandRepository *repository.GlossaryTypeCommandRepository) spi.GlossaryTypeCommandRepository {
	return &GlossaryTypeCommandAdapter{glossaryTypeCommandRepository: glossaryTypeCommandRepository}
}

func (glossaryTypeCommandAdapter *GlossaryTypeCommandAdapter) Insert(glossaryType model.GlossaryType) (model.GlossaryType, error) {
	inserted, err := glossaryTypeCommandAdapter.glossaryTypeCommandRepository.Insert(mapper.ToGlossaryTypeEntity(glossaryType))
	if err != nil {
		return model.GlossaryType{}, err
	}

	return mapper.ToGlossaryTypeModel(inserted), nil
}

func (glossaryTypeCommandAdapter *GlossaryTypeCommandAdapter) UpdateByIDAndOwner(id string, ownerID string, glossaryType model.GlossaryType) (model.GlossaryType, error) {
	rowsAffected, err := glossaryTypeCommandAdapter.glossaryTypeCommandRepository.UpdateByIDAndOwner(id, ownerID, mapper.ToGlossaryTypeEntity(glossaryType))
	if err != nil {
		return model.GlossaryType{}, err
	}

	if rowsAffected == 0 {
		return model.GlossaryType{}, exception.NewGlossaryTypeNotFoundError()
	}

	updated, err := glossaryTypeCommandAdapter.glossaryTypeCommandRepository.FindByIDAndOwner(id, ownerID)
	if err != nil {
		return model.GlossaryType{}, err
	}

	return mapper.ToGlossaryTypeModel(updated), nil
}

func (glossaryTypeCommandAdapter *GlossaryTypeCommandAdapter) DeleteByIDAndOwner(id string, ownerID string) error {
	rowsAffected, err := glossaryTypeCommandAdapter.glossaryTypeCommandRepository.DeleteByIDAndOwner(id, ownerID)
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return exception.NewGlossaryTypeNotFoundError()
	}

	return nil
}
