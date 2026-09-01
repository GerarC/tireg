package controller_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gerarc/tireg/internal/health/domain/model"
	"github.com/gerarc/tireg/internal/health/infrastructure/in/rest/v1/controller"
	"github.com/gerarc/tireg/internal/health/infrastructure/in/rest/v1/dto"
	"github.com/gerarc/tireg/internal/health/infrastructure/in/rest/v1/util/constant"
)

type stubHealthUseCase struct {
	healthStatus model.HealthStatus
}

func (stub *stubHealthUseCase) CheckHealth() model.HealthStatus {
	return stub.healthStatus
}

func TestCheckHealthController_CheckHealth_ReturnsMappedStatus(t *testing.T) {
	stub := &stubHealthUseCase{healthStatus: model.HealthStatus{Status: "UP", Service: "tireg"}}
	healthController := controller.NewCheckHealthController(stub)

	request := httptest.NewRequest(http.MethodGet, constant.HealthRoutePath, nil)
	responseRecorder := httptest.NewRecorder()

	healthController.CheckHealth(responseRecorder, request)

	if responseRecorder.Code != http.StatusOK {
		t.Fatalf("expected status code %d, got %d", http.StatusOK, responseRecorder.Code)
	}

	var healthResponseDTO dto.HealthResponseDTO
	if err := json.NewDecoder(responseRecorder.Body).Decode(&healthResponseDTO); err != nil {
		t.Fatalf("unexpected error decoding response body: %v", err)
	}

	if healthResponseDTO.Status != "UP" || healthResponseDTO.Service != "tireg" {
		t.Fatalf("unexpected response body: %+v", healthResponseDTO)
	}
}

func TestCheckHealthController_RegisterRoutes_RegistersHealthRoute(t *testing.T) {
	stub := &stubHealthUseCase{healthStatus: model.HealthStatus{Status: "UP", Service: "tireg"}}
	healthController := controller.NewCheckHealthController(stub)

	mux := http.NewServeMux()
	healthController.RegisterRoutes(mux)

	request := httptest.NewRequest(http.MethodGet, constant.HealthRoutePath, nil)
	responseRecorder := httptest.NewRecorder()

	mux.ServeHTTP(responseRecorder, request)

	if responseRecorder.Code != http.StatusOK {
		t.Fatalf("expected status code %d, got %d", http.StatusOK, responseRecorder.Code)
	}
}
