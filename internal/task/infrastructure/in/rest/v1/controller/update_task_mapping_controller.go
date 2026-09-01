package controller

import (
	"encoding/json"
	"net/http"

	sharedRest "github.com/gerarc/tireg/internal/shared/infrastructure/in/rest"
	sharedMiddleware "github.com/gerarc/tireg/internal/shared/infrastructure/in/rest/middleware"

	"github.com/gerarc/tireg/internal/task/domain/api"
	"github.com/gerarc/tireg/internal/task/infrastructure/in/rest/v1/dto"
	"github.com/gerarc/tireg/internal/task/infrastructure/in/rest/v1/mapper"
	"github.com/gerarc/tireg/internal/task/infrastructure/in/rest/v1/util/constant"
)

type UpdateTaskMappingController struct {
	updateTaskMappingUseCase api.UpdateTaskMappingUseCase
	requireAuthMiddleware    *sharedMiddleware.RequireAuthMiddleware
}

func NewUpdateTaskMappingController(updateTaskMappingUseCase api.UpdateTaskMappingUseCase, requireAuthMiddleware *sharedMiddleware.RequireAuthMiddleware) *UpdateTaskMappingController {
	return &UpdateTaskMappingController{updateTaskMappingUseCase: updateTaskMappingUseCase, requireAuthMiddleware: requireAuthMiddleware}
}

func (updateTaskMappingController *UpdateTaskMappingController) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc(constant.UpdateTaskMappingRoutePath, updateTaskMappingController.requireAuthMiddleware.Wrap(updateTaskMappingController.Update))
}

// Update godoc
// @Summary Update a task mapping
// @Description Updates a task mapping owned by the authenticated user
// @Tags task-mapping
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Task mapping id"
// @Param request body dto.TaskMappingRequestDTO true "Task mapping payload"
// @Success 200 {object} dto.TaskMappingResponseDTO
// @Failure 400 {object} sharedRest.ErrorResponseDTO
// @Failure 401 {object} sharedRest.ErrorResponseDTO
// @Failure 404 {object} sharedRest.ErrorResponseDTO
// @Router /api/v1/task-mappings/{id} [put]
func (updateTaskMappingController *UpdateTaskMappingController) Update(responseWriter http.ResponseWriter, request *http.Request) {
	authenticatedUser, _ := sharedMiddleware.AuthenticatedUserFromContext(request.Context())

	var requestDTO dto.TaskMappingRequestDTO
	if err := json.NewDecoder(request.Body).Decode(&requestDTO); err != nil {
		sharedRest.WriteError(responseWriter, sharedRest.InvalidRequestBodyError(err))
		return
	}

	taskMapping, err := updateTaskMappingController.updateTaskMappingUseCase.Update(authenticatedUser.ID, request.PathValue("id"), mapper.ToTaskMapping(requestDTO))
	if err != nil {
		sharedRest.WriteError(responseWriter, err)
		return
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(mapper.ToTaskMappingResponseDTO(taskMapping))
}
