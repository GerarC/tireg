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

type FindTaskMappingByIDController struct {
	findTaskMappingByIDUseCase api.FindTaskMappingByIDUseCase
	requireAuthMiddleware      *sharedMiddleware.RequireAuthMiddleware
}

func NewFindTaskMappingByIDController(findTaskMappingByIDUseCase api.FindTaskMappingByIDUseCase, requireAuthMiddleware *sharedMiddleware.RequireAuthMiddleware) *FindTaskMappingByIDController {
	return &FindTaskMappingByIDController{findTaskMappingByIDUseCase: findTaskMappingByIDUseCase, requireAuthMiddleware: requireAuthMiddleware}
}

func (findTaskMappingByIDController *FindTaskMappingByIDController) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc(constant.FindTaskMappingByIDRoutePath, findTaskMappingByIDController.requireAuthMiddleware.Wrap(findTaskMappingByIDController.FindByID))
}

// FindByID godoc
// @Summary Get a task mapping
// @Description Returns a single task mapping owned by the authenticated user
// @Tags task-mapping
// @Produce json
// @Security BearerAuth
// @Param id path string true "Task mapping id"
// @Success 200 {object} dto.TaskMappingResponseDTO
// @Failure 401 {object} sharedRest.ErrorResponseDTO
// @Failure 404 {object} sharedRest.ErrorResponseDTO
// @Router /api/v1/task-mappings/{id} [get]
func (findTaskMappingByIDController *FindTaskMappingByIDController) FindByID(responseWriter http.ResponseWriter, request *http.Request) {
	authenticatedUser, _ := sharedMiddleware.AuthenticatedUserFromContext(request.Context())

	taskMapping, err := findTaskMappingByIDController.findTaskMappingByIDUseCase.FindByID(authenticatedUser.ID, request.PathValue("id"))
	if err != nil {
		sharedRest.WriteError(responseWriter, err)
		return
	}

	var responseDTO dto.TaskMappingResponseDTO = mapper.ToTaskMappingResponseDTO(taskMapping)

	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(responseDTO)
}
