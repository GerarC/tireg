package controller_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gerarc/tireg/internal/task/domain/model"
	"github.com/gerarc/tireg/internal/task/infrastructure/in/rest/v1/controller"
	"github.com/gerarc/tireg/internal/task/infrastructure/in/rest/v1/dto"
)

type stubListTaskMappingsUseCase struct {
	listFunc func(string) ([]model.TaskMapping, error)
}

func (stub *stubListTaskMappingsUseCase) List(ownerID string) ([]model.TaskMapping, error) {
	return stub.listFunc(ownerID)
}

func TestList_Returns200_WithOwnersMappings(t *testing.T) {
	requireAuthMiddleware := authenticatedMiddleware("owner-id")
	useCase := &stubListTaskMappingsUseCase{
		listFunc: func(ownerID string) ([]model.TaskMapping, error) {
			return []model.TaskMapping{{ID: "mapping-id", OwnerID: ownerID}}, nil
		},
	}
	listTaskMappingsController := controller.NewListTaskMappingsController(useCase, requireAuthMiddleware)

	responseRecorder := httptest.NewRecorder()
	requireAuthMiddleware.Wrap(listTaskMappingsController.List)(responseRecorder, authenticatedRequest(http.MethodGet, "/api/v1/task-mappings", nil))

	if responseRecorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, responseRecorder.Code)
	}

	var responseDTOs []dto.TaskMappingResponseDTO
	if err := json.NewDecoder(responseRecorder.Body).Decode(&responseDTOs); err != nil {
		t.Fatalf("unexpected error decoding response: %v", err)
	}

	if len(responseDTOs) != 1 || responseDTOs[0].ID != "mapping-id" {
		t.Fatalf("unexpected response: %+v", responseDTOs)
	}
}

func TestList_Returns401_WhenTokenInvalid(t *testing.T) {
	requireAuthMiddleware := unauthenticatedMiddleware()
	listTaskMappingsController := controller.NewListTaskMappingsController(&stubListTaskMappingsUseCase{}, requireAuthMiddleware)

	responseRecorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/api/v1/task-mappings", nil)

	requireAuthMiddleware.Wrap(listTaskMappingsController.List)(responseRecorder, request)

	if responseRecorder.Code != http.StatusUnauthorized {
		t.Fatalf("expected status %d, got %d", http.StatusUnauthorized, responseRecorder.Code)
	}
}
