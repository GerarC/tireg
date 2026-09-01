package controller

import (
	"encoding/json"
	"net/http"

	sharedRest "github.com/gerarc/tireg/internal/shared/infrastructure/in/rest"

	"github.com/gerarc/tireg/internal/auth/domain/api"
	"github.com/gerarc/tireg/internal/auth/infrastructure/in/rest/v1/dto"
	"github.com/gerarc/tireg/internal/auth/infrastructure/in/rest/v1/mapper"
	"github.com/gerarc/tireg/internal/auth/infrastructure/in/rest/v1/util/constant"
)

type LoginController struct {
	loginUseCase api.LoginUseCase
}

func NewLoginController(loginUseCase api.LoginUseCase) *LoginController {
	return &LoginController{loginUseCase: loginUseCase}
}

func (loginController *LoginController) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc(constant.LoginRoutePath, loginController.Login)
}

// Login godoc
// @Summary Authenticate a user
// @Description Validates credentials and returns a signed JWT access token
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.LoginRequestDTO true "Login credentials"
// @Success 200 {object} dto.LoginResponseDTO
// @Failure 400 {object} sharedRest.ErrorResponseDTO
// @Failure 401 {object} sharedRest.ErrorResponseDTO
// @Router /api/v1/auth/login [post]
func (loginController *LoginController) Login(responseWriter http.ResponseWriter, request *http.Request) {
	var requestDTO dto.LoginRequestDTO
	if err := json.NewDecoder(request.Body).Decode(&requestDTO); err != nil {
		sharedRest.WriteError(responseWriter, sharedRest.InvalidRequestBodyError(err))
		return
	}

	accessToken, err := loginController.loginUseCase.Login(mapper.ToCredentials(requestDTO))
	if err != nil {
		sharedRest.WriteError(responseWriter, err)
		return
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(mapper.ToLoginResponseDTO(accessToken))
}
