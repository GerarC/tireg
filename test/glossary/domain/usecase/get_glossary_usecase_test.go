package usecase_test

import (
	"errors"
	"testing"

	"github.com/gerarc/tireg/internal/glossary/domain/model"
	"github.com/gerarc/tireg/internal/glossary/domain/usecase"
)

type stubListGlossaryTypesUseCase struct {
	listFunc func(string) ([]model.GlossaryType, error)
}

func (stub *stubListGlossaryTypesUseCase) List(ownerID string) ([]model.GlossaryType, error) {
	return stub.listFunc(ownerID)
}

type stubListGlossaryProjectsUseCase struct {
	listFunc func(string) ([]model.GlossaryProject, error)
}

func (stub *stubListGlossaryProjectsUseCase) List(ownerID string) ([]model.GlossaryProject, error) {
	return stub.listFunc(ownerID)
}

func TestGet_ReturnsCombinedTypesAndProjects(t *testing.T) {
	listGlossaryTypesUseCase := &stubListGlossaryTypesUseCase{
		listFunc: func(ownerID string) ([]model.GlossaryType, error) {
			return []model.GlossaryType{{ID: "type-id", OwnerID: ownerID}}, nil
		},
	}
	listGlossaryProjectsUseCase := &stubListGlossaryProjectsUseCase{
		listFunc: func(ownerID string) ([]model.GlossaryProject, error) {
			return []model.GlossaryProject{{ID: "project-id", OwnerID: ownerID}}, nil
		},
	}

	getGlossaryUseCase := usecase.NewGetGlossaryUseCase(listGlossaryTypesUseCase, listGlossaryProjectsUseCase)

	glossary, err := getGlossaryUseCase.Get("owner-id")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(glossary.Types) != 1 || glossary.Types[0].ID != "type-id" {
		t.Fatalf("unexpected types: %+v", glossary.Types)
	}

	if len(glossary.Projects) != 1 || glossary.Projects[0].ID != "project-id" {
		t.Fatalf("unexpected projects: %+v", glossary.Projects)
	}
}

func TestGet_PropagatesError_WhenListingTypesFails(t *testing.T) {
	typesErr := errors.New("query failed")
	listGlossaryTypesUseCase := &stubListGlossaryTypesUseCase{
		listFunc: func(ownerID string) ([]model.GlossaryType, error) {
			return nil, typesErr
		},
	}
	listGlossaryProjectsUseCase := &stubListGlossaryProjectsUseCase{
		listFunc: func(ownerID string) ([]model.GlossaryProject, error) {
			return nil, nil
		},
	}

	getGlossaryUseCase := usecase.NewGetGlossaryUseCase(listGlossaryTypesUseCase, listGlossaryProjectsUseCase)

	_, err := getGlossaryUseCase.Get("owner-id")
	if !errors.Is(err, typesErr) {
		t.Fatalf("expected types error to propagate, got %v", err)
	}
}

func TestGet_PropagatesError_WhenListingProjectsFails(t *testing.T) {
	projectsErr := errors.New("query failed")
	listGlossaryTypesUseCase := &stubListGlossaryTypesUseCase{
		listFunc: func(ownerID string) ([]model.GlossaryType, error) {
			return nil, nil
		},
	}
	listGlossaryProjectsUseCase := &stubListGlossaryProjectsUseCase{
		listFunc: func(ownerID string) ([]model.GlossaryProject, error) {
			return nil, projectsErr
		},
	}

	getGlossaryUseCase := usecase.NewGetGlossaryUseCase(listGlossaryTypesUseCase, listGlossaryProjectsUseCase)

	_, err := getGlossaryUseCase.Get("owner-id")
	if !errors.Is(err, projectsErr) {
		t.Fatalf("expected projects error to propagate, got %v", err)
	}
}
