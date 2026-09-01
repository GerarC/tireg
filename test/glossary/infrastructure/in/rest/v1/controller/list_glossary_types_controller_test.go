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

type stubListGlossaryTypesControllerUseCase struct {
	listFunc func(string) ([]model.GlossaryType, error)
}

func (stub *stubListGlossaryTypesControllerUseCase) List(ownerID string) ([]model.GlossaryType, error) {
	return stub.listFunc(ownerID)
}

func TestList_Returns200_WithOwnersTypes(t *testing.T) {
	requireAuthMiddleware := authenticatedMiddleware("owner-id")
	useCase := &stubListGlossaryTypesControllerUseCase{
		listFunc: func(ownerID string) ([]model.GlossaryType, error) {
			return []model.GlossaryType{{ID: "type-id", OwnerID: ownerID, TypeKey: "dev"}}, nil
		},
	}
	listGlossaryTypesController := controller.NewListGlossaryTypesController(useCase, requireAuthMiddleware)

	responseRecorder := httptest.NewRecorder()
	requireAuthMiddleware.Wrap(listGlossaryTypesController.List)(responseRecorder, authenticatedRequest(http.MethodGet, "/api/v1/glossary/types", nil))

	if responseRecorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, responseRecorder.Code)
	}

	var responseDTOs []dto.GlossaryTypeResponseDTO
	if err := json.NewDecoder(responseRecorder.Body).Decode(&responseDTOs); err != nil {
		t.Fatalf("unexpected error decoding response: %v", err)
	}

	if len(responseDTOs) != 1 || responseDTOs[0].ID != "type-id" {
		t.Fatalf("unexpected response: %+v", responseDTOs)
	}
}

func TestList_Returns401_WhenTokenInvalid(t *testing.T) {
	requireAuthMiddleware := unauthenticatedMiddleware()
	listGlossaryTypesController := controller.NewListGlossaryTypesController(&stubListGlossaryTypesControllerUseCase{}, requireAuthMiddleware)

	responseRecorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/api/v1/glossary/types", nil)

	requireAuthMiddleware.Wrap(listGlossaryTypesController.List)(responseRecorder, request)

	if responseRecorder.Code != http.StatusUnauthorized {
		t.Fatalf("expected status %d, got %d", http.StatusUnauthorized, responseRecorder.Code)
	}
}
