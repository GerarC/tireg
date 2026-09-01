package controller_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	timeEntryException "github.com/gerarc/tireg/internal/time-registry/domain/exception"
	"github.com/gerarc/tireg/internal/time-registry/domain/model"
	"github.com/gerarc/tireg/internal/time-registry/infrastructure/in/rest/v1/controller"
	"github.com/gerarc/tireg/internal/time-registry/infrastructure/in/rest/v1/dto"
)

type stubFindTimeEntryByIDUseCase struct {
	findByIDFunc func(string, string) (model.TimeEntry, error)
}

func (stub *stubFindTimeEntryByIDUseCase) FindByID(ownerID string, id string) (model.TimeEntry, error) {
	return stub.findByIDFunc(ownerID, id)
}

func TestFindByID_Returns200_WhenOwnedByAuthenticatedUser(t *testing.T) {
	requireAuthMiddleware := authenticatedMiddleware("owner-id")
	useCase := &stubFindTimeEntryByIDUseCase{
		findByIDFunc: func(ownerID string, id string) (model.TimeEntry, error) {
			return model.TimeEntry{ID: id, OwnerID: ownerID}, nil
		},
	}
	findTimeEntryByIDController := controller.NewFindTimeEntryByIDController(useCase, requireAuthMiddleware)

	request := authenticatedRequest(http.MethodGet, "/api/v1/time-entries/entry-id", nil)
	request.SetPathValue("id", "entry-id")
	responseRecorder := httptest.NewRecorder()

	requireAuthMiddleware.Wrap(findTimeEntryByIDController.FindByID)(responseRecorder, request)

	if responseRecorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, responseRecorder.Code)
	}

	var responseDTO dto.TimeEntryResponseDTO
	if err := json.NewDecoder(responseRecorder.Body).Decode(&responseDTO); err != nil {
		t.Fatalf("unexpected error decoding response: %v", err)
	}

	if responseDTO.ID != "entry-id" {
		t.Fatalf("unexpected response: %+v", responseDTO)
	}
}

func TestFindByID_Returns404_WhenNotOwnedByAuthenticatedUser(t *testing.T) {
	requireAuthMiddleware := authenticatedMiddleware("owner-id")
	useCase := &stubFindTimeEntryByIDUseCase{
		findByIDFunc: func(ownerID string, id string) (model.TimeEntry, error) {
			return model.TimeEntry{}, timeEntryException.NewTimeEntryNotFoundError()
		},
	}
	findTimeEntryByIDController := controller.NewFindTimeEntryByIDController(useCase, requireAuthMiddleware)

	request := authenticatedRequest(http.MethodGet, "/api/v1/time-entries/someone-elses-id", nil)
	request.SetPathValue("id", "someone-elses-id")
	responseRecorder := httptest.NewRecorder()

	requireAuthMiddleware.Wrap(findTimeEntryByIDController.FindByID)(responseRecorder, request)

	if responseRecorder.Code != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, responseRecorder.Code)
	}
}
