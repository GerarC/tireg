package controller

import (
	"net/http"

	sharedRest "github.com/gerarc/tireg/internal/shared/infrastructure/in/rest"
	sharedMiddleware "github.com/gerarc/tireg/internal/shared/infrastructure/in/rest/middleware"

	"github.com/gerarc/tireg/internal/glossary/domain/api"
	"github.com/gerarc/tireg/internal/glossary/infrastructure/in/rest/v1/util/constant"
)

type DeleteGlossaryTypeController struct {
	deleteGlossaryTypeUseCase api.DeleteGlossaryTypeUseCase
	requireAuthMiddleware     *sharedMiddleware.RequireAuthMiddleware
}

func NewDeleteGlossaryTypeController(deleteGlossaryTypeUseCase api.DeleteGlossaryTypeUseCase, requireAuthMiddleware *sharedMiddleware.RequireAuthMiddleware) *DeleteGlossaryTypeController {
	return &DeleteGlossaryTypeController{deleteGlossaryTypeUseCase: deleteGlossaryTypeUseCase, requireAuthMiddleware: requireAuthMiddleware}
}

func (deleteGlossaryTypeController *DeleteGlossaryTypeController) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc(constant.DeleteGlossaryTypeRoutePath, deleteGlossaryTypeController.requireAuthMiddleware.Wrap(deleteGlossaryTypeController.Delete))
}

// Delete godoc
// @Summary Delete a glossary type
// @Description Deletes a glossary type owned by the authenticated user
// @Tags glossary
// @Security BearerAuth
// @Param id path string true "Glossary type id"
// @Success 204
// @Failure 401 {object} sharedRest.ErrorResponseDTO
// @Failure 404 {object} sharedRest.ErrorResponseDTO
// @Router /api/v1/glossary/types/{id} [delete]
func (deleteGlossaryTypeController *DeleteGlossaryTypeController) Delete(responseWriter http.ResponseWriter, request *http.Request) {
	authenticatedUser, _ := sharedMiddleware.AuthenticatedUserFromContext(request.Context())

	if err := deleteGlossaryTypeController.deleteGlossaryTypeUseCase.Delete(authenticatedUser.ID, request.PathValue("id")); err != nil {
		sharedRest.WriteError(responseWriter, err)
		return
	}

	responseWriter.WriteHeader(http.StatusNoContent)
}
