package usecase

import (
	"github.com/gerarc/tireg/internal/health/domain/api"
	"github.com/gerarc/tireg/internal/health/domain/model"
	"github.com/gerarc/tireg/internal/health/domain/util/constant"
)

type HealthUseCaseImplemented struct{}

func NewHealthUseCase() api.HealthUseCase {
	return &HealthUseCaseImplemented{}
}

func (healthUseCase *HealthUseCaseImplemented) CheckHealth() model.HealthStatus {
	return model.HealthStatus{
		Status:  constant.StatusUp,
		Service: constant.ServiceName,
	}
}
