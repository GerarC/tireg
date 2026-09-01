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

type ListGlossaryTypesController struct {
	listGlossaryTypesUseCase api.ListGlossaryTypesUseCase
	requireAuthMiddleware    *sharedMiddleware.RequireAuthMiddleware
}

func NewListGlossaryTypesController(listGlossaryTypesUseCase api.ListGlossaryTypesUseCase, requireAuthMiddleware *sharedMiddleware.RequireAuthMiddleware) *ListGlossaryTypesController {
	return &ListGlossaryTypesController{listGlossaryTypesUseCase: listGlossaryTypesUseCase, requireAuthMiddleware: requireAuthMiddleware}
}

func (listGlossaryTypesController *ListGlossaryTypesController) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc(constant.ListGlossaryTypesRoutePath, listGlossaryTypesController.requireAuthMiddleware.Wrap(listGlossaryTypesController.List))
}

// List godoc
// @Summary List glossary types
// @Description Returns every glossary type owned by the authenticated user, seeding the defaults on first access
// @Tags glossary
// @Produce json
// @Security BearerAuth
// @Success 200 {array} dto.GlossaryTypeResponseDTO
// @Failure 401 {object} sharedRest.ErrorResponseDTO
// @Router /api/v1/glossary/types [get]
func (listGlossaryTypesController *ListGlossaryTypesController) List(responseWriter http.ResponseWriter, request *http.Request) {
	authenticatedUser, _ := sharedMiddleware.AuthenticatedUserFromContext(request.Context())

	glossaryTypes, err := listGlossaryTypesController.listGlossaryTypesUseCase.List(authenticatedUser.ID)
	if err != nil {
		sharedRest.WriteError(responseWriter, err)
		return
	}

	var responseDTOs []dto.GlossaryTypeResponseDTO = mapper.ToGlossaryTypeResponseDTOList(glossaryTypes)

	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(responseDTOs)
}
