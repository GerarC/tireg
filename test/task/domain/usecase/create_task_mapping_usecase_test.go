package usecase_test

import (
	"errors"
	"testing"

	sharedException "github.com/gerarc/tireg/internal/shared/domain/exception"

	"github.com/gerarc/tireg/internal/task/domain/model"
	"github.com/gerarc/tireg/internal/task/domain/usecase"
)

type stubTaskMappingCommandRepository struct {
	insertFunc             func(model.TaskMapping) (model.TaskMapping, error)
	updateByIDAndOwnerFunc func(string, string, model.TaskMapping) (model.TaskMapping, error)
	deleteByIDAndOwnerFunc func(string, string) error
}

func (stub *stubTaskMappingCommandRepository) Insert(taskMapping model.TaskMapping) (model.TaskMapping, error) {
	return stub.insertFunc(taskMapping)
}

func (stub *stubTaskMappingCommandRepository) UpdateByIDAndOwner(id string, ownerID string, taskMapping model.TaskMapping) (model.TaskMapping, error) {
	return stub.updateByIDAndOwnerFunc(id, ownerID, taskMapping)
}

func (stub *stubTaskMappingCommandRepository) DeleteByIDAndOwner(id string, ownerID string) error {
	return stub.deleteByIDAndOwnerFunc(id, ownerID)
}

type stubTaskMappingQueryRepository struct {
	selectAllByOwnerFunc   func(string) ([]model.TaskMapping, error)
	selectByIDAndOwnerFunc func(string, string) (model.TaskMapping, error)
}

func (stub *stubTaskMappingQueryRepository) SelectAllByOwner(ownerID string) ([]model.TaskMapping, error) {
	return stub.selectAllByOwnerFunc(ownerID)
}

func (stub *stubTaskMappingQueryRepository) SelectByIDAndOwner(id string, ownerID string) (model.TaskMapping, error) {
	return stub.selectByIDAndOwnerFunc(id, ownerID)
}

func TestCreate_Succeeds_WhenTaskMappingValid(t *testing.T) {
	commandRepository := &stubTaskMappingCommandRepository{
		insertFunc: func(taskMapping model.TaskMapping) (model.TaskMapping, error) {
			taskMapping.ID = "generated-id"
			return taskMapping, nil
		},
	}

	createTaskMappingUseCase := usecase.NewCreateTaskMappingUseCase(commandRepository)

	taskMapping, err := createTaskMappingUseCase.Create("owner-id", model.TaskMapping{ProjectLabel: "Abastecar", Pattern: "reunión con el cliente"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if taskMapping.ID != "generated-id" || taskMapping.OwnerID != "owner-id" {
		t.Fatalf("unexpected task mapping: %+v", taskMapping)
	}
}

func TestCreate_ReturnsValidationError_WhenRequiredFieldsMissing(t *testing.T) {
	createTaskMappingUseCase := usecase.NewCreateTaskMappingUseCase(&stubTaskMappingCommandRepository{})

	_, err := createTaskMappingUseCase.Create("owner-id", model.TaskMapping{})

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
