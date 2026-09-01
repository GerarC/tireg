package usecase

import (
	"github.com/gerarc/tireg/internal/glossary/domain/api"
	"github.com/gerarc/tireg/internal/glossary/domain/model"
	"github.com/gerarc/tireg/internal/glossary/domain/spi"
	"github.com/gerarc/tireg/internal/glossary/domain/util/constant"
)

type ListGlossaryTypesUseCaseImplemented struct {
	glossaryTypeCommandRepository spi.GlossaryTypeCommandRepository
	glossaryTypeQueryRepository   spi.GlossaryTypeQueryRepository
}

func NewListGlossaryTypesUseCase(glossaryTypeCommandRepository spi.GlossaryTypeCommandRepository, glossaryTypeQueryRepository spi.GlossaryTypeQueryRepository) api.ListGlossaryTypesUseCase {
	return &ListGlossaryTypesUseCaseImplemented{
		glossaryTypeCommandRepository: glossaryTypeCommandRepository,
		glossaryTypeQueryRepository:   glossaryTypeQueryRepository,
	}
}

func (listGlossaryTypesUseCase *ListGlossaryTypesUseCaseImplemented) List(ownerID string) ([]model.GlossaryType, error) {
	glossaryTypes, err := listGlossaryTypesUseCase.glossaryTypeQueryRepository.SelectAllByOwner(ownerID)
	if err != nil {
		return nil, err
	}

	if len(glossaryTypes) > 0 {
		return glossaryTypes, nil
	}

	return listGlossaryTypesUseCase.seedDefaultTypes(ownerID)
}

func (listGlossaryTypesUseCase *ListGlossaryTypesUseCaseImplemented) seedDefaultTypes(ownerID string) ([]model.GlossaryType, error) {
	seeded := make([]model.GlossaryType, 0, len(constant.DefaultGlossaryTypes))

	for _, defaultType := range constant.DefaultGlossaryTypes {
		inserted, err := listGlossaryTypesUseCase.glossaryTypeCommandRepository.Insert(model.GlossaryType{
			OwnerID:     ownerID,
			TypeKey:     defaultType.TypeKey,
			Label:       defaultType.Label,
			Description: defaultType.Description,
		})
		if err != nil {
			return nil, err
		}

		seeded = append(seeded, inserted)
	}

	return seeded, nil
}
