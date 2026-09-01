package usecase

import (
	"github.com/gerarc/tireg/internal/glossary/domain/api"
	"github.com/gerarc/tireg/internal/glossary/domain/exception"
	"github.com/gerarc/tireg/internal/glossary/domain/model"
	"github.com/gerarc/tireg/internal/glossary/domain/spi"
)

type UpdateGlossaryTypeUseCaseImplemented struct {
	glossaryTypeCommandRepository spi.GlossaryTypeCommandRepository
}

func NewUpdateGlossaryTypeUseCase(glossaryTypeCommandRepository spi.GlossaryTypeCommandRepository) api.UpdateGlossaryTypeUseCase {
	return &UpdateGlossaryTypeUseCaseImplemented{glossaryTypeCommandRepository: glossaryTypeCommandRepository}
}

func (updateGlossaryTypeUseCase *UpdateGlossaryTypeUseCaseImplemented) Update(ownerID string, id string, glossaryType model.GlossaryType) (model.GlossaryType, error) {
	if details := validateGlossaryType(glossaryType); len(details) > 0 {
		return model.GlossaryType{}, exception.NewValidationFailedError(details...)
	}

	return updateGlossaryTypeUseCase.glossaryTypeCommandRepository.UpdateByIDAndOwner(id, ownerID, glossaryType)
}
