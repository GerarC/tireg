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

type ListTaskMappingsController struct {
	listTaskMappingsUseCase api.ListTaskMappingsUseCase
	requireAuthMiddleware   *sharedMiddleware.RequireAuthMiddleware
}

func NewListTaskMappingsController(listTaskMappingsUseCase api.ListTaskMappingsUseCase, requireAuthMiddleware *sharedMiddleware.RequireAuthMiddleware) *ListTaskMappingsController {
	return &ListTaskMappingsController{listTaskMappingsUseCase: listTaskMappingsUseCase, requireAuthMiddleware: requireAuthMiddleware}
}

func (listTaskMappingsController *ListTaskMappingsController) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc(constant.ListTaskMappingsRoutePath, listTaskMappingsController.requireAuthMiddleware.Wrap(listTaskMappingsController.List))
}

// List godoc
// @Summary List task mappings
// @Description Returns every task mapping owned by the authenticated user
// @Tags task-mapping
// @Produce json
// @Security BearerAuth
// @Success 200 {array} dto.TaskMappingResponseDTO
// @Failure 401 {object} sharedRest.ErrorResponseDTO
// @Router /api/v1/task-mappings [get]
func (listTaskMappingsController *ListTaskMappingsController) List(responseWriter http.ResponseWriter, request *http.Request) {
	authenticatedUser, _ := sharedMiddleware.AuthenticatedUserFromContext(request.Context())

	taskMappings, err := listTaskMappingsController.listTaskMappingsUseCase.List(authenticatedUser.ID)
	if err != nil {
		sharedRest.WriteError(responseWriter, err)
		return
	}

	var responseDTOs []dto.TaskMappingResponseDTO = mapper.ToTaskMappingResponseDTOList(taskMappings)

	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(responseDTOs)
}
