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

type UpdateGlossaryTypeController struct {
	updateGlossaryTypeUseCase api.UpdateGlossaryTypeUseCase
	requireAuthMiddleware     *sharedMiddleware.RequireAuthMiddleware
}

func NewUpdateGlossaryTypeController(updateGlossaryTypeUseCase api.UpdateGlossaryTypeUseCase, requireAuthMiddleware *sharedMiddleware.RequireAuthMiddleware) *UpdateGlossaryTypeController {
	return &UpdateGlossaryTypeController{updateGlossaryTypeUseCase: updateGlossaryTypeUseCase, requireAuthMiddleware: requireAuthMiddleware}
}

func (updateGlossaryTypeController *UpdateGlossaryTypeController) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc(constant.UpdateGlossaryTypeRoutePath, updateGlossaryTypeController.requireAuthMiddleware.Wrap(updateGlossaryTypeController.Update))
}

// Update godoc
// @Summary Update a glossary type
// @Description Updates a glossary type owned by the authenticated user
// @Tags glossary
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Glossary type id"
// @Param request body dto.GlossaryTypeRequestDTO true "Glossary type payload"
// @Success 200 {object} dto.GlossaryTypeResponseDTO
// @Failure 400 {object} sharedRest.ErrorResponseDTO
// @Failure 401 {object} sharedRest.ErrorResponseDTO
// @Failure 404 {object} sharedRest.ErrorResponseDTO
// @Router /api/v1/glossary/types/{id} [put]
func (updateGlossaryTypeController *UpdateGlossaryTypeController) Update(responseWriter http.ResponseWriter, request *http.Request) {
	authenticatedUser, _ := sharedMiddleware.AuthenticatedUserFromContext(request.Context())

	var requestDTO dto.GlossaryTypeRequestDTO
	if err := json.NewDecoder(request.Body).Decode(&requestDTO); err != nil {
		sharedRest.WriteError(responseWriter, sharedRest.InvalidRequestBodyError(err))
		return
	}

	glossaryType, err := updateGlossaryTypeController.updateGlossaryTypeUseCase.Update(authenticatedUser.ID, request.PathValue("id"), mapper.ToGlossaryType(requestDTO))
	if err != nil {
		sharedRest.WriteError(responseWriter, err)
		return
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(mapper.ToGlossaryTypeResponseDTO(glossaryType))
}
