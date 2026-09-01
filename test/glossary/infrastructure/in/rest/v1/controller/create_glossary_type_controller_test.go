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

type stubCreateGlossaryTypeUseCase struct {
	createFunc func(string, model.GlossaryType) (model.GlossaryType, error)
}

func (stub *stubCreateGlossaryTypeUseCase) Create(ownerID string, glossaryType model.GlossaryType) (model.GlossaryType, error) {
	return stub.createFunc(ownerID, glossaryType)
}

func TestCreate_Returns201_WhenValid(t *testing.T) {
	requireAuthMiddleware := authenticatedMiddleware("owner-id")
	useCase := &stubCreateGlossaryTypeUseCase{
		createFunc: func(ownerID string, glossaryType model.GlossaryType) (model.GlossaryType, error) {
			glossaryType.ID = "generated-id"
			glossaryType.OwnerID = ownerID
			return glossaryType, nil
		},
	}
	createGlossaryTypeController := controller.NewCreateGlossaryTypeController(useCase, requireAuthMiddleware)

	requestBody, _ := json.Marshal(dto.GlossaryTypeRequestDTO{TypeKey: "dev", Label: "Desarrollo"})
	responseRecorder := httptest.NewRecorder()
	requireAuthMiddleware.Wrap(createGlossaryTypeController.Create)(responseRecorder, authenticatedRequest(http.MethodPost, "/api/v1/glossary/types", requestBody))

	if responseRecorder.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d", http.StatusCreated, responseRecorder.Code)
	}

	var responseDTO dto.GlossaryTypeResponseDTO
	if err := json.NewDecoder(responseRecorder.Body).Decode(&responseDTO); err != nil {
		t.Fatalf("unexpected error decoding response: %v", err)
	}

	if responseDTO.ID != "generated-id" {
		t.Fatalf("unexpected response: %+v", responseDTO)
	}
}

func TestCreate_Returns400_WhenBodyIsInvalidJSON(t *testing.T) {
	requireAuthMiddleware := authenticatedMiddleware("owner-id")
	createGlossaryTypeController := controller.NewCreateGlossaryTypeController(&stubCreateGlossaryTypeUseCase{}, requireAuthMiddleware)

	responseRecorder := httptest.NewRecorder()
	requireAuthMiddleware.Wrap(createGlossaryTypeController.Create)(responseRecorder, authenticatedRequest(http.MethodPost, "/api/v1/glossary/types", []byte("{invalid")))

	if responseRecorder.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, responseRecorder.Code)
	}
}

func TestCreate_Returns401_WhenTokenInvalid(t *testing.T) {
	requireAuthMiddleware := unauthenticatedMiddleware()
	createGlossaryTypeController := controller.NewCreateGlossaryTypeController(&stubCreateGlossaryTypeUseCase{}, requireAuthMiddleware)

	responseRecorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "/api/v1/glossary/types", nil)

	requireAuthMiddleware.Wrap(createGlossaryTypeController.Create)(responseRecorder, request)

	if responseRecorder.Code != http.StatusUnauthorized {
		t.Fatalf("expected status %d, got %d", http.StatusUnauthorized, responseRecorder.Code)
	}
}
