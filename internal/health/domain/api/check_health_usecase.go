package api

import "github.com/gerarc/tireg/internal/health/domain/model"

// CheckHealthUseCase exposes the operation to check the health of the service.
type CheckHealthUseCase interface {
	// CheckHealth returns the current health status of the service.
	CheckHealth() model.HealthStatus
}
