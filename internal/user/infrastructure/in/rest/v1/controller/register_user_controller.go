package controller

import (
	"encoding/json"
	"net/http"

	sharedRest "github.com/gerarc/tireg/internal/shared/infrastructure/in/rest"

	"github.com/gerarc/tireg/internal/user/domain/api"
	"github.com/gerarc/tireg/internal/user/infrastructure/in/rest/v1/dto"
	"github.com/gerarc/tireg/internal/user/infrastructure/in/rest/v1/mapper"
	"github.com/gerarc/tireg/internal/user/infrastructure/in/rest/v1/util/constant"
)

type RegisterUserController struct {
	registerUserUseCase api.RegisterUserUseCase
}

func NewRegisterUserController(registerUserUseCase api.RegisterUserUseCase) *RegisterUserController {
	return &RegisterUserController{registerUserUseCase: registerUserUseCase}
}

func (registerUserController *RegisterUserController) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc(constant.UserRegisterRoutePath, registerUserController.Register)
}

// Register godoc
// @Summary Register a new user
// @Description Creates a new user account with a hashed password
// @Tags user
// @Accept json
// @Produce json
// @Param request body dto.UserRegisterRequestDTO true "User registration payload"
// @Success 201 {object} dto.UserResponseDTO
// @Failure 400 {object} sharedRest.ErrorResponseDTO
// @Failure 409 {object} sharedRest.ErrorResponseDTO
// @Router /api/v1/users [post]
func (registerUserController *RegisterUserController) Register(responseWriter http.ResponseWriter, request *http.Request) {
	var requestDTO dto.UserRegisterRequestDTO
	if err := json.NewDecoder(request.Body).Decode(&requestDTO); err != nil {
		sharedRest.WriteError(responseWriter, sharedRest.InvalidRequestBodyError(err))
		return
	}

	user, err := registerUserController.registerUserUseCase.Register(mapper.ToUserRegistration(requestDTO))
	if err != nil {
		sharedRest.WriteError(responseWriter, err)
		return
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusCreated)
	json.NewEncoder(responseWriter).Encode(mapper.ToUserResponseDTO(user))
}
