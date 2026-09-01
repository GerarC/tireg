package controller_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	taskException "github.com/gerarc/tireg/internal/task/domain/exception"
	"github.com/gerarc/tireg/internal/task/domain/model"
	"github.com/gerarc/tireg/internal/task/infrastructure/in/rest/v1/controller"
	"github.com/gerarc/tireg/internal/task/infrastructure/in/rest/v1/dto"
)

type stubFindTaskMappingByIDUseCase struct {
	findByIDFunc func(string, string) (model.TaskMapping, error)
}

func (stub *stubFindTaskMappingByIDUseCase) FindByID(ownerID string, id string) (model.TaskMapping, error) {
	return stub.findByIDFunc(ownerID, id)
}

func TestFindByID_Returns200_WhenOwnedByAuthenticatedUser(t *testing.T) {
	requireAuthMiddleware := authenticatedMiddleware("owner-id")
	useCase := &stubFindTaskMappingByIDUseCase{
		findByIDFunc: func(ownerID string, id string) (model.TaskMapping, error) {
			return model.TaskMapping{ID: id, OwnerID: ownerID}, nil
		},
	}
	findTaskMappingByIDController := controller.NewFindTaskMappingByIDController(useCase, requireAuthMiddleware)

	request := authenticatedRequest(http.MethodGet, "/api/v1/task-mappings/mapping-id", nil)
	request.SetPathValue("id", "mapping-id")
	responseRecorder := httptest.NewRecorder()

	requireAuthMiddleware.Wrap(findTaskMappingByIDController.FindByID)(responseRecorder, request)

	if responseRecorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, responseRecorder.Code)
	}

	var responseDTO dto.TaskMappingResponseDTO
	if err := json.NewDecoder(responseRecorder.Body).Decode(&responseDTO); err != nil {
		t.Fatalf("unexpected error decoding response: %v", err)
	}

	if responseDTO.ID != "mapping-id" {
		t.Fatalf("unexpected response: %+v", responseDTO)
	}
}

func TestFindByID_Returns404_WhenNotOwnedByAuthenticatedUser(t *testing.T) {
	requireAuthMiddleware := authenticatedMiddleware("owner-id")
	useCase := &stubFindTaskMappingByIDUseCase{
		findByIDFunc: func(ownerID string, id string) (model.TaskMapping, error) {
			return model.TaskMapping{}, taskException.NewTaskMappingNotFoundError()
		},
	}
	findTaskMappingByIDController := controller.NewFindTaskMappingByIDController(useCase, requireAuthMiddleware)

	request := authenticatedRequest(http.MethodGet, "/api/v1/task-mappings/someone-elses-id", nil)
	request.SetPathValue("id", "someone-elses-id")
	responseRecorder := httptest.NewRecorder()

	requireAuthMiddleware.Wrap(findTaskMappingByIDController.FindByID)(responseRecorder, request)

	if responseRecorder.Code != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, responseRecorder.Code)
	}
}
