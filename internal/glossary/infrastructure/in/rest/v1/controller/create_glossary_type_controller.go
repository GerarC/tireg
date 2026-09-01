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

type CreateGlossaryTypeController struct {
	createGlossaryTypeUseCase api.CreateGlossaryTypeUseCase
	requireAuthMiddleware     *sharedMiddleware.RequireAuthMiddleware
}

func NewCreateGlossaryTypeController(createGlossaryTypeUseCase api.CreateGlossaryTypeUseCase, requireAuthMiddleware *sharedMiddleware.RequireAuthMiddleware) *CreateGlossaryTypeController {
	return &CreateGlossaryTypeController{createGlossaryTypeUseCase: createGlossaryTypeUseCase, requireAuthMiddleware: requireAuthMiddleware}
}

func (createGlossaryTypeController *CreateGlossaryTypeController) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc(constant.CreateGlossaryTypeRoutePath, createGlossaryTypeController.requireAuthMiddleware.Wrap(createGlossaryTypeController.Create))
}

// Create godoc
// @Summary Create a glossary type
// @Description Creates a new glossary type owned by the authenticated user
// @Tags glossary
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.GlossaryTypeRequestDTO true "Glossary type payload"
// @Success 201 {object} dto.GlossaryTypeResponseDTO
// @Failure 400 {object} sharedRest.ErrorResponseDTO
// @Failure 401 {object} sharedRest.ErrorResponseDTO
// @Router /api/v1/glossary/types [post]
func (createGlossaryTypeController *CreateGlossaryTypeController) Create(responseWriter http.ResponseWriter, request *http.Request) {
	authenticatedUser, _ := sharedMiddleware.AuthenticatedUserFromContext(request.Context())

	var requestDTO dto.GlossaryTypeRequestDTO
	if err := json.NewDecoder(request.Body).Decode(&requestDTO); err != nil {
		sharedRest.WriteError(responseWriter, sharedRest.InvalidRequestBodyError(err))
		return
	}

	glossaryType, err := createGlossaryTypeController.createGlossaryTypeUseCase.Create(authenticatedUser.ID, mapper.ToGlossaryType(requestDTO))
	if err != nil {
		sharedRest.WriteError(responseWriter, err)
		return
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusCreated)
	json.NewEncoder(responseWriter).Encode(mapper.ToGlossaryTypeResponseDTO(glossaryType))
}
