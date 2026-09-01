package mapper_test

import (
	"testing"
	"time"

	sharedModel "github.com/gerarc/tireg/internal/shared/domain/model"

	"github.com/gerarc/tireg/internal/user/domain/model"
	"github.com/gerarc/tireg/internal/user/infrastructure/in/rest/v1/dto"
	"github.com/gerarc/tireg/internal/user/infrastructure/in/rest/v1/mapper"
)

func TestToUserRegistration_MapsAllFields(t *testing.T) {
	requestDTO := dto.UserRegisterRequestDTO{
		FirstName: "Ada",
		LastName:  "Lovelace",
		Username:  "ada",
		Email:     "ada@example.com",
		Password:  "super-secret",
	}

	registration := mapper.ToUserRegistration(requestDTO)

	if registration.FirstName != requestDTO.FirstName ||
		registration.LastName != requestDTO.LastName ||
		registration.Username != requestDTO.Username ||
		registration.Email != requestDTO.Email ||
		registration.Password != requestDTO.Password {
		t.Fatalf("expected registration to match request DTO, got %+v", registration)
	}
}

func TestToUserResponseDTO_DoesNotIncludePassword(t *testing.T) {
	user := model.User{
		ID:           "user-id",
		FirstName:    "Ada",
		LastName:     "Lovelace",
		Username:     "ada",
		Email:        "ada@example.com",
		PasswordHash: "hashed-password",
		Audit:        sharedModel.Audit{CreatedAt: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)},
	}

	responseDTO := mapper.ToUserResponseDTO(user)

	if responseDTO.ID != user.ID || responseDTO.Username != user.Username || responseDTO.Email != user.Email {
		t.Fatalf("expected mapped fields to match, got %+v", responseDTO)
	}
}
