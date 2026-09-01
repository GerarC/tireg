package controller_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	authException "github.com/gerarc/tireg/internal/auth/domain/exception"
	authModel "github.com/gerarc/tireg/internal/auth/domain/model"
	sharedMiddleware "github.com/gerarc/tireg/internal/shared/infrastructure/in/rest/middleware"

	"github.com/gerarc/tireg/internal/time-registry/domain/model"
	"github.com/gerarc/tireg/internal/time-registry/infrastructure/in/rest/v1/controller"
	"github.com/gerarc/tireg/internal/time-registry/infrastructure/in/rest/v1/dto"
)

type stubVerifyTokenUseCase struct {
	verifyFunc func(string) (authModel.AuthenticatedUser, error)
}

func (stub *stubVerifyTokenUseCase) Verify(token string) (authModel.AuthenticatedUser, error) {
	return stub.verifyFunc(token)
}

func authenticatedMiddleware(ownerID string) *sharedMiddleware.RequireAuthMiddleware {
	return sharedMiddleware.NewRequireAuthMiddleware(&stubVerifyTokenUseCase{
		verifyFunc: func(token string) (authModel.AuthenticatedUser, error) {
			return authModel.AuthenticatedUser{ID: ownerID, Username: "ada"}, nil
		},
	})
}

func unauthenticatedMiddleware() *sharedMiddleware.RequireAuthMiddleware {
	return sharedMiddleware.NewRequireAuthMiddleware(&stubVerifyTokenUseCase{
		verifyFunc: func(token string) (authModel.AuthenticatedUser, error) {
			return authModel.AuthenticatedUser{}, authException.NewInvalidCredentialsError()
		},
	})
}

func authenticatedRequest(method string, path string, body []byte) *http.Request {
	request := httptest.NewRequest(method, path, bytes.NewReader(body))
	request.Header.Set("Authorization", "Bearer valid-token")
	return request
}

type stubCreateTimeEntryUseCase struct {
	createFunc func(string, model.TimeEntry) (model.TimeEntry, error)
}

func (stub *stubCreateTimeEntryUseCase) Create(ownerID string, timeEntry model.TimeEntry) (model.TimeEntry, error) {
	return stub.createFunc(ownerID, timeEntry)
}

func TestCreate_Returns201_WhenValid(t *testing.T) {
	requireAuthMiddleware := authenticatedMiddleware("owner-id")
	useCase := &stubCreateTimeEntryUseCase{
		createFunc: func(ownerID string, timeEntry model.TimeEntry) (model.TimeEntry, error) {
			timeEntry.ID = "generated-id"
			timeEntry.OwnerID = ownerID
			return timeEntry, nil
		},
	}
	createTimeEntryController := controller.NewCreateTimeEntryController(useCase, requireAuthMiddleware)

	requestBody, _ := json.Marshal(dto.TimeEntryRequestDTO{Date: "2026-08-03", ProjectLabel: "Abastecar", Start: "08:00", End: "10:30", Hours: 2.5})
	responseRecorder := httptest.NewRecorder()
	requireAuthMiddleware.Wrap(createTimeEntryController.Create)(responseRecorder, authenticatedRequest(http.MethodPost, "/api/v1/time-entries", requestBody))

	if responseRecorder.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d", http.StatusCreated, responseRecorder.Code)
	}

	var responseDTO dto.TimeEntryResponseDTO
	if err := json.NewDecoder(responseRecorder.Body).Decode(&responseDTO); err != nil {
		t.Fatalf("unexpected error decoding response: %v", err)
	}

	if responseDTO.ID != "generated-id" {
		t.Fatalf("unexpected response: %+v", responseDTO)
	}
}

func TestCreate_Returns400_WhenBodyIsInvalidJSON(t *testing.T) {
	requireAuthMiddleware := authenticatedMiddleware("owner-id")
	createTimeEntryController := controller.NewCreateTimeEntryController(&stubCreateTimeEntryUseCase{}, requireAuthMiddleware)

	responseRecorder := httptest.NewRecorder()
	requireAuthMiddleware.Wrap(createTimeEntryController.Create)(responseRecorder, authenticatedRequest(http.MethodPost, "/api/v1/time-entries", []byte("{invalid")))

	if responseRecorder.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, responseRecorder.Code)
	}
}

func TestCreate_Returns401_WhenTokenInvalid(t *testing.T) {
	requireAuthMiddleware := unauthenticatedMiddleware()
	createTimeEntryController := controller.NewCreateTimeEntryController(&stubCreateTimeEntryUseCase{}, requireAuthMiddleware)

	responseRecorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "/api/v1/time-entries", nil)

	requireAuthMiddleware.Wrap(createTimeEntryController.Create)(responseRecorder, request)

	if responseRecorder.Code != http.StatusUnauthorized {
		t.Fatalf("expected status %d, got %d", http.StatusUnauthorized, responseRecorder.Code)
	}
}
