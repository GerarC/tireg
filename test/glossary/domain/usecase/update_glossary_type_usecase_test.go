package usecase_test

import (
	"errors"
	"testing"

	sharedException "github.com/gerarc/tireg/internal/shared/domain/exception"

	"github.com/gerarc/tireg/internal/glossary/domain/exception"
	"github.com/gerarc/tireg/internal/glossary/domain/model"
	"github.com/gerarc/tireg/internal/glossary/domain/usecase"
)

func TestUpdate_Succeeds_WhenGlossaryTypeValid(t *testing.T) {
	commandRepository := &stubGlossaryTypeCommandRepository{
		updateByIDAndOwnerFunc: func(id string, ownerID string, glossaryType model.GlossaryType) (model.GlossaryType, error) {
			glossaryType.ID = id
			glossaryType.OwnerID = ownerID
			return glossaryType, nil
		},
	}

	updateGlossaryTypeUseCase := usecase.NewUpdateGlossaryTypeUseCase(commandRepository)

	glossaryType, err := updateGlossaryTypeUseCase.Update("owner-id", "type-id", model.GlossaryType{TypeKey: "dev", Label: "Desarrollo"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if glossaryType.ID != "type-id" || glossaryType.OwnerID != "owner-id" {
		t.Fatalf("unexpected glossary type: %+v", glossaryType)
	}
}

func TestUpdate_ReturnsValidationError_WhenRequiredFieldsMissing(t *testing.T) {
	updateGlossaryTypeUseCase := usecase.NewUpdateGlossaryTypeUseCase(&stubGlossaryTypeCommandRepository{})

	_, err := updateGlossaryTypeUseCase.Update("owner-id", "type-id", model.GlossaryType{})

	var domainError *sharedException.DomainError
	if !errors.As(err, &domainError) || domainError.Code != 400 {
		t.Fatalf("expected a 400 DomainError, got %v", err)
	}
}

func TestUpdate_ReturnsNotFoundError_WhenRepositoryReportsNotFound(t *testing.T) {
	commandRepository := &stubGlossaryTypeCommandRepository{
		updateByIDAndOwnerFunc: func(id string, ownerID string, glossaryType model.GlossaryType) (model.GlossaryType, error) {
			return model.GlossaryType{}, exception.NewGlossaryTypeNotFoundError()
		},
	}

	updateGlossaryTypeUseCase := usecase.NewUpdateGlossaryTypeUseCase(commandRepository)

	_, err := updateGlossaryTypeUseCase.Update("owner-id", "someone-elses-id", model.GlossaryType{TypeKey: "dev", Label: "Desarrollo"})

	var domainError *sharedException.DomainError
	if !errors.As(err, &domainError) || domainError.Code != 404 {
		t.Fatalf("expected a 404 DomainError, got %v", err)
	}
}
