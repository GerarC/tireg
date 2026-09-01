package controller

import (
	"encoding/json"
	"net/http"

	sharedRest "github.com/gerarc/tireg/internal/shared/infrastructure/in/rest"
	sharedMiddleware "github.com/gerarc/tireg/internal/shared/infrastructure/in/rest/middleware"

	"github.com/gerarc/tireg/internal/time-registry/domain/api"
	"github.com/gerarc/tireg/internal/time-registry/infrastructure/in/rest/v1/dto"
	"github.com/gerarc/tireg/internal/time-registry/infrastructure/in/rest/v1/mapper"
	"github.com/gerarc/tireg/internal/time-registry/infrastructure/in/rest/v1/util/constant"
)

type FindTimeEntryByIDController struct {
	findTimeEntryByIDUseCase api.FindTimeEntryByIDUseCase
	requireAuthMiddleware    *sharedMiddleware.RequireAuthMiddleware
}

func NewFindTimeEntryByIDController(findTimeEntryByIDUseCase api.FindTimeEntryByIDUseCase, requireAuthMiddleware *sharedMiddleware.RequireAuthMiddleware) *FindTimeEntryByIDController {
	return &FindTimeEntryByIDController{findTimeEntryByIDUseCase: findTimeEntryByIDUseCase, requireAuthMiddleware: requireAuthMiddleware}
}

func (findTimeEntryByIDController *FindTimeEntryByIDController) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc(constant.FindTimeEntryByIDRoutePath, findTimeEntryByIDController.requireAuthMiddleware.Wrap(findTimeEntryByIDController.FindByID))
}

// FindByID godoc
// @Summary Get a time entry
// @Description Returns a single time entry owned by the authenticated user
// @Tags time-entry
// @Produce json
// @Security BearerAuth
// @Param id path string true "Time entry id"
// @Success 200 {object} dto.TimeEntryResponseDTO
// @Failure 401 {object} sharedRest.ErrorResponseDTO
// @Failure 404 {object} sharedRest.ErrorResponseDTO
// @Router /api/v1/time-entries/{id} [get]
func (findTimeEntryByIDController *FindTimeEntryByIDController) FindByID(responseWriter http.ResponseWriter, request *http.Request) {
	authenticatedUser, _ := sharedMiddleware.AuthenticatedUserFromContext(request.Context())

	timeEntry, err := findTimeEntryByIDController.findTimeEntryByIDUseCase.FindByID(authenticatedUser.ID, request.PathValue("id"))
	if err != nil {
		sharedRest.WriteError(responseWriter, err)
		return
	}

	var responseDTO dto.TimeEntryResponseDTO = mapper.ToTimeEntryResponseDTO(timeEntry)

	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(responseDTO)
}
