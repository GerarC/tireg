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

type stubUpdateGlossaryTypeUseCase struct {
	updateFunc func(string, string, model.GlossaryType) (model.GlossaryType, error)
}

func (stub *stubUpdateGlossaryTypeUseCase) Update(ownerID string, id string, glossaryType model.GlossaryType) (model.GlossaryType, error) {
	return stub.updateFunc(ownerID, id, glossaryType)
}

func TestUpdate_Returns200_WhenValid(t *testing.T) {
	requireAuthMiddleware := authenticatedMiddleware("owner-id")
	useCase := &stubUpdateGlossaryTypeUseCase{
		updateFunc: func(ownerID string, id string, glossaryType model.GlossaryType) (model.GlossaryType, error) {
			glossaryType.ID = id
			glossaryType.OwnerID = ownerID
			return glossaryType, nil
		},
	}
	updateGlossaryTypeController := controller.NewUpdateGlossaryTypeController(useCase, requireAuthMiddleware)

	requestBody, _ := json.Marshal(dto.GlossaryTypeRequestDTO{TypeKey: "dev", Label: "Desarrollo"})
	request := authenticatedRequest(http.MethodPut, "/api/v1/glossary/types/type-id", requestBody)
	request.SetPathValue("id", "type-id")
	responseRecorder := httptest.NewRecorder()

	requireAuthMiddleware.Wrap(updateGlossaryTypeController.Update)(responseRecorder, request)

	if responseRecorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, responseRecorder.Code)
	}

	var responseDTO dto.GlossaryTypeResponseDTO
	if err := json.NewDecoder(responseRecorder.Body).Decode(&responseDTO); err != nil {
		t.Fatalf("unexpected error decoding response: %v", err)
	}

	if responseDTO.ID != "type-id" {
		t.Fatalf("unexpected response: %+v", responseDTO)
	}
}

func TestUpdate_Returns404_WhenNotOwnedByAuthenticatedUser(t *testing.T) {
	requireAuthMiddleware := authenticatedMiddleware("owner-id")
	useCase := &stubUpdateGlossaryTypeUseCase{
		updateFunc: func(ownerID string, id string, glossaryType model.GlossaryType) (model.GlossaryType, error) {
			return model.GlossaryType{}, glossaryException.NewGlossaryTypeNotFoundError()
		},
	}
	updateGlossaryTypeController := controller.NewUpdateGlossaryTypeController(useCase, requireAuthMiddleware)

	requestBody, _ := json.Marshal(dto.GlossaryTypeRequestDTO{TypeKey: "dev", Label: "Desarrollo"})
	request := authenticatedRequest(http.MethodPut, "/api/v1/glossary/types/someone-elses-id", requestBody)
	request.SetPathValue("id", "someone-elses-id")
	responseRecorder := httptest.NewRecorder()

	requireAuthMiddleware.Wrap(updateGlossaryTypeController.Update)(responseRecorder, request)

	if responseRecorder.Code != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, responseRecorder.Code)
	}
}
