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

	"github.com/gerarc/tireg/internal/task/domain/model"
	"github.com/gerarc/tireg/internal/task/infrastructure/in/rest/v1/controller"
	"github.com/gerarc/tireg/internal/task/infrastructure/in/rest/v1/dto"
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

type stubCreateTaskMappingUseCase struct {
	createFunc func(string, model.TaskMapping) (model.TaskMapping, error)
}

func (stub *stubCreateTaskMappingUseCase) Create(ownerID string, taskMapping model.TaskMapping) (model.TaskMapping, error) {
	return stub.createFunc(ownerID, taskMapping)
}

func TestCreate_Returns201_WhenValid(t *testing.T) {
	requireAuthMiddleware := authenticatedMiddleware("owner-id")
	useCase := &stubCreateTaskMappingUseCase{
		createFunc: func(ownerID string, taskMapping model.TaskMapping) (model.TaskMapping, error) {
			taskMapping.ID = "generated-id"
			taskMapping.OwnerID = ownerID
			return taskMapping, nil
		},
	}
	createTaskMappingController := controller.NewCreateTaskMappingController(useCase, requireAuthMiddleware)

	requestBody, _ := json.Marshal(dto.TaskMappingRequestDTO{ProjectLabel: "Abastecar", Pattern: "reunión con el cliente"})
	responseRecorder := httptest.NewRecorder()
	requireAuthMiddleware.Wrap(createTaskMappingController.Create)(responseRecorder, authenticatedRequest(http.MethodPost, "/api/v1/task-mappings", requestBody))

	if responseRecorder.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d", http.StatusCreated, responseRecorder.Code)
	}

	var responseDTO dto.TaskMappingResponseDTO
	if err := json.NewDecoder(responseRecorder.Body).Decode(&responseDTO); err != nil {
		t.Fatalf("unexpected error decoding response: %v", err)
	}

	if responseDTO.ID != "generated-id" {
		t.Fatalf("unexpected response: %+v", responseDTO)
	}
}

func TestCreate_Returns400_WhenBodyIsInvalidJSON(t *testing.T) {
	requireAuthMiddleware := authenticatedMiddleware("owner-id")
	createTaskMappingController := controller.NewCreateTaskMappingController(&stubCreateTaskMappingUseCase{}, requireAuthMiddleware)

	responseRecorder := httptest.NewRecorder()
	requireAuthMiddleware.Wrap(createTaskMappingController.Create)(responseRecorder, authenticatedRequest(http.MethodPost, "/api/v1/task-mappings", []byte("{invalid")))

	if responseRecorder.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, responseRecorder.Code)
	}
}

func TestCreate_Returns401_WhenTokenInvalid(t *testing.T) {
	requireAuthMiddleware := unauthenticatedMiddleware()
	createTaskMappingController := controller.NewCreateTaskMappingController(&stubCreateTaskMappingUseCase{}, requireAuthMiddleware)

	responseRecorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "/api/v1/task-mappings", nil)

	requireAuthMiddleware.Wrap(createTaskMappingController.Create)(responseRecorder, request)

	if responseRecorder.Code != http.StatusUnauthorized {
		t.Fatalf("expected status %d, got %d", http.StatusUnauthorized, responseRecorder.Code)
	}
}
