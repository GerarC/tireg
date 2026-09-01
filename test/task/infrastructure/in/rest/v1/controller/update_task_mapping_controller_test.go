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

type stubUpdateTaskMappingUseCase struct {
	updateFunc func(string, string, model.TaskMapping) (model.TaskMapping, error)
}

func (stub *stubUpdateTaskMappingUseCase) Update(ownerID string, id string, taskMapping model.TaskMapping) (model.TaskMapping, error) {
	return stub.updateFunc(ownerID, id, taskMapping)
}

func TestUpdate_Returns200_WhenValid(t *testing.T) {
	requireAuthMiddleware := authenticatedMiddleware("owner-id")
	useCase := &stubUpdateTaskMappingUseCase{
		updateFunc: func(ownerID string, id string, taskMapping model.TaskMapping) (model.TaskMapping, error) {
			taskMapping.ID = id
			taskMapping.OwnerID = ownerID
			return taskMapping, nil
		},
	}
	updateTaskMappingController := controller.NewUpdateTaskMappingController(useCase, requireAuthMiddleware)

	requestBody, _ := json.Marshal(dto.TaskMappingRequestDTO{ProjectLabel: "Abastecar", Pattern: "reunión con el cliente"})
	request := authenticatedRequest(http.MethodPut, "/api/v1/task-mappings/mapping-id", requestBody)
	request.SetPathValue("id", "mapping-id")
	responseRecorder := httptest.NewRecorder()

	requireAuthMiddleware.Wrap(updateTaskMappingController.Update)(responseRecorder, request)

	if responseRecorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, responseRecorder.Code)
	}
}

func TestUpdate_Returns404_WhenNotOwnedByAuthenticatedUser(t *testing.T) {
	requireAuthMiddleware := authenticatedMiddleware("owner-id")
	useCase := &stubUpdateTaskMappingUseCase{
		updateFunc: func(ownerID string, id string, taskMapping model.TaskMapping) (model.TaskMapping, error) {
			return model.TaskMapping{}, taskException.NewTaskMappingNotFoundError()
		},
	}
	updateTaskMappingController := controller.NewUpdateTaskMappingController(useCase, requireAuthMiddleware)

	requestBody, _ := json.Marshal(dto.TaskMappingRequestDTO{ProjectLabel: "Abastecar", Pattern: "reunión con el cliente"})
	request := authenticatedRequest(http.MethodPut, "/api/v1/task-mappings/someone-elses-id", requestBody)
	request.SetPathValue("id", "someone-elses-id")
	responseRecorder := httptest.NewRecorder()

	requireAuthMiddleware.Wrap(updateTaskMappingController.Update)(responseRecorder, request)

	if responseRecorder.Code != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, responseRecorder.Code)
	}
}
