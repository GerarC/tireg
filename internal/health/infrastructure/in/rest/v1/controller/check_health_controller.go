package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gerarc/tireg/internal/health/domain/api"
	"github.com/gerarc/tireg/internal/health/infrastructure/in/rest/v1/dto"
	"github.com/gerarc/tireg/internal/health/infrastructure/in/rest/v1/mapper"
	"github.com/gerarc/tireg/internal/health/infrastructure/in/rest/v1/util/constant"
)

type CheckHealthController struct {
	checkHealthUseCase api.CheckHealthUseCase
}

func NewCheckHealthController(checkHealthUseCase api.CheckHealthUseCase) *CheckHealthController {
	return &CheckHealthController{
		checkHealthUseCase: checkHealthUseCase,
	}
}

func (checkHealthController *CheckHealthController) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc(constant.HealthRoutePath, checkHealthController.CheckHealth)
}

// CheckHealth godoc
// @Summary Check service health
// @Description Returns the current health status of the service
// @Tags health
// @Produce json
// @Success 200 {object} dto.HealthResponseDTO
// @Router /api/v1/health [get]
func (checkHealthController *CheckHealthController) CheckHealth(responseWriter http.ResponseWriter, request *http.Request) {
	healthStatus := checkHealthController.checkHealthUseCase.CheckHealth()

	var healthResponseDTO dto.HealthResponseDTO = mapper.ToHealthResponseDTO(healthStatus)

	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(healthResponseDTO)
}
