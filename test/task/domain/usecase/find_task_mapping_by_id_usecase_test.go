package usecase_test

import (
	"errors"
	"testing"

	sharedException "github.com/gerarc/tireg/internal/shared/domain/exception"

	"github.com/gerarc/tireg/internal/task/domain/exception"
	"github.com/gerarc/tireg/internal/task/domain/model"
	"github.com/gerarc/tireg/internal/task/domain/usecase"
)

func TestFindByID_ReturnsMapping_WhenOwnedByOwner(t *testing.T) {
	queryRepository := &stubTaskMappingQueryRepository{
		selectByIDAndOwnerFunc: func(id string, ownerID string) (model.TaskMapping, error) {
			return model.TaskMapping{ID: id, OwnerID: ownerID}, nil
		},
	}

	findTaskMappingByIDUseCase := usecase.NewFindTaskMappingByIDUseCase(queryRepository)

	taskMapping, err := findTaskMappingByIDUseCase.FindByID("owner-id", "mapping-id")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if taskMapping.ID != "mapping-id" || taskMapping.OwnerID != "owner-id" {
		t.Fatalf("unexpected task mapping: %+v", taskMapping)
	}
}

func TestFindByID_ReturnsNotFoundError_WhenNotOwnedByOwner(t *testing.T) {
	queryRepository := &stubTaskMappingQueryRepository{
		selectByIDAndOwnerFunc: func(id string, ownerID string) (model.TaskMapping, error) {
			return model.TaskMapping{}, exception.NewTaskMappingNotFoundError()
		},
	}

	findTaskMappingByIDUseCase := usecase.NewFindTaskMappingByIDUseCase(queryRepository)

	_, err := findTaskMappingByIDUseCase.FindByID("owner-id", "someone-elses-id")

	var domainError *sharedException.DomainError
	if !errors.As(err, &domainError) || domainError.Code != 404 {
		t.Fatalf("expected a 404 DomainError, got %v", err)
	}
}
