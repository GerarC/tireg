package usecase

import (
	"github.com/gerarc/tireg/internal/glossary/domain/api"
	"github.com/gerarc/tireg/internal/glossary/domain/spi"
)

type DeleteGlossaryProjectUseCaseImplemented struct {
	glossaryProjectCommandRepository spi.GlossaryProjectCommandRepository
}

func NewDeleteGlossaryProjectUseCase(glossaryProjectCommandRepository spi.GlossaryProjectCommandRepository) api.DeleteGlossaryProjectUseCase {
	return &DeleteGlossaryProjectUseCaseImplemented{glossaryProjectCommandRepository: glossaryProjectCommandRepository}
}

func (deleteGlossaryProjectUseCase *DeleteGlossaryProjectUseCaseImplemented) Delete(ownerID string, id string) error {
	return deleteGlossaryProjectUseCase.glossaryProjectCommandRepository.DeleteByIDAndOwner(id, ownerID)
}
