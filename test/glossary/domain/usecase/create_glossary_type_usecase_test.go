package usecase_test

import (
	"errors"
	"testing"

	sharedException "github.com/gerarc/tireg/internal/shared/domain/exception"

	"github.com/gerarc/tireg/internal/glossary/domain/model"
	"github.com/gerarc/tireg/internal/glossary/domain/usecase"
)

type stubGlossaryTypeCommandRepository struct {
	insertFunc             func(model.GlossaryType) (model.GlossaryType, error)
	updateByIDAndOwnerFunc func(string, string, model.GlossaryType) (model.GlossaryType, error)
	deleteByIDAndOwnerFunc func(string, string) error
}

func (stub *stubGlossaryTypeCommandRepository) Insert(glossaryType model.GlossaryType) (model.GlossaryType, error) {
	return stub.insertFunc(glossaryType)
}

func (stub *stubGlossaryTypeCommandRepository) UpdateByIDAndOwner(id string, ownerID string, glossaryType model.GlossaryType) (model.GlossaryType, error) {
	return stub.updateByIDAndOwnerFunc(id, ownerID, glossaryType)
}

func (stub *stubGlossaryTypeCommandRepository) DeleteByIDAndOwner(id string, ownerID string) error {
	return stub.deleteByIDAndOwnerFunc(id, ownerID)
}

type stubGlossaryTypeQueryRepository struct {
	selectAllByOwnerFunc func(string) ([]model.GlossaryType, error)
}

func (stub *stubGlossaryTypeQueryRepository) SelectAllByOwner(ownerID string) ([]model.GlossaryType, error) {
	return stub.selectAllByOwnerFunc(ownerID)
}

func TestCreate_Succeeds_WhenGlossaryTypeValid(t *testing.T) {
	commandRepository := &stubGlossaryTypeCommandRepository{
		insertFunc: func(glossaryType model.GlossaryType) (model.GlossaryType, error) {
			glossaryType.ID = "generated-id"
			return glossaryType, nil
		},
	}

	createGlossaryTypeUseCase := usecase.NewCreateGlossaryTypeUseCase(commandRepository)

	glossaryType, err := createGlossaryTypeUseCase.Create("owner-id", model.GlossaryType{TypeKey: "dev", Label: "Desarrollo"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if glossaryType.ID != "generated-id" || glossaryType.OwnerID != "owner-id" {
		t.Fatalf("unexpected glossary type: %+v", glossaryType)
	}
}

func TestCreate_ReturnsValidationError_WhenRequiredFieldsMissing(t *testing.T) {
	createGlossaryTypeUseCase := usecase.NewCreateGlossaryTypeUseCase(&stubGlossaryTypeCommandRepository{})

	_, err := createGlossaryTypeUseCase.Create("owner-id", model.GlossaryType{})

	var domainError *sharedException.DomainError
	if !errors.As(err, &domainError) {
		t.Fatalf("expected a DomainError, got %v", err)
	}

	if domainError.Code != 400 {
		t.Fatalf("expected code 400, got %d", domainError.Code)
	}

	if len(domainError.Details) != 2 {
		t.Fatalf("expected 2 validation details, got %d: %v", len(domainError.Details), domainError.Details)
	}
}
