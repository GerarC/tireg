package api

import "github.com/gerarc/tireg/internal/health/domain/model"

// HealthUseCase exposes the operations available to check the health of the service.
type HealthUseCase interface {
	// CheckHealth returns the current health status of the service.
	CheckHealth() model.HealthStatus
}
