package usecase_test

import (
	"errors"
	"testing"

	sharedException "github.com/gerarc/tireg/internal/shared/domain/exception"

	"github.com/gerarc/tireg/internal/glossary/domain/exception"
	"github.com/gerarc/tireg/internal/glossary/domain/model"
	"github.com/gerarc/tireg/internal/glossary/domain/usecase"
)

func TestUpdate_Succeeds_WhenGlossaryProjectValid(t *testing.T) {
	commandRepository := &stubGlossaryProjectCommandRepository{
		updateByIDAndOwnerFunc: func(id string, ownerID string, glossaryProject model.GlossaryProject) (model.GlossaryProject, error) {
			glossaryProject.ID = id
			glossaryProject.OwnerID = ownerID
			return glossaryProject, nil
		},
	}

	updateGlossaryProjectUseCase := usecase.NewUpdateGlossaryProjectUseCase(commandRepository)

	glossaryProject, err := updateGlossaryProjectUseCase.Update("owner-id", "project-id", model.GlossaryProject{ProjectLabel: "Abastecar"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if glossaryProject.ID != "project-id" || glossaryProject.OwnerID != "owner-id" {
		t.Fatalf("unexpected glossary project: %+v", glossaryProject)
	}
}

func TestUpdate_ReturnsValidationError_WhenProjectLabelMissing(t *testing.T) {
	updateGlossaryProjectUseCase := usecase.NewUpdateGlossaryProjectUseCase(&stubGlossaryProjectCommandRepository{})

	_, err := updateGlossaryProjectUseCase.Update("owner-id", "project-id", model.GlossaryProject{})

	var domainError *sharedException.DomainError
	if !errors.As(err, &domainError) || domainError.Code != 400 {
		t.Fatalf("expected a 400 DomainError, got %v", err)
	}
}

func TestUpdateProject_ReturnsNotFoundError_WhenRepositoryReportsNotFound(t *testing.T) {
	commandRepository := &stubGlossaryProjectCommandRepository{
		updateByIDAndOwnerFunc: func(id string, ownerID string, glossaryProject model.GlossaryProject) (model.GlossaryProject, error) {
			return model.GlossaryProject{}, exception.NewGlossaryProjectNotFoundError()
		},
	}

	updateGlossaryProjectUseCase := usecase.NewUpdateGlossaryProjectUseCase(commandRepository)

	_, err := updateGlossaryProjectUseCase.Update("owner-id", "someone-elses-id", model.GlossaryProject{ProjectLabel: "Abastecar"})

	var domainError *sharedException.DomainError
	if !errors.As(err, &domainError) || domainError.Code != 404 {
		t.Fatalf("expected a 404 DomainError, got %v", err)
	}
}
