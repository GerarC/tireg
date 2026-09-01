package usecase_test

import (
	"testing"

	"github.com/gerarc/tireg/internal/task/domain/model"
	"github.com/gerarc/tireg/internal/task/domain/usecase"
)

func TestList_ReturnsMappingsOwnedByOwner(t *testing.T) {
	queryRepository := &stubTaskMappingQueryRepository{
		selectAllByOwnerFunc: func(ownerID string) ([]model.TaskMapping, error) {
			return []model.TaskMapping{{ID: "mapping-id", OwnerID: ownerID}}, nil
		},
	}

	listTaskMappingsUseCase := usecase.NewListTaskMappingsUseCase(queryRepository)

	taskMappings, err := listTaskMappingsUseCase.List("owner-id")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(taskMappings) != 1 || taskMappings[0].ID != "mapping-id" {
		t.Fatalf("unexpected task mappings: %+v", taskMappings)
	}
}

func TestList_ReturnsEmptySlice_WhenOwnerHasNoMappings(t *testing.T) {
	queryRepository := &stubTaskMappingQueryRepository{
		selectAllByOwnerFunc: func(ownerID string) ([]model.TaskMapping, error) {
			return nil, nil
		},
	}

	listTaskMappingsUseCase := usecase.NewListTaskMappingsUseCase(queryRepository)

	taskMappings, err := listTaskMappingsUseCase.List("owner-id")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(taskMappings) != 0 {
		t.Fatalf("expected no mappings, got %d", len(taskMappings))
	}
}
