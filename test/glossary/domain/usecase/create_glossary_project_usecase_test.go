package usecase_test

import (
	"errors"
	"testing"

	sharedException "github.com/gerarc/tireg/internal/shared/domain/exception"

	"github.com/gerarc/tireg/internal/glossary/domain/model"
	"github.com/gerarc/tireg/internal/glossary/domain/usecase"
)

type stubGlossaryProjectCommandRepository struct {
	insertFunc             func(model.GlossaryProject) (model.GlossaryProject, error)
	updateByIDAndOwnerFunc func(string, string, model.GlossaryProject) (model.GlossaryProject, error)
	deleteByIDAndOwnerFunc func(string, string) error
}

func (stub *stubGlossaryProjectCommandRepository) Insert(glossaryProject model.GlossaryProject) (model.GlossaryProject, error) {
	return stub.insertFunc(glossaryProject)
}

func (stub *stubGlossaryProjectCommandRepository) UpdateByIDAndOwner(id string, ownerID string, glossaryProject model.GlossaryProject) (model.GlossaryProject, error) {
	return stub.updateByIDAndOwnerFunc(id, ownerID, glossaryProject)
}

func (stub *stubGlossaryProjectCommandRepository) DeleteByIDAndOwner(id string, ownerID string) error {
	return stub.deleteByIDAndOwnerFunc(id, ownerID)
}

type stubGlossaryProjectQueryRepository struct {
	selectAllByOwnerFunc func(string) ([]model.GlossaryProject, error)
}

func (stub *stubGlossaryProjectQueryRepository) SelectAllByOwner(ownerID string) ([]model.GlossaryProject, error) {
	return stub.selectAllByOwnerFunc(ownerID)
}

func TestCreate_Succeeds_WhenGlossaryProjectValid(t *testing.T) {
	commandRepository := &stubGlossaryProjectCommandRepository{
		insertFunc: func(glossaryProject model.GlossaryProject) (model.GlossaryProject, error) {
			glossaryProject.ID = "generated-id"
			return glossaryProject, nil
		},
	}

	createGlossaryProjectUseCase := usecase.NewCreateGlossaryProjectUseCase(commandRepository)

	glossaryProject, err := createGlossaryProjectUseCase.Create("owner-id", model.GlossaryProject{ProjectLabel: "Abastecar"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if glossaryProject.ID != "generated-id" || glossaryProject.OwnerID != "owner-id" {
		t.Fatalf("unexpected glossary project: %+v", glossaryProject)
	}
}

func TestCreate_ReturnsValidationError_WhenProjectLabelMissing(t *testing.T) {
	createGlossaryProjectUseCase := usecase.NewCreateGlossaryProjectUseCase(&stubGlossaryProjectCommandRepository{})

	_, err := createGlossaryProjectUseCase.Create("owner-id", model.GlossaryProject{})

	var domainError *sharedException.DomainError
	if !errors.As(err, &domainError) || domainError.Code != 400 {
		t.Fatalf("expected a 400 DomainError, got %v", err)
	}
}
