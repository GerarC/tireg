package adapter

import (
	"github.com/gerarc/tireg/internal/glossary/domain/exception"
	"github.com/gerarc/tireg/internal/glossary/domain/model"
	"github.com/gerarc/tireg/internal/glossary/domain/spi"
	"github.com/gerarc/tireg/internal/glossary/infrastructure/out/postgres/mapper"
	"github.com/gerarc/tireg/internal/glossary/infrastructure/out/postgres/repository"
)

type GlossaryProjectCommandAdapter struct {
	glossaryProjectCommandRepository *repository.GlossaryProjectCommandRepository
}

func NewGlossaryProjectCommandAdapter(glossaryProjectCommandRepository *repository.GlossaryProjectCommandRepository) spi.GlossaryProjectCommandRepository {
	return &GlossaryProjectCommandAdapter{glossaryProjectCommandRepository: glossaryProjectCommandRepository}
}

func (glossaryProjectCommandAdapter *GlossaryProjectCommandAdapter) Insert(glossaryProject model.GlossaryProject) (model.GlossaryProject, error) {
	inserted, err := glossaryProjectCommandAdapter.glossaryProjectCommandRepository.Insert(mapper.ToGlossaryProjectEntity(glossaryProject))
	if err != nil {
		return model.GlossaryProject{}, err
	}

	return mapper.ToGlossaryProjectModel(inserted), nil
}

func (glossaryProjectCommandAdapter *GlossaryProjectCommandAdapter) UpdateByIDAndOwner(id string, ownerID string, glossaryProject model.GlossaryProject) (model.GlossaryProject, error) {
	rowsAffected, err := glossaryProjectCommandAdapter.glossaryProjectCommandRepository.UpdateByIDAndOwner(id, ownerID, mapper.ToGlossaryProjectEntity(glossaryProject))
	if err != nil {
		return model.GlossaryProject{}, err
	}

	if rowsAffected == 0 {
		return model.GlossaryProject{}, exception.NewGlossaryProjectNotFoundError()
	}

	updated, err := glossaryProjectCommandAdapter.glossaryProjectCommandRepository.FindByIDAndOwner(id, ownerID)
	if err != nil {
		return model.GlossaryProject{}, err
	}

	return mapper.ToGlossaryProjectModel(updated), nil
}

func (glossaryProjectCommandAdapter *GlossaryProjectCommandAdapter) DeleteByIDAndOwner(id string, ownerID string) error {
	rowsAffected, err := glossaryProjectCommandAdapter.glossaryProjectCommandRepository.DeleteByIDAndOwner(id, ownerID)
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return exception.NewGlossaryProjectNotFoundError()
	}

	return nil
}
