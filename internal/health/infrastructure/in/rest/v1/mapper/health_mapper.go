package mapper

import (
	"github.com/gerarc/tireg/internal/health/domain/model"
	"github.com/gerarc/tireg/internal/health/infrastructure/in/rest/v1/dto"
)

func ToHealthResponseDTO(healthStatus model.HealthStatus) dto.HealthResponseDTO {
	return dto.HealthResponseDTO{
		Status:  healthStatus.Status,
		Service: healthStatus.Service,
	}
}
