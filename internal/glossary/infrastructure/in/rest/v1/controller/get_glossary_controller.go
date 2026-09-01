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

type GetGlossaryController struct {
	getGlossaryUseCase    api.GetGlossaryUseCase
	requireAuthMiddleware *sharedMiddleware.RequireAuthMiddleware
}

func NewGetGlossaryController(getGlossaryUseCase api.GetGlossaryUseCase, requireAuthMiddleware *sharedMiddleware.RequireAuthMiddleware) *GetGlossaryController {
	return &GetGlossaryController{getGlossaryUseCase: getGlossaryUseCase, requireAuthMiddleware: requireAuthMiddleware}
}

func (getGlossaryController *GetGlossaryController) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc(constant.GetGlossaryRoutePath, getGlossaryController.requireAuthMiddleware.Wrap(getGlossaryController.Get))
}

// Get godoc
// @Summary Get the authenticated user's glossary
// @Description Returns every glossary type and project owned by the authenticated user
// @Tags glossary
// @Produce json
// @Security BearerAuth
// @Success 200 {object} dto.GlossaryResponseDTO
// @Failure 401 {object} sharedRest.ErrorResponseDTO
// @Router /api/v1/glossary [get]
func (getGlossaryController *GetGlossaryController) Get(responseWriter http.ResponseWriter, request *http.Request) {
	authenticatedUser, _ := sharedMiddleware.AuthenticatedUserFromContext(request.Context())

	glossary, err := getGlossaryController.getGlossaryUseCase.Get(authenticatedUser.ID)
	if err != nil {
		sharedRest.WriteError(responseWriter, err)
		return
	}

	var responseDTO dto.GlossaryResponseDTO = mapper.ToGlossaryResponseDTO(glossary)

	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(responseDTO)
}
