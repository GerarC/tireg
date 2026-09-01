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

	"github.com/gerarc/tireg/internal/glossary/domain/model"
	"github.com/gerarc/tireg/internal/glossary/infrastructure/in/rest/v1/controller"
	"github.com/gerarc/tireg/internal/glossary/infrastructure/in/rest/v1/dto"
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

type stubGetGlossaryUseCase struct {
	getFunc func(string) (model.Glossary, error)
}

func (stub *stubGetGlossaryUseCase) Get(ownerID string) (model.Glossary, error) {
	return stub.getFunc(ownerID)
}

func TestGet_Returns200_WithOwnersGlossary(t *testing.T) {
	requireAuthMiddleware := authenticatedMiddleware("owner-id")
	useCase := &stubGetGlossaryUseCase{
		getFunc: func(ownerID string) (model.Glossary, error) {
			return model.Glossary{
				Types:    []model.GlossaryType{{ID: "type-id", OwnerID: ownerID}},
				Projects: []model.GlossaryProject{{ID: "project-id", OwnerID: ownerID}},
			}, nil
		},
	}
	getGlossaryController := controller.NewGetGlossaryController(useCase, requireAuthMiddleware)

	responseRecorder := httptest.NewRecorder()
	requireAuthMiddleware.Wrap(getGlossaryController.Get)(responseRecorder, authenticatedRequest(http.MethodGet, "/api/v1/glossary", nil))

	if responseRecorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, responseRecorder.Code)
	}

	var responseDTO dto.GlossaryResponseDTO
	if err := json.NewDecoder(responseRecorder.Body).Decode(&responseDTO); err != nil {
		t.Fatalf("unexpected error decoding response: %v", err)
	}

	if len(responseDTO.Types) != 1 || len(responseDTO.Projects) != 1 {
		t.Fatalf("unexpected response: %+v", responseDTO)
	}
}

func TestGet_Returns401_WhenTokenInvalid(t *testing.T) {
	requireAuthMiddleware := unauthenticatedMiddleware()
	getGlossaryController := controller.NewGetGlossaryController(&stubGetGlossaryUseCase{}, requireAuthMiddleware)

	responseRecorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/api/v1/glossary", nil)

	requireAuthMiddleware.Wrap(getGlossaryController.Get)(responseRecorder, request)

	if responseRecorder.Code != http.StatusUnauthorized {
		t.Fatalf("expected status %d, got %d", http.StatusUnauthorized, responseRecorder.Code)
	}
}
