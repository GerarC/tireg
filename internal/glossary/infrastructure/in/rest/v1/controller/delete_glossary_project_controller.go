package controller

import (
	"net/http"

	sharedRest "github.com/gerarc/tireg/internal/shared/infrastructure/in/rest"
	sharedMiddleware "github.com/gerarc/tireg/internal/shared/infrastructure/in/rest/middleware"

	"github.com/gerarc/tireg/internal/glossary/domain/api"
	"github.com/gerarc/tireg/internal/glossary/infrastructure/in/rest/v1/util/constant"
)

type DeleteGlossaryProjectController struct {
	deleteGlossaryProjectUseCase api.DeleteGlossaryProjectUseCase
	requireAuthMiddleware        *sharedMiddleware.RequireAuthMiddleware
}

func NewDeleteGlossaryProjectController(deleteGlossaryProjectUseCase api.DeleteGlossaryProjectUseCase, requireAuthMiddleware *sharedMiddleware.RequireAuthMiddleware) *DeleteGlossaryProjectController {
	return &DeleteGlossaryProjectController{deleteGlossaryProjectUseCase: deleteGlossaryProjectUseCase, requireAuthMiddleware: requireAuthMiddleware}
}

func (deleteGlossaryProjectController *DeleteGlossaryProjectController) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc(constant.DeleteGlossaryProjectRoutePath, deleteGlossaryProjectController.requireAuthMiddleware.Wrap(deleteGlossaryProjectController.Delete))
}

// Delete godoc
// @Summary Delete a glossary project
// @Description Deletes a glossary project owned by the authenticated user
// @Tags glossary
// @Security BearerAuth
// @Param id path string true "Glossary project id"
// @Success 204
// @Failure 401 {object} sharedRest.ErrorResponseDTO
// @Failure 404 {object} sharedRest.ErrorResponseDTO
// @Router /api/v1/glossary/projects/{id} [delete]
func (deleteGlossaryProjectController *DeleteGlossaryProjectController) Delete(responseWriter http.ResponseWriter, request *http.Request) {
	authenticatedUser, _ := sharedMiddleware.AuthenticatedUserFromContext(request.Context())

	if err := deleteGlossaryProjectController.deleteGlossaryProjectUseCase.Delete(authenticatedUser.ID, request.PathValue("id")); err != nil {
		sharedRest.WriteError(responseWriter, err)
		return
	}

	responseWriter.WriteHeader(http.StatusNoContent)
}
