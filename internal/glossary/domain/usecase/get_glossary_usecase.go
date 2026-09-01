package usecase

import (
	"github.com/gerarc/tireg/internal/glossary/domain/api"
	"github.com/gerarc/tireg/internal/glossary/domain/model"
)

type GetGlossaryUseCaseImplemented struct {
	listGlossaryTypesUseCase    api.ListGlossaryTypesUseCase
	listGlossaryProjectsUseCase api.ListGlossaryProjectsUseCase
}

func NewGetGlossaryUseCase(listGlossaryTypesUseCase api.ListGlossaryTypesUseCase, listGlossaryProjectsUseCase api.ListGlossaryProjectsUseCase) api.GetGlossaryUseCase {
	return &GetGlossaryUseCaseImplemented{
		listGlossaryTypesUseCase:    listGlossaryTypesUseCase,
		listGlossaryProjectsUseCase: listGlossaryProjectsUseCase,
	}
}

func (getGlossaryUseCase *GetGlossaryUseCaseImplemented) Get(ownerID string) (model.Glossary, error) {
	glossaryTypes, err := getGlossaryUseCase.listGlossaryTypesUseCase.List(ownerID)
	if err != nil {
		return model.Glossary{}, err
	}

	glossaryProjects, err := getGlossaryUseCase.listGlossaryProjectsUseCase.List(ownerID)
	if err != nil {
		return model.Glossary{}, err
	}

	return model.Glossary{Types: glossaryTypes, Projects: glossaryProjects}, nil
}
