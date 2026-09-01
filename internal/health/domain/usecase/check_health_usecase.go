package usecase

import (
	"github.com/gerarc/tireg/internal/health/domain/api"
	"github.com/gerarc/tireg/internal/health/domain/model"
	"github.com/gerarc/tireg/internal/health/domain/util/constant"
)

type CheckHealthUseCaseImplemented struct{}

func NewCheckHealthUseCase() api.CheckHealthUseCase {
	return &CheckHealthUseCaseImplemented{}
}

func (checkHealthUseCase *CheckHealthUseCaseImplemented) CheckHealth() model.HealthStatus {
	return model.HealthStatus{
		Status:  constant.StatusUp,
		Service: constant.ServiceName,
	}
}
