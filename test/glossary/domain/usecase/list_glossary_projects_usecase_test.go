package usecase_test

import (
	"testing"

	"github.com/gerarc/tireg/internal/glossary/domain/model"
	"github.com/gerarc/tireg/internal/glossary/domain/usecase"
)

func TestList_ReturnsProjectsOwnedByOwner(t *testing.T) {
	queryRepository := &stubGlossaryProjectQueryRepository{
		selectAllByOwnerFunc: func(ownerID string) ([]model.GlossaryProject, error) {
			return []model.GlossaryProject{{ID: "project-id", OwnerID: ownerID, ProjectLabel: "Abastecar"}}, nil
		},
	}

	listGlossaryProjectsUseCase := usecase.NewListGlossaryProjectsUseCase(queryRepository)

	glossaryProjects, err := listGlossaryProjectsUseCase.List("owner-id")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(glossaryProjects) != 1 || glossaryProjects[0].ID != "project-id" {
		t.Fatalf("unexpected glossary projects: %+v", glossaryProjects)
	}
}

func TestList_ReturnsEmptySlice_WhenOwnerHasNoProjects(t *testing.T) {
	queryRepository := &stubGlossaryProjectQueryRepository{
		selectAllByOwnerFunc: func(ownerID string) ([]model.GlossaryProject, error) {
			return nil, nil
		},
	}

	listGlossaryProjectsUseCase := usecase.NewListGlossaryProjectsUseCase(queryRepository)

	glossaryProjects, err := listGlossaryProjectsUseCase.List("owner-id")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(glossaryProjects) != 0 {
		t.Fatalf("expected no projects, got %d", len(glossaryProjects))
	}
}
