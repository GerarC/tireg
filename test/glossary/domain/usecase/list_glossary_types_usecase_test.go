package usecase_test

import (
	"testing"

	"github.com/gerarc/tireg/internal/glossary/domain/model"
	"github.com/gerarc/tireg/internal/glossary/domain/usecase"
)

func TestList_ReturnsExistingTypes_WithoutSeeding(t *testing.T) {
	insertCalls := 0
	commandRepository := &stubGlossaryTypeCommandRepository{
		insertFunc: func(glossaryType model.GlossaryType) (model.GlossaryType, error) {
			insertCalls++
			return glossaryType, nil
		},
	}
	queryRepository := &stubGlossaryTypeQueryRepository{
		selectAllByOwnerFunc: func(ownerID string) ([]model.GlossaryType, error) {
			return []model.GlossaryType{{ID: "existing-id", OwnerID: ownerID, TypeKey: "dev"}}, nil
		},
	}

	listGlossaryTypesUseCase := usecase.NewListGlossaryTypesUseCase(commandRepository, queryRepository)

	glossaryTypes, err := listGlossaryTypesUseCase.List("owner-id")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(glossaryTypes) != 1 || glossaryTypes[0].ID != "existing-id" {
		t.Fatalf("unexpected glossary types: %+v", glossaryTypes)
	}

	if insertCalls != 0 {
		t.Fatalf("expected no seeding when types already exist, got %d inserts", insertCalls)
	}
}

func TestList_SeedsDefaultTypes_WhenOwnerHasNone(t *testing.T) {
	var inserted []model.GlossaryType
	commandRepository := &stubGlossaryTypeCommandRepository{
		insertFunc: func(glossaryType model.GlossaryType) (model.GlossaryType, error) {
			glossaryType.ID = "seeded-" + glossaryType.TypeKey
			inserted = append(inserted, glossaryType)
			return glossaryType, nil
		},
	}
	queryRepository := &stubGlossaryTypeQueryRepository{
		selectAllByOwnerFunc: func(ownerID string) ([]model.GlossaryType, error) {
			return nil, nil
		},
	}

	listGlossaryTypesUseCase := usecase.NewListGlossaryTypesUseCase(commandRepository, queryRepository)

	glossaryTypes, err := listGlossaryTypesUseCase.List("owner-id")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(glossaryTypes) != 5 {
		t.Fatalf("expected 5 seeded default types, got %d", len(glossaryTypes))
	}

	if len(inserted) != 5 {
		t.Fatalf("expected 5 inserts, got %d", len(inserted))
	}

	for _, glossaryType := range inserted {
		if glossaryType.OwnerID != "owner-id" {
			t.Fatalf("expected seeded type to be owned by owner-id, got %+v", glossaryType)
		}
	}
}
