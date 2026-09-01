package controller_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	authException "github.com/gerarc/tireg/internal/auth/domain/exception"
	"github.com/gerarc/tireg/internal/auth/domain/model"
	"github.com/gerarc/tireg/internal/auth/infrastructure/in/rest/v1/controller"
	"github.com/gerarc/tireg/internal/auth/infrastructure/in/rest/v1/dto"
)

type stubAuthUseCase struct {
	loginFunc func(model.Credentials) (model.AccessToken, error)
}

func (stub *stubAuthUseCase) Login(credentials model.Credentials) (model.AccessToken, error) {
	return stub.loginFunc(credentials)
}

func doLoginRequest(t *testing.T, useCase *stubAuthUseCase, body []byte) *httptest.ResponseRecorder {
	t.Helper()

	request := httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewReader(body))
	responseRecorder := httptest.NewRecorder()

	authController := controller.NewLoginController(useCase)
	authController.Login(responseRecorder, request)

	return responseRecorder
}

func TestLogin_Returns200_WhenValid(t *testing.T) {
	useCase := &stubAuthUseCase{
		loginFunc: func(credentials model.Credentials) (model.AccessToken, error) {
			return model.AccessToken{Token: "jwt-token", ExpiresAt: time.Now()}, nil
		},
	}

	requestBody, _ := json.Marshal(dto.LoginRequestDTO{Identifier: "ada", Password: "secret"})
	responseRecorder := doLoginRequest(t, useCase, requestBody)

	if responseRecorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, responseRecorder.Code)
	}

	var responseDTO dto.LoginResponseDTO
	if err := json.NewDecoder(responseRecorder.Body).Decode(&responseDTO); err != nil {
		t.Fatalf("unexpected error decoding response: %v", err)
	}

	if responseDTO.AccessToken != "jwt-token" {
		t.Fatalf("expected access token %q, got %q", "jwt-token", responseDTO.AccessToken)
	}
}

func TestLogin_Returns401_WhenInvalidCredentials(t *testing.T) {
	useCase := &stubAuthUseCase{
		loginFunc: func(credentials model.Credentials) (model.AccessToken, error) {
			return model.AccessToken{}, authException.NewInvalidCredentialsError()
		},
	}

	requestBody, _ := json.Marshal(dto.LoginRequestDTO{Identifier: "ada", Password: "wrong"})
	responseRecorder := doLoginRequest(t, useCase, requestBody)

	if responseRecorder.Code != http.StatusUnauthorized {
		t.Fatalf("expected status %d, got %d", http.StatusUnauthorized, responseRecorder.Code)
	}
}

func TestLogin_Returns400_WhenBodyIsInvalidJSON(t *testing.T) {
	useCase := &stubAuthUseCase{}

	responseRecorder := doLoginRequest(t, useCase, []byte("{invalid"))

	if responseRecorder.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, responseRecorder.Code)
	}
}
