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

type CreateTaskMappingController struct {
	createTaskMappingUseCase api.CreateTaskMappingUseCase
	requireAuthMiddleware    *sharedMiddleware.RequireAuthMiddleware
}

func NewCreateTaskMappingController(createTaskMappingUseCase api.CreateTaskMappingUseCase, requireAuthMiddleware *sharedMiddleware.RequireAuthMiddleware) *CreateTaskMappingController {
	return &CreateTaskMappingController{createTaskMappingUseCase: createTaskMappingUseCase, requireAuthMiddleware: requireAuthMiddleware}
}

func (createTaskMappingController *CreateTaskMappingController) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc(constant.CreateTaskMappingRoutePath, createTaskMappingController.requireAuthMiddleware.Wrap(createTaskMappingController.Create))
}

// Create godoc
// @Summary Create a task mapping
// @Description Creates a new task/meeting classification rule owned by the authenticated user
// @Tags task-mapping
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.TaskMappingRequestDTO true "Task mapping payload"
// @Success 201 {object} dto.TaskMappingResponseDTO
// @Failure 400 {object} sharedRest.ErrorResponseDTO
// @Failure 401 {object} sharedRest.ErrorResponseDTO
// @Router /api/v1/task-mappings [post]
func (createTaskMappingController *CreateTaskMappingController) Create(responseWriter http.ResponseWriter, request *http.Request) {
	authenticatedUser, _ := sharedMiddleware.AuthenticatedUserFromContext(request.Context())

	var requestDTO dto.TaskMappingRequestDTO
	if err := json.NewDecoder(request.Body).Decode(&requestDTO); err != nil {
		sharedRest.WriteError(responseWriter, sharedRest.InvalidRequestBodyError(err))
		return
	}

	taskMapping, err := createTaskMappingController.createTaskMappingUseCase.Create(authenticatedUser.ID, mapper.ToTaskMapping(requestDTO))
	if err != nil {
		sharedRest.WriteError(responseWriter, err)
		return
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusCreated)
	json.NewEncoder(responseWriter).Encode(mapper.ToTaskMappingResponseDTO(taskMapping))
}
