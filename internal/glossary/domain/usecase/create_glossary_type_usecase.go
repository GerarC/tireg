package usecase

import (
	"strings"

	"github.com/gerarc/tireg/internal/glossary/domain/api"
	"github.com/gerarc/tireg/internal/glossary/domain/exception"
	"github.com/gerarc/tireg/internal/glossary/domain/model"
	"github.com/gerarc/tireg/internal/glossary/domain/spi"
	"github.com/gerarc/tireg/internal/glossary/domain/util/constant"
)

type CreateGlossaryTypeUseCaseImplemented struct {
	glossaryTypeCommandRepository spi.GlossaryTypeCommandRepository
}

func NewCreateGlossaryTypeUseCase(glossaryTypeCommandRepository spi.GlossaryTypeCommandRepository) api.CreateGlossaryTypeUseCase {
	return &CreateGlossaryTypeUseCaseImplemented{glossaryTypeCommandRepository: glossaryTypeCommandRepository}
}

func (createGlossaryTypeUseCase *CreateGlossaryTypeUseCaseImplemented) Create(ownerID string, glossaryType model.GlossaryType) (model.GlossaryType, error) {
	if details := validateGlossaryType(glossaryType); len(details) > 0 {
		return model.GlossaryType{}, exception.NewValidationFailedError(details...)
	}

	glossaryType.OwnerID = ownerID

	return createGlossaryTypeUseCase.glossaryTypeCommandRepository.Insert(glossaryType)
}

func validateGlossaryType(glossaryType model.GlossaryType) []string {
	var details []string

	if strings.TrimSpace(glossaryType.TypeKey) == "" {
		details = append(details, constant.DetailTypeKeyRequired)
	}

	if strings.TrimSpace(glossaryType.Label) == "" {
		details = append(details, constant.DetailLabelRequired)
	}

	return details
}
