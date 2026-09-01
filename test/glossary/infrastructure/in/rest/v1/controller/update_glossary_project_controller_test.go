package controller_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	glossaryException "github.com/gerarc/tireg/internal/glossary/domain/exception"
	"github.com/gerarc/tireg/internal/glossary/domain/model"
	"github.com/gerarc/tireg/internal/glossary/infrastructure/in/rest/v1/controller"
	"github.com/gerarc/tireg/internal/glossary/infrastructure/in/rest/v1/dto"
)

type stubUpdateGlossaryProjectUseCase struct {
	updateFunc func(string, string, model.GlossaryProject) (model.GlossaryProject, error)
}

func (stub *stubUpdateGlossaryProjectUseCase) Update(ownerID string, id string, glossaryProject model.GlossaryProject) (model.GlossaryProject, error) {
	return stub.updateFunc(ownerID, id, glossaryProject)
}

func TestUpdateProject_Returns200_WhenValid(t *testing.T) {
	requireAuthMiddleware := authenticatedMiddleware("owner-id")
	useCase := &stubUpdateGlossaryProjectUseCase{
		updateFunc: func(ownerID string, id string, glossaryProject model.GlossaryProject) (model.GlossaryProject, error) {
			glossaryProject.ID = id
			glossaryProject.OwnerID = ownerID
			return glossaryProject, nil
		},
	}
	updateGlossaryProjectController := controller.NewUpdateGlossaryProjectController(useCase, requireAuthMiddleware)

	requestBody, _ := json.Marshal(dto.GlossaryProjectRequestDTO{ProjectLabel: "Abastecar"})
	request := authenticatedRequest(http.MethodPut, "/api/v1/glossary/projects/project-id", requestBody)
	request.SetPathValue("id", "project-id")
	responseRecorder := httptest.NewRecorder()

	requireAuthMiddleware.Wrap(updateGlossaryProjectController.Update)(responseRecorder, request)

	if responseRecorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, responseRecorder.Code)
	}
}

func TestUpdateProject_Returns404_WhenNotOwnedByAuthenticatedUser(t *testing.T) {
	requireAuthMiddleware := authenticatedMiddleware("owner-id")
	useCase := &stubUpdateGlossaryProjectUseCase{
		updateFunc: func(ownerID string, id string, glossaryProject model.GlossaryProject) (model.GlossaryProject, error) {
			return model.GlossaryProject{}, glossaryException.NewGlossaryProjectNotFoundError()
		},
	}
	updateGlossaryProjectController := controller.NewUpdateGlossaryProjectController(useCase, requireAuthMiddleware)

	requestBody, _ := json.Marshal(dto.GlossaryProjectRequestDTO{ProjectLabel: "Abastecar"})
	request := authenticatedRequest(http.MethodPut, "/api/v1/glossary/projects/someone-elses-id", requestBody)
	request.SetPathValue("id", "someone-elses-id")
	responseRecorder := httptest.NewRecorder()

	requireAuthMiddleware.Wrap(updateGlossaryProjectController.Update)(responseRecorder, request)

	if responseRecorder.Code != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, responseRecorder.Code)
	}
}
