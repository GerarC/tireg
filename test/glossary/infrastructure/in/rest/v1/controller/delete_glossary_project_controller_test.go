package controller_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	glossaryException "github.com/gerarc/tireg/internal/glossary/domain/exception"
	"github.com/gerarc/tireg/internal/glossary/infrastructure/in/rest/v1/controller"
)

type stubDeleteGlossaryProjectUseCase struct {
	deleteFunc func(string, string) error
}

func (stub *stubDeleteGlossaryProjectUseCase) Delete(ownerID string, id string) error {
	return stub.deleteFunc(ownerID, id)
}

func TestDeleteProject_Returns204_WhenOwnedByAuthenticatedUser(t *testing.T) {
	requireAuthMiddleware := authenticatedMiddleware("owner-id")
	useCase := &stubDeleteGlossaryProjectUseCase{deleteFunc: func(ownerID string, id string) error { return nil }}
	deleteGlossaryProjectController := controller.NewDeleteGlossaryProjectController(useCase, requireAuthMiddleware)

	request := authenticatedRequest(http.MethodDelete, "/api/v1/glossary/projects/project-id", nil)
	request.SetPathValue("id", "project-id")
	responseRecorder := httptest.NewRecorder()

	requireAuthMiddleware.Wrap(deleteGlossaryProjectController.Delete)(responseRecorder, request)

	if responseRecorder.Code != http.StatusNoContent {
		t.Fatalf("expected status %d, got %d", http.StatusNoContent, responseRecorder.Code)
	}
}

func TestDeleteProject_Returns404_WhenNotOwnedByAuthenticatedUser(t *testing.T) {
	requireAuthMiddleware := authenticatedMiddleware("owner-id")
	useCase := &stubDeleteGlossaryProjectUseCase{
		deleteFunc: func(ownerID string, id string) error {
			return glossaryException.NewGlossaryProjectNotFoundError()
		},
	}
	deleteGlossaryProjectController := controller.NewDeleteGlossaryProjectController(useCase, requireAuthMiddleware)

	request := authenticatedRequest(http.MethodDelete, "/api/v1/glossary/projects/someone-elses-id", nil)
	request.SetPathValue("id", "someone-elses-id")
	responseRecorder := httptest.NewRecorder()

	requireAuthMiddleware.Wrap(deleteGlossaryProjectController.Delete)(responseRecorder, request)

	if responseRecorder.Code != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, responseRecorder.Code)
	}
}
