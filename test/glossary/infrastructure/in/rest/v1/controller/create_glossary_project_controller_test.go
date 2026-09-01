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

type stubCreateGlossaryProjectUseCase struct {
	createFunc func(string, model.GlossaryProject) (model.GlossaryProject, error)
}

func (stub *stubCreateGlossaryProjectUseCase) Create(ownerID string, glossaryProject model.GlossaryProject) (model.GlossaryProject, error) {
	return stub.createFunc(ownerID, glossaryProject)
}

func TestCreateProject_Returns201_WhenValid(t *testing.T) {
	requireAuthMiddleware := authenticatedMiddleware("owner-id")
	useCase := &stubCreateGlossaryProjectUseCase{
		createFunc: func(ownerID string, glossaryProject model.GlossaryProject) (model.GlossaryProject, error) {
			glossaryProject.ID = "generated-id"
			glossaryProject.OwnerID = ownerID
			return glossaryProject, nil
		},
	}
	createGlossaryProjectController := controller.NewCreateGlossaryProjectController(useCase, requireAuthMiddleware)

	requestBody, _ := json.Marshal(dto.GlossaryProjectRequestDTO{ProjectLabel: "Abastecar"})
	responseRecorder := httptest.NewRecorder()
	requireAuthMiddleware.Wrap(createGlossaryProjectController.Create)(responseRecorder, authenticatedRequest(http.MethodPost, "/api/v1/glossary/projects", requestBody))

	if responseRecorder.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d", http.StatusCreated, responseRecorder.Code)
	}

	var responseDTO dto.GlossaryProjectResponseDTO
	if err := json.NewDecoder(responseRecorder.Body).Decode(&responseDTO); err != nil {
		t.Fatalf("unexpected error decoding response: %v", err)
	}

	if responseDTO.ID != "generated-id" {
		t.Fatalf("unexpected response: %+v", responseDTO)
	}
}

func TestCreateProject_Returns400_WhenBodyIsInvalidJSON(t *testing.T) {
	requireAuthMiddleware := authenticatedMiddleware("owner-id")
	createGlossaryProjectController := controller.NewCreateGlossaryProjectController(&stubCreateGlossaryProjectUseCase{}, requireAuthMiddleware)

	responseRecorder := httptest.NewRecorder()
	requireAuthMiddleware.Wrap(createGlossaryProjectController.Create)(responseRecorder, authenticatedRequest(http.MethodPost, "/api/v1/glossary/projects", []byte("{invalid")))

	if responseRecorder.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, responseRecorder.Code)
	}
}
