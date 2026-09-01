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

type ListGlossaryProjectsController struct {
	listGlossaryProjectsUseCase api.ListGlossaryProjectsUseCase
	requireAuthMiddleware       *sharedMiddleware.RequireAuthMiddleware
}

func NewListGlossaryProjectsController(listGlossaryProjectsUseCase api.ListGlossaryProjectsUseCase, requireAuthMiddleware *sharedMiddleware.RequireAuthMiddleware) *ListGlossaryProjectsController {
	return &ListGlossaryProjectsController{listGlossaryProjectsUseCase: listGlossaryProjectsUseCase, requireAuthMiddleware: requireAuthMiddleware}
}

func (listGlossaryProjectsController *ListGlossaryProjectsController) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc(constant.ListGlossaryProjectsRoutePath, listGlossaryProjectsController.requireAuthMiddleware.Wrap(listGlossaryProjectsController.List))
}

// List godoc
// @Summary List glossary projects
// @Description Returns every glossary project owned by the authenticated user
// @Tags glossary
// @Produce json
// @Security BearerAuth
// @Success 200 {array} dto.GlossaryProjectResponseDTO
// @Failure 401 {object} sharedRest.ErrorResponseDTO
// @Router /api/v1/glossary/projects [get]
func (listGlossaryProjectsController *ListGlossaryProjectsController) List(responseWriter http.ResponseWriter, request *http.Request) {
	authenticatedUser, _ := sharedMiddleware.AuthenticatedUserFromContext(request.Context())

	glossaryProjects, err := listGlossaryProjectsController.listGlossaryProjectsUseCase.List(authenticatedUser.ID)
	if err != nil {
		sharedRest.WriteError(responseWriter, err)
		return
	}

	var responseDTOs []dto.GlossaryProjectResponseDTO = mapper.ToGlossaryProjectResponseDTOList(glossaryProjects)

	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(responseDTOs)
}
