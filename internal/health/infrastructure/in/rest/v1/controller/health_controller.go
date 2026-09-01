package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gerarc/tireg/internal/health/domain/api"
	"github.com/gerarc/tireg/internal/health/infrastructure/in/rest/v1/mapper"
	"github.com/gerarc/tireg/internal/health/infrastructure/in/rest/v1/util/constant"
)

type HealthController struct {
	healthUseCase api.HealthUseCase
}

func NewHealthController(healthUseCase api.HealthUseCase) *HealthController {
	return &HealthController{
		healthUseCase: healthUseCase,
	}
}

func (healthController *HealthController) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc(constant.HealthRoutePath, healthController.HealthCheck)
}

// HealthCheck godoc
// @Summary Check service health
// @Description Returns the current health status of the service
// @Tags health
// @Produce json
// @Success 200 {object} dto.HealthResponseDTO
// @Router /api/v1/health [get]
func (healthController *HealthController) HealthCheck(responseWriter http.ResponseWriter, request *http.Request) {
	healthStatus := healthController.healthUseCase.CheckHealth()
	healthResponseDTO := mapper.ToHealthResponseDTO(healthStatus)

	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(healthResponseDTO)
}
