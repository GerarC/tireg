package controller_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	glossaryException "github.com/gerarc/tireg/internal/glossary/domain/exception"
	"github.com/gerarc/tireg/internal/glossary/infrastructure/in/rest/v1/controller"
)

type stubDeleteGlossaryTypeUseCase struct {
	deleteFunc func(string, string) error
}

func (stub *stubDeleteGlossaryTypeUseCase) Delete(ownerID string, id string) error {
	return stub.deleteFunc(ownerID, id)
}

func TestDelete_Returns204_WhenOwnedByAuthenticatedUser(t *testing.T) {
	requireAuthMiddleware := authenticatedMiddleware("owner-id")
	useCase := &stubDeleteGlossaryTypeUseCase{deleteFunc: func(ownerID string, id string) error { return nil }}
	deleteGlossaryTypeController := controller.NewDeleteGlossaryTypeController(useCase, requireAuthMiddleware)

	request := authenticatedRequest(http.MethodDelete, "/api/v1/glossary/types/type-id", nil)
	request.SetPathValue("id", "type-id")
	responseRecorder := httptest.NewRecorder()

	requireAuthMiddleware.Wrap(deleteGlossaryTypeController.Delete)(responseRecorder, request)

	if responseRecorder.Code != http.StatusNoContent {
		t.Fatalf("expected status %d, got %d", http.StatusNoContent, responseRecorder.Code)
	}
}

func TestDelete_Returns404_WhenNotOwnedByAuthenticatedUser(t *testing.T) {
	requireAuthMiddleware := authenticatedMiddleware("owner-id")
	useCase := &stubDeleteGlossaryTypeUseCase{
		deleteFunc: func(ownerID string, id string) error {
			return glossaryException.NewGlossaryTypeNotFoundError()
		},
	}
	deleteGlossaryTypeController := controller.NewDeleteGlossaryTypeController(useCase, requireAuthMiddleware)

	request := authenticatedRequest(http.MethodDelete, "/api/v1/glossary/types/someone-elses-id", nil)
	request.SetPathValue("id", "someone-elses-id")
	responseRecorder := httptest.NewRecorder()

	requireAuthMiddleware.Wrap(deleteGlossaryTypeController.Delete)(responseRecorder, request)

	if responseRecorder.Code != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, responseRecorder.Code)
	}
}
