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

type CreateGlossaryProjectController struct {
	createGlossaryProjectUseCase api.CreateGlossaryProjectUseCase
	requireAuthMiddleware        *sharedMiddleware.RequireAuthMiddleware
}

func NewCreateGlossaryProjectController(createGlossaryProjectUseCase api.CreateGlossaryProjectUseCase, requireAuthMiddleware *sharedMiddleware.RequireAuthMiddleware) *CreateGlossaryProjectController {
	return &CreateGlossaryProjectController{createGlossaryProjectUseCase: createGlossaryProjectUseCase, requireAuthMiddleware: requireAuthMiddleware}
}

func (createGlossaryProjectController *CreateGlossaryProjectController) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc(constant.CreateGlossaryProjectRoutePath, createGlossaryProjectController.requireAuthMiddleware.Wrap(createGlossaryProjectController.Create))
}

// Create godoc
// @Summary Create a glossary project
// @Description Creates a new glossary project owned by the authenticated user
// @Tags glossary
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.GlossaryProjectRequestDTO true "Glossary project payload"
// @Success 201 {object} dto.GlossaryProjectResponseDTO
// @Failure 400 {object} sharedRest.ErrorResponseDTO
// @Failure 401 {object} sharedRest.ErrorResponseDTO
// @Router /api/v1/glossary/projects [post]
func (createGlossaryProjectController *CreateGlossaryProjectController) Create(responseWriter http.ResponseWriter, request *http.Request) {
	authenticatedUser, _ := sharedMiddleware.AuthenticatedUserFromContext(request.Context())

	var requestDTO dto.GlossaryProjectRequestDTO
	if err := json.NewDecoder(request.Body).Decode(&requestDTO); err != nil {
		sharedRest.WriteError(responseWriter, sharedRest.InvalidRequestBodyError(err))
		return
	}

	glossaryProject, err := createGlossaryProjectController.createGlossaryProjectUseCase.Create(authenticatedUser.ID, mapper.ToGlossaryProject(requestDTO))
	if err != nil {
		sharedRest.WriteError(responseWriter, err)
		return
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusCreated)
	json.NewEncoder(responseWriter).Encode(mapper.ToGlossaryProjectResponseDTO(glossaryProject))
}
