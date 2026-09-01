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

type UpdateTimeEntryController struct {
	updateTimeEntryUseCase api.UpdateTimeEntryUseCase
	requireAuthMiddleware  *sharedMiddleware.RequireAuthMiddleware
}

func NewUpdateTimeEntryController(updateTimeEntryUseCase api.UpdateTimeEntryUseCase, requireAuthMiddleware *sharedMiddleware.RequireAuthMiddleware) *UpdateTimeEntryController {
	return &UpdateTimeEntryController{updateTimeEntryUseCase: updateTimeEntryUseCase, requireAuthMiddleware: requireAuthMiddleware}
}

func (updateTimeEntryController *UpdateTimeEntryController) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc(constant.UpdateTimeEntryRoutePath, updateTimeEntryController.requireAuthMiddleware.Wrap(updateTimeEntryController.Update))
}

// Update godoc
// @Summary Update a time entry
// @Description Updates a time entry owned by the authenticated user
// @Tags time-entry
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Time entry id"
// @Param request body dto.TimeEntryRequestDTO true "Time entry payload"
// @Success 200 {object} dto.TimeEntryResponseDTO
// @Failure 400 {object} sharedRest.ErrorResponseDTO
// @Failure 401 {object} sharedRest.ErrorResponseDTO
// @Failure 404 {object} sharedRest.ErrorResponseDTO
// @Router /api/v1/time-entries/{id} [put]
func (updateTimeEntryController *UpdateTimeEntryController) Update(responseWriter http.ResponseWriter, request *http.Request) {
	authenticatedUser, _ := sharedMiddleware.AuthenticatedUserFromContext(request.Context())

	var requestDTO dto.TimeEntryRequestDTO
	if err := json.NewDecoder(request.Body).Decode(&requestDTO); err != nil {
		sharedRest.WriteError(responseWriter, sharedRest.InvalidRequestBodyError(err))
		return
	}

	timeEntry, err := updateTimeEntryController.updateTimeEntryUseCase.Update(authenticatedUser.ID, request.PathValue("id"), mapper.ToTimeEntry(requestDTO))
	if err != nil {
		sharedRest.WriteError(responseWriter, err)
		return
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(mapper.ToTimeEntryResponseDTO(timeEntry))
}
