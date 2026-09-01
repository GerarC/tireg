package usecase

import (
	"github.com/gerarc/tireg/internal/glossary/domain/api"
	"github.com/gerarc/tireg/internal/glossary/domain/exception"
	"github.com/gerarc/tireg/internal/glossary/domain/model"
	"github.com/gerarc/tireg/internal/glossary/domain/spi"
)

type UpdateGlossaryProjectUseCaseImplemented struct {
	glossaryProjectCommandRepository spi.GlossaryProjectCommandRepository
}

func NewUpdateGlossaryProjectUseCase(glossaryProjectCommandRepository spi.GlossaryProjectCommandRepository) api.UpdateGlossaryProjectUseCase {
	return &UpdateGlossaryProjectUseCaseImplemented{glossaryProjectCommandRepository: glossaryProjectCommandRepository}
}

func (updateGlossaryProjectUseCase *UpdateGlossaryProjectUseCaseImplemented) Update(ownerID string, id string, glossaryProject model.GlossaryProject) (model.GlossaryProject, error) {
	if details := validateGlossaryProject(glossaryProject); len(details) > 0 {
		return model.GlossaryProject{}, exception.NewValidationFailedError(details...)
	}

	return updateGlossaryProjectUseCase.glossaryProjectCommandRepository.UpdateByIDAndOwner(id, ownerID, glossaryProject)
}
