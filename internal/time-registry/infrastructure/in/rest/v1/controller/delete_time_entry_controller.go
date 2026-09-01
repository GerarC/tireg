package controller

import (
	"net/http"

	sharedRest "github.com/gerarc/tireg/internal/shared/infrastructure/in/rest"
	sharedMiddleware "github.com/gerarc/tireg/internal/shared/infrastructure/in/rest/middleware"

	"github.com/gerarc/tireg/internal/time-registry/domain/api"
	"github.com/gerarc/tireg/internal/time-registry/infrastructure/in/rest/v1/util/constant"
)

type DeleteTimeEntryController struct {
	deleteTimeEntryUseCase api.DeleteTimeEntryUseCase
	requireAuthMiddleware  *sharedMiddleware.RequireAuthMiddleware
}

func NewDeleteTimeEntryController(deleteTimeEntryUseCase api.DeleteTimeEntryUseCase, requireAuthMiddleware *sharedMiddleware.RequireAuthMiddleware) *DeleteTimeEntryController {
	return &DeleteTimeEntryController{deleteTimeEntryUseCase: deleteTimeEntryUseCase, requireAuthMiddleware: requireAuthMiddleware}
}

func (deleteTimeEntryController *DeleteTimeEntryController) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc(constant.DeleteTimeEntryRoutePath, deleteTimeEntryController.requireAuthMiddleware.Wrap(deleteTimeEntryController.Delete))
}

// Delete godoc
// @Summary Delete a time entry
// @Description Deletes a time entry owned by the authenticated user
// @Tags time-entry
// @Security BearerAuth
// @Param id path string true "Time entry id"
// @Success 204
// @Failure 401 {object} sharedRest.ErrorResponseDTO
// @Failure 404 {object} sharedRest.ErrorResponseDTO
// @Router /api/v1/time-entries/{id} [delete]
func (deleteTimeEntryController *DeleteTimeEntryController) Delete(responseWriter http.ResponseWriter, request *http.Request) {
	authenticatedUser, _ := sharedMiddleware.AuthenticatedUserFromContext(request.Context())

	if err := deleteTimeEntryController.deleteTimeEntryUseCase.Delete(authenticatedUser.ID, request.PathValue("id")); err != nil {
		sharedRest.WriteError(responseWriter, err)
		return
	}

	responseWriter.WriteHeader(http.StatusNoContent)
}
