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

type CreateTimeEntryController struct {
	createTimeEntryUseCase api.CreateTimeEntryUseCase
	requireAuthMiddleware  *sharedMiddleware.RequireAuthMiddleware
}

func NewCreateTimeEntryController(createTimeEntryUseCase api.CreateTimeEntryUseCase, requireAuthMiddleware *sharedMiddleware.RequireAuthMiddleware) *CreateTimeEntryController {
	return &CreateTimeEntryController{createTimeEntryUseCase: createTimeEntryUseCase, requireAuthMiddleware: requireAuthMiddleware}
}

func (createTimeEntryController *CreateTimeEntryController) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc(constant.CreateTimeEntryRoutePath, createTimeEntryController.requireAuthMiddleware.Wrap(createTimeEntryController.Create))
}

// Create godoc
// @Summary Create a time entry
// @Description Creates a new time entry owned by the authenticated user, auto-filling project/type/issue from the owner's task mappings when left blank
// @Tags time-entry
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.TimeEntryRequestDTO true "Time entry payload"
// @Success 201 {object} dto.TimeEntryResponseDTO
// @Failure 400 {object} sharedRest.ErrorResponseDTO
// @Failure 401 {object} sharedRest.ErrorResponseDTO
// @Router /api/v1/time-entries [post]
func (createTimeEntryController *CreateTimeEntryController) Create(responseWriter http.ResponseWriter, request *http.Request) {
	authenticatedUser, _ := sharedMiddleware.AuthenticatedUserFromContext(request.Context())

	var requestDTO dto.TimeEntryRequestDTO
	if err := json.NewDecoder(request.Body).Decode(&requestDTO); err != nil {
		sharedRest.WriteError(responseWriter, sharedRest.InvalidRequestBodyError(err))
		return
	}

	timeEntry, err := createTimeEntryController.createTimeEntryUseCase.Create(authenticatedUser.ID, mapper.ToTimeEntry(requestDTO))
	if err != nil {
		sharedRest.WriteError(responseWriter, err)
		return
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusCreated)
	json.NewEncoder(responseWriter).Encode(mapper.ToTimeEntryResponseDTO(timeEntry))
}
