package mapper

import (
	"time"

	"github.com/gerarc/tireg/internal/auth/domain/model"
	"github.com/gerarc/tireg/internal/auth/infrastructure/in/rest/v1/dto"
)

func ToCredentials(requestDTO dto.LoginRequestDTO) model.Credentials {
	return model.Credentials{
		Identifier: requestDTO.Identifier,
		Password:   requestDTO.Password,
	}
}

func ToLoginResponseDTO(accessToken model.AccessToken) dto.LoginResponseDTO {
	return dto.LoginResponseDTO{
		AccessToken: accessToken.Token,
		ExpiresAt:   accessToken.ExpiresAt.Format(time.RFC3339),
	}
}
