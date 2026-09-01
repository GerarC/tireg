package controller

import (
	"encoding/json"
	"net/http"

	sharedRest "github.com/gerarc/tireg/internal/shared/infrastructure/in/rest"
	sharedMiddleware "github.com/gerarc/tireg/internal/shared/infrastructure/in/rest/middleware"

	"github.com/gerarc/tireg/internal/glossary/domain/api"
	"github.com/gerarc/tireg/internal/glossary/infrastructure/in/rest/v1/dto"
	"github.com/gerarc/tireg/internal/glossary/infrastructure/in/rest/v1/mapper"
	"github.com/gerarc/tireg/internal/glossary/infrastructure/in/rest/v1/util/constant"
)

type UpdateGlossaryProjectController struct {
	updateGlossaryProjectUseCase api.UpdateGlossaryProjectUseCase
	requireAuthMiddleware        *sharedMiddleware.RequireAuthMiddleware
}

func NewUpdateGlossaryProjectController(updateGlossaryProjectUseCase api.UpdateGlossaryProjectUseCase, requireAuthMiddleware *sharedMiddleware.RequireAuthMiddleware) *UpdateGlossaryProjectController {
	return &UpdateGlossaryProjectController{updateGlossaryProjectUseCase: updateGlossaryProjectUseCase, requireAuthMiddleware: requireAuthMiddleware}
}

func (updateGlossaryProjectController *UpdateGlossaryProjectController) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc(constant.UpdateGlossaryProjectRoutePath, updateGlossaryProjectController.requireAuthMiddleware.Wrap(updateGlossaryProjectController.Update))
}

// Update godoc
// @Summary Update a glossary project
// @Description Updates a glossary project owned by the authenticated user
// @Tags glossary
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Glossary project id"
// @Param request body dto.GlossaryProjectRequestDTO true "Glossary project payload"
// @Success 200 {object} dto.GlossaryProjectResponseDTO
// @Failure 400 {object} sharedRest.ErrorResponseDTO
// @Failure 401 {object} sharedRest.ErrorResponseDTO
// @Failure 404 {object} sharedRest.ErrorResponseDTO
// @Router /api/v1/glossary/projects/{id} [put]
func (updateGlossaryProjectController *UpdateGlossaryProjectController) Update(responseWriter http.ResponseWriter, request *http.Request) {
	authenticatedUser, _ := sharedMiddleware.AuthenticatedUserFromContext(request.Context())

	var requestDTO dto.GlossaryProjectRequestDTO
	if err := json.NewDecoder(request.Body).Decode(&requestDTO); err != nil {
		sharedRest.WriteError(responseWriter, sharedRest.InvalidRequestBodyError(err))
		return
	}

	glossaryProject, err := updateGlossaryProjectController.updateGlossaryProjectUseCase.Update(authenticatedUser.ID, request.PathValue("id"), mapper.ToGlossaryProject(requestDTO))
	if err != nil {
		sharedRest.WriteError(responseWriter, err)
		return
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(mapper.ToGlossaryProjectResponseDTO(glossaryProject))
}
