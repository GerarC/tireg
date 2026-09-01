package controller

import (
	"net/http"

	sharedRest "github.com/gerarc/tireg/internal/shared/infrastructure/in/rest"
	sharedMiddleware "github.com/gerarc/tireg/internal/shared/infrastructure/in/rest/middleware"

	"github.com/gerarc/tireg/internal/task/domain/api"
	"github.com/gerarc/tireg/internal/task/infrastructure/in/rest/v1/util/constant"
)

type DeleteTaskMappingController struct {
	deleteTaskMappingUseCase api.DeleteTaskMappingUseCase
	requireAuthMiddleware    *sharedMiddleware.RequireAuthMiddleware
}

func NewDeleteTaskMappingController(deleteTaskMappingUseCase api.DeleteTaskMappingUseCase, requireAuthMiddleware *sharedMiddleware.RequireAuthMiddleware) *DeleteTaskMappingController {
	return &DeleteTaskMappingController{deleteTaskMappingUseCase: deleteTaskMappingUseCase, requireAuthMiddleware: requireAuthMiddleware}
}

func (deleteTaskMappingController *DeleteTaskMappingController) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc(constant.DeleteTaskMappingRoutePath, deleteTaskMappingController.requireAuthMiddleware.Wrap(deleteTaskMappingController.Delete))
}

// Delete godoc
// @Summary Delete a task mapping
// @Description Deletes a task mapping owned by the authenticated user
// @Tags task-mapping
// @Security BearerAuth
// @Param id path string true "Task mapping id"
// @Success 204
// @Failure 401 {object} sharedRest.ErrorResponseDTO
// @Failure 404 {object} sharedRest.ErrorResponseDTO
// @Router /api/v1/task-mappings/{id} [delete]
func (deleteTaskMappingController *DeleteTaskMappingController) Delete(responseWriter http.ResponseWriter, request *http.Request) {
	authenticatedUser, _ := sharedMiddleware.AuthenticatedUserFromContext(request.Context())

	if err := deleteTaskMappingController.deleteTaskMappingUseCase.Delete(authenticatedUser.ID, request.PathValue("id")); err != nil {
		sharedRest.WriteError(responseWriter, err)
		return
	}

	responseWriter.WriteHeader(http.StatusNoContent)
}
