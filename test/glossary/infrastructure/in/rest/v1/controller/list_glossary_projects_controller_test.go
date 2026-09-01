package controller_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gerarc/tireg/internal/glossary/domain/model"
	"github.com/gerarc/tireg/internal/glossary/infrastructure/in/rest/v1/controller"
	"github.com/gerarc/tireg/internal/glossary/infrastructure/in/rest/v1/dto"
)

type stubListGlossaryProjectsControllerUseCase struct {
	listFunc func(string) ([]model.GlossaryProject, error)
}

func (stub *stubListGlossaryProjectsControllerUseCase) List(ownerID string) ([]model.GlossaryProject, error) {
	return stub.listFunc(ownerID)
}

func TestListProjects_Returns200_WithOwnersProjects(t *testing.T) {
	requireAuthMiddleware := authenticatedMiddleware("owner-id")
	useCase := &stubListGlossaryProjectsControllerUseCase{
		listFunc: func(ownerID string) ([]model.GlossaryProject, error) {
			return []model.GlossaryProject{{ID: "project-id", OwnerID: ownerID, ProjectLabel: "Abastecar"}}, nil
		},
	}
	listGlossaryProjectsController := controller.NewListGlossaryProjectsController(useCase, requireAuthMiddleware)

	responseRecorder := httptest.NewRecorder()
	requireAuthMiddleware.Wrap(listGlossaryProjectsController.List)(responseRecorder, authenticatedRequest(http.MethodGet, "/api/v1/glossary/projects", nil))

	if responseRecorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, responseRecorder.Code)
	}

	var responseDTOs []dto.GlossaryProjectResponseDTO
	if err := json.NewDecoder(responseRecorder.Body).Decode(&responseDTOs); err != nil {
		t.Fatalf("unexpected error decoding response: %v", err)
	}

	if len(responseDTOs) != 1 || responseDTOs[0].ID != "project-id" {
		t.Fatalf("unexpected response: %+v", responseDTOs)
	}
}
