package mapper

import (
	"time"

	"github.com/gerarc/tireg/internal/user/domain/model"
	"github.com/gerarc/tireg/internal/user/infrastructure/in/rest/v1/dto"
)

func ToUserRegistration(requestDTO dto.UserRegisterRequestDTO) model.UserRegistration {
	return model.UserRegistration{
		FirstName: requestDTO.FirstName,
		LastName:  requestDTO.LastName,
		Username:  requestDTO.Username,
		Email:     requestDTO.Email,
		Password:  requestDTO.Password,
	}
}

func ToUserResponseDTO(user model.User) dto.UserResponseDTO {
	return dto.UserResponseDTO{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		CreatedBy: user.CreatedBy,
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
		UpdatedBy: user.UpdatedBy,
	}
}
