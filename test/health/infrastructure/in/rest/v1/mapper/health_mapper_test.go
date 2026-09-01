package mapper_test

import (
	"testing"

	"github.com/gerarc/tireg/internal/health/domain/model"
	"github.com/gerarc/tireg/internal/health/infrastructure/in/rest/v1/mapper"
)

func TestToHealthResponseDTO_MapsAllFields(t *testing.T) {
	healthStatus := model.HealthStatus{Status: "UP", Service: "tireg"}

	healthResponseDTO := mapper.ToHealthResponseDTO(healthStatus)

	if healthResponseDTO.Status != healthStatus.Status {
		t.Fatalf("expected status %q, got %q", healthStatus.Status, healthResponseDTO.Status)
	}

	if healthResponseDTO.Service != healthStatus.Service {
		t.Fatalf("expected service %q, got %q", healthStatus.Service, healthResponseDTO.Service)
	}
}
