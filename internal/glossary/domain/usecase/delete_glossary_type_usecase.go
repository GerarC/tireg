package usecase

import (
	"github.com/gerarc/tireg/internal/glossary/domain/api"
	"github.com/gerarc/tireg/internal/glossary/domain/spi"
)

type DeleteGlossaryTypeUseCaseImplemented struct {
	glossaryTypeCommandRepository spi.GlossaryTypeCommandRepository
}

func NewDeleteGlossaryTypeUseCase(glossaryTypeCommandRepository spi.GlossaryTypeCommandRepository) api.DeleteGlossaryTypeUseCase {
	return &DeleteGlossaryTypeUseCaseImplemented{glossaryTypeCommandRepository: glossaryTypeCommandRepository}
}

func (deleteGlossaryTypeUseCase *DeleteGlossaryTypeUseCaseImplemented) Delete(ownerID string, id string) error {
	return deleteGlossaryTypeUseCase.glossaryTypeCommandRepository.DeleteByIDAndOwner(id, ownerID)
}
