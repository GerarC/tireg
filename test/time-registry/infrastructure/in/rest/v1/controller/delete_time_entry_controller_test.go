package controller_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	timeEntryException "github.com/gerarc/tireg/internal/time-registry/domain/exception"
	"github.com/gerarc/tireg/internal/time-registry/infrastructure/in/rest/v1/controller"
)

type stubDeleteTimeEntryUseCase struct {
	deleteFunc func(string, string) error
}

func (stub *stubDeleteTimeEntryUseCase) Delete(ownerID string, id string) error {
	return stub.deleteFunc(ownerID, id)
}

func TestDelete_Returns204_WhenOwnedByAuthenticatedUser(t *testing.T) {
	requireAuthMiddleware := authenticatedMiddleware("owner-id")
	useCase := &stubDeleteTimeEntryUseCase{deleteFunc: func(ownerID string, id string) error { return nil }}
	deleteTimeEntryController := controller.NewDeleteTimeEntryController(useCase, requireAuthMiddleware)

	request := authenticatedRequest(http.MethodDelete, "/api/v1/time-entries/entry-id", nil)
	request.SetPathValue("id", "entry-id")
	responseRecorder := httptest.NewRecorder()

	requireAuthMiddleware.Wrap(deleteTimeEntryController.Delete)(responseRecorder, request)

	if responseRecorder.Code != http.StatusNoContent {
		t.Fatalf("expected status %d, got %d", http.StatusNoContent, responseRecorder.Code)
	}
}

func TestDelete_Returns404_WhenNotOwnedByAuthenticatedUser(t *testing.T) {
	requireAuthMiddleware := authenticatedMiddleware("owner-id")
	useCase := &stubDeleteTimeEntryUseCase{
		deleteFunc: func(ownerID string, id string) error {
			return timeEntryException.NewTimeEntryNotFoundError()
		},
	}
	deleteTimeEntryController := controller.NewDeleteTimeEntryController(useCase, requireAuthMiddleware)

	request := authenticatedRequest(http.MethodDelete, "/api/v1/time-entries/someone-elses-id", nil)
	request.SetPathValue("id", "someone-elses-id")
	responseRecorder := httptest.NewRecorder()

	requireAuthMiddleware.Wrap(deleteTimeEntryController.Delete)(responseRecorder, request)

	if responseRecorder.Code != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, responseRecorder.Code)
	}
}
