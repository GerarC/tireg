package usecase

import (
	"strings"

	"github.com/gerarc/tireg/internal/glossary/domain/api"
	"github.com/gerarc/tireg/internal/glossary/domain/exception"
	"github.com/gerarc/tireg/internal/glossary/domain/model"
	"github.com/gerarc/tireg/internal/glossary/domain/spi"
	"github.com/gerarc/tireg/internal/glossary/domain/util/constant"
)

type CreateGlossaryProjectUseCaseImplemented struct {
	glossaryProjectCommandRepository spi.GlossaryProjectCommandRepository
}

func NewCreateGlossaryProjectUseCase(glossaryProjectCommandRepository spi.GlossaryProjectCommandRepository) api.CreateGlossaryProjectUseCase {
	return &CreateGlossaryProjectUseCaseImplemented{glossaryProjectCommandRepository: glossaryProjectCommandRepository}
}

func (createGlossaryProjectUseCase *CreateGlossaryProjectUseCaseImplemented) Create(ownerID string, glossaryProject model.GlossaryProject) (model.GlossaryProject, error) {
	if details := validateGlossaryProject(glossaryProject); len(details) > 0 {
		return model.GlossaryProject{}, exception.NewValidationFailedError(details...)
	}

	glossaryProject.OwnerID = ownerID

	return createGlossaryProjectUseCase.glossaryProjectCommandRepository.Insert(glossaryProject)
}

func validateGlossaryProject(glossaryProject model.GlossaryProject) []string {
	var details []string

	if strings.TrimSpace(glossaryProject.ProjectLabel) == "" {
		details = append(details, constant.DetailProjectLabelRequired)
	}

	return details
}
