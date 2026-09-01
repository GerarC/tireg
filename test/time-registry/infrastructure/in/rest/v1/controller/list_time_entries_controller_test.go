package controller_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gerarc/tireg/internal/time-registry/domain/model"
	"github.com/gerarc/tireg/internal/time-registry/infrastructure/in/rest/v1/controller"
	"github.com/gerarc/tireg/internal/time-registry/infrastructure/in/rest/v1/dto"
)

type stubListTimeEntriesUseCase struct {
	listFunc func(string) ([]model.TimeEntry, error)
}

func (stub *stubListTimeEntriesUseCase) List(ownerID string) ([]model.TimeEntry, error) {
	return stub.listFunc(ownerID)
}

func TestList_Returns200_WithOwnersEntries(t *testing.T) {
	requireAuthMiddleware := authenticatedMiddleware("owner-id")
	useCase := &stubListTimeEntriesUseCase{
		listFunc: func(ownerID string) ([]model.TimeEntry, error) {
			return []model.TimeEntry{{ID: "entry-id", OwnerID: ownerID}}, nil
		},
	}
	listTimeEntriesController := controller.NewListTimeEntriesController(useCase, requireAuthMiddleware)

	responseRecorder := httptest.NewRecorder()
	requireAuthMiddleware.Wrap(listTimeEntriesController.List)(responseRecorder, authenticatedRequest(http.MethodGet, "/api/v1/time-entries", nil))

	if responseRecorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, responseRecorder.Code)
	}

	var responseDTOs []dto.TimeEntryResponseDTO
	if err := json.NewDecoder(responseRecorder.Body).Decode(&responseDTOs); err != nil {
		t.Fatalf("unexpected error decoding response: %v", err)
	}

	if len(responseDTOs) != 1 || responseDTOs[0].ID != "entry-id" {
		t.Fatalf("unexpected response: %+v", responseDTOs)
	}
}

func TestList_Returns401_WhenTokenInvalid(t *testing.T) {
	requireAuthMiddleware := unauthenticatedMiddleware()
	listTimeEntriesController := controller.NewListTimeEntriesController(&stubListTimeEntriesUseCase{}, requireAuthMiddleware)

	responseRecorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/api/v1/time-entries", nil)

	requireAuthMiddleware.Wrap(listTimeEntriesController.List)(responseRecorder, request)

	if responseRecorder.Code != http.StatusUnauthorized {
		t.Fatalf("expected status %d, got %d", http.StatusUnauthorized, responseRecorder.Code)
	}
}
