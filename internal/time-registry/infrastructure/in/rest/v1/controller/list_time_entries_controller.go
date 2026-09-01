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

type ListTimeEntriesController struct {
	listTimeEntriesUseCase api.ListTimeEntriesUseCase
	requireAuthMiddleware  *sharedMiddleware.RequireAuthMiddleware
}

func NewListTimeEntriesController(listTimeEntriesUseCase api.ListTimeEntriesUseCase, requireAuthMiddleware *sharedMiddleware.RequireAuthMiddleware) *ListTimeEntriesController {
	return &ListTimeEntriesController{listTimeEntriesUseCase: listTimeEntriesUseCase, requireAuthMiddleware: requireAuthMiddleware}
}

func (listTimeEntriesController *ListTimeEntriesController) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc(constant.ListTimeEntriesRoutePath, listTimeEntriesController.requireAuthMiddleware.Wrap(listTimeEntriesController.List))
}

// List godoc
// @Summary List time entries
// @Description Returns every time entry owned by the authenticated user
// @Tags time-entry
// @Produce json
// @Security BearerAuth
// @Success 200 {array} dto.TimeEntryResponseDTO
// @Failure 401 {object} sharedRest.ErrorResponseDTO
// @Router /api/v1/time-entries [get]
func (listTimeEntriesController *ListTimeEntriesController) List(responseWriter http.ResponseWriter, request *http.Request) {
	authenticatedUser, _ := sharedMiddleware.AuthenticatedUserFromContext(request.Context())

	timeEntries, err := listTimeEntriesController.listTimeEntriesUseCase.List(authenticatedUser.ID)
	if err != nil {
		sharedRest.WriteError(responseWriter, err)
		return
	}

	var responseDTOs []dto.TimeEntryResponseDTO = mapper.ToTimeEntryResponseDTOList(timeEntries)

	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(responseDTOs)
}
