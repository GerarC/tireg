package controller_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	sharedRest "github.com/gerarc/tireg/internal/shared/infrastructure/in/rest"

	userException "github.com/gerarc/tireg/internal/user/domain/exception"
	"github.com/gerarc/tireg/internal/user/domain/model"
	"github.com/gerarc/tireg/internal/user/infrastructure/in/rest/v1/controller"
	"github.com/gerarc/tireg/internal/user/infrastructure/in/rest/v1/dto"
)

type stubRegisterUserUseCase struct {
	registerFunc func(model.UserRegistration) (model.User, error)
}

func (stub *stubRegisterUserUseCase) Register(registration model.UserRegistration) (model.User, error) {
	return stub.registerFunc(registration)
}

func doRegisterRequest(t *testing.T, useCase *stubRegisterUserUseCase, payload any) *httptest.ResponseRecorder {
	t.Helper()

	body, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("unexpected error marshaling payload: %v", err)
	}

	request := httptest.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewReader(body))
	responseRecorder := httptest.NewRecorder()

	userController := controller.NewRegisterUserController(useCase)
	userController.Register(responseRecorder, request)

	return responseRecorder
}

func TestRegister_Returns201_WhenValid(t *testing.T) {
	useCase := &stubRegisterUserUseCase{
		registerFunc: func(registration model.UserRegistration) (model.User, error) {
			return model.User{ID: "user-id", Username: registration.Username}, nil
		},
	}

	responseRecorder := doRegisterRequest(t, useCase, dto.UserRegisterRequestDTO{
		FirstName: "Ada",
		LastName:  "Lovelace",
		Username:  "ada",
		Email:     "ada@example.com",
		Password:  "super-secret",
	})

	if responseRecorder.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d", http.StatusCreated, responseRecorder.Code)
	}

	var responseDTO dto.UserResponseDTO
	if err := json.NewDecoder(responseRecorder.Body).Decode(&responseDTO); err != nil {
		t.Fatalf("unexpected error decoding response: %v", err)
	}

	if responseDTO.Username != "ada" {
		t.Fatalf("expected username %q, got %q", "ada", responseDTO.Username)
	}
}

func TestRegister_Returns409_WhenUsernameAndEmailAlreadyTaken(t *testing.T) {
	useCase := &stubRegisterUserUseCase{
		registerFunc: func(registration model.UserRegistration) (model.User, error) {
			return model.User{}, userException.NewAlreadyTakenError("username already taken", "email already registered")
		},
	}

	responseRecorder := doRegisterRequest(t, useCase, dto.UserRegisterRequestDTO{
		FirstName: "Ada",
		LastName:  "Lovelace",
		Username:  "ada",
		Email:     "ada@example.com",
		Password:  "super-secret",
	})

	if responseRecorder.Code != http.StatusConflict {
		t.Fatalf("expected status %d, got %d", http.StatusConflict, responseRecorder.Code)
	}

	var errorResponse sharedRest.ErrorResponseDTO
	if err := json.NewDecoder(responseRecorder.Body).Decode(&errorResponse); err != nil {
		t.Fatalf("unexpected error decoding error response: %v", err)
	}

	if len(errorResponse.Details) != 2 {
		t.Fatalf("expected 2 details, got %d: %v", len(errorResponse.Details), errorResponse.Details)
	}
}

func TestRegister_Returns400_WhenBodyIsInvalidJSON(t *testing.T) {
	useCase := &stubRegisterUserUseCase{}

	request := httptest.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewReader([]byte("{invalid")))
	responseRecorder := httptest.NewRecorder()

	userController := controller.NewRegisterUserController(useCase)
	userController.Register(responseRecorder, request)

	if responseRecorder.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, responseRecorder.Code)
	}
}
