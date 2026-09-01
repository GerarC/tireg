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

type stubUpdateTimeEntryUseCase struct {
	updateFunc func(string, string, model.TimeEntry) (model.TimeEntry, error)
}

func (stub *stubUpdateTimeEntryUseCase) Update(ownerID string, id string, timeEntry model.TimeEntry) (model.TimeEntry, error) {
	return stub.updateFunc(ownerID, id, timeEntry)
}

func TestUpdate_Returns200_WhenValid(t *testing.T) {
	requireAuthMiddleware := authenticatedMiddleware("owner-id")
	useCase := &stubUpdateTimeEntryUseCase{
		updateFunc: func(ownerID string, id string, timeEntry model.TimeEntry) (model.TimeEntry, error) {
			timeEntry.ID = id
			timeEntry.OwnerID = ownerID
			return timeEntry, nil
		},
	}
	updateTimeEntryController := controller.NewUpdateTimeEntryController(useCase, requireAuthMiddleware)

	requestBody, _ := json.Marshal(dto.TimeEntryRequestDTO{Date: "2026-08-03", ProjectLabel: "Abastecar", Start: "08:00", End: "10:30", Hours: 2.5})
	request := authenticatedRequest(http.MethodPut, "/api/v1/time-entries/entry-id", requestBody)
	request.SetPathValue("id", "entry-id")
	responseRecorder := httptest.NewRecorder()

	requireAuthMiddleware.Wrap(updateTimeEntryController.Update)(responseRecorder, request)

	if responseRecorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, responseRecorder.Code)
	}
}

func TestUpdate_Returns404_WhenNotOwnedByAuthenticatedUser(t *testing.T) {
	requireAuthMiddleware := authenticatedMiddleware("owner-id")
	useCase := &stubUpdateTimeEntryUseCase{
		updateFunc: func(ownerID string, id string, timeEntry model.TimeEntry) (model.TimeEntry, error) {
			return model.TimeEntry{}, timeEntryException.NewTimeEntryNotFoundError()
		},
	}
	updateTimeEntryController := controller.NewUpdateTimeEntryController(useCase, requireAuthMiddleware)

	requestBody, _ := json.Marshal(dto.TimeEntryRequestDTO{Date: "2026-08-03", ProjectLabel: "Abastecar", Start: "08:00", End: "10:30", Hours: 2.5})
	request := authenticatedRequest(http.MethodPut, "/api/v1/time-entries/someone-elses-id", requestBody)
	request.SetPathValue("id", "someone-elses-id")
	responseRecorder := httptest.NewRecorder()

	requireAuthMiddleware.Wrap(updateTimeEntryController.Update)(responseRecorder, request)

	if responseRecorder.Code != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, responseRecorder.Code)
	}
}
