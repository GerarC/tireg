package usecase_test

import (
	"errors"
	"testing"

	sharedException "github.com/gerarc/tireg/internal/shared/domain/exception"

	"github.com/gerarc/tireg/internal/task/domain/exception"
	"github.com/gerarc/tireg/internal/task/domain/model"
	"github.com/gerarc/tireg/internal/task/domain/usecase"
)

func TestUpdate_Succeeds_WhenTaskMappingValid(t *testing.T) {
	commandRepository := &stubTaskMappingCommandRepository{
		updateByIDAndOwnerFunc: func(id string, ownerID string, taskMapping model.TaskMapping) (model.TaskMapping, error) {
			taskMapping.ID = id
			taskMapping.OwnerID = ownerID
			return taskMapping, nil
		},
	}

	updateTaskMappingUseCase := usecase.NewUpdateTaskMappingUseCase(commandRepository)

	taskMapping, err := updateTaskMappingUseCase.Update("owner-id", "mapping-id", model.TaskMapping{ProjectLabel: "Abastecar", Pattern: "reunión con el cliente"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if taskMapping.ID != "mapping-id" || taskMapping.OwnerID != "owner-id" {
		t.Fatalf("unexpected task mapping: %+v", taskMapping)
	}
}

func TestUpdate_ReturnsValidationError_WhenRequiredFieldsMissing(t *testing.T) {
	updateTaskMappingUseCase := usecase.NewUpdateTaskMappingUseCase(&stubTaskMappingCommandRepository{})

	_, err := updateTaskMappingUseCase.Update("owner-id", "mapping-id", model.TaskMapping{})

	var domainError *sharedException.DomainError
	if !errors.As(err, &domainError) || domainError.Code != 400 {
		t.Fatalf("expected a 400 DomainError, got %v", err)
	}
}

func TestUpdate_ReturnsNotFoundError_WhenNotOwnedByOwner(t *testing.T) {
	commandRepository := &stubTaskMappingCommandRepository{
		updateByIDAndOwnerFunc: func(id string, ownerID string, taskMapping model.TaskMapping) (model.TaskMapping, error) {
			return model.TaskMapping{}, exception.NewTaskMappingNotFoundError()
		},
	}

	updateTaskMappingUseCase := usecase.NewUpdateTaskMappingUseCase(commandRepository)

	_, err := updateTaskMappingUseCase.Update("owner-id", "someone-elses-id", model.TaskMapping{ProjectLabel: "Abastecar", Pattern: "reunión con el cliente"})

	var domainError *sharedException.DomainError
	if !errors.As(err, &domainError) || domainError.Code != 404 {
		t.Fatalf("expected a 404 DomainError, got %v", err)
	}
}
