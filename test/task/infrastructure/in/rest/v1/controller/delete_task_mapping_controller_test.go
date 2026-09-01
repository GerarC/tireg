package controller_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	taskException "github.com/gerarc/tireg/internal/task/domain/exception"
	"github.com/gerarc/tireg/internal/task/infrastructure/in/rest/v1/controller"
)

type stubDeleteTaskMappingUseCase struct {
	deleteFunc func(string, string) error
}

func (stub *stubDeleteTaskMappingUseCase) Delete(ownerID string, id string) error {
	return stub.deleteFunc(ownerID, id)
}

func TestDelete_Returns204_WhenOwnedByAuthenticatedUser(t *testing.T) {
	requireAuthMiddleware := authenticatedMiddleware("owner-id")
	useCase := &stubDeleteTaskMappingUseCase{deleteFunc: func(ownerID string, id string) error { return nil }}
	deleteTaskMappingController := controller.NewDeleteTaskMappingController(useCase, requireAuthMiddleware)

	request := authenticatedRequest(http.MethodDelete, "/api/v1/task-mappings/mapping-id", nil)
	request.SetPathValue("id", "mapping-id")
	responseRecorder := httptest.NewRecorder()

	requireAuthMiddleware.Wrap(deleteTaskMappingController.Delete)(responseRecorder, request)

	if responseRecorder.Code != http.StatusNoContent {
		t.Fatalf("expected status %d, got %d", http.StatusNoContent, responseRecorder.Code)
	}
}

func TestDelete_Returns404_WhenNotOwnedByAuthenticatedUser(t *testing.T) {
	requireAuthMiddleware := authenticatedMiddleware("owner-id")
	useCase := &stubDeleteTaskMappingUseCase{
		deleteFunc: func(ownerID string, id string) error {
			return taskException.NewTaskMappingNotFoundError()
		},
	}
	deleteTaskMappingController := controller.NewDeleteTaskMappingController(useCase, requireAuthMiddleware)

	request := authenticatedRequest(http.MethodDelete, "/api/v1/task-mappings/someone-elses-id", nil)
	request.SetPathValue("id", "someone-elses-id")
	responseRecorder := httptest.NewRecorder()

	requireAuthMiddleware.Wrap(deleteTaskMappingController.Delete)(responseRecorder, request)

	if responseRecorder.Code != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, responseRecorder.Code)
	}
}
