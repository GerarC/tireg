package usecase_test

import (
	"testing"

	"github.com/gerarc/tireg/internal/health/domain/usecase"
	"github.com/gerarc/tireg/internal/health/domain/util/constant"
)

func TestNewHealthUseCase_CheckHealth_ReturnsUpStatus(t *testing.T) {
	healthUseCase := usecase.NewHealthUseCase()

	healthStatus := healthUseCase.CheckHealth()

	if healthStatus.Status != constant.StatusUp {
		t.Fatalf("expected status %q, got %q", constant.StatusUp, healthStatus.Status)
	}

	if healthStatus.Service != constant.ServiceName {
		t.Fatalf("expected service %q, got %q", constant.ServiceName, healthStatus.Service)
	}
}
