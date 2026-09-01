package mapper

import (
	sharedModel "github.com/gerarc/tireg/internal/shared/domain/model"
	sharedEntity "github.com/gerarc/tireg/internal/shared/infrastructure/out/postgres/entity"

	"github.com/gerarc/tireg/internal/user/domain/model"
	"github.com/gerarc/tireg/internal/user/infrastructure/out/postgres/entity"
)

func ToUserEntity(user model.User) entity.UserEntity {
	return entity.UserEntity{
		ID:           user.ID,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Username:     user.Username,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		AuditEntity: sharedEntity.AuditEntity{
			CreatedAt: user.CreatedAt,
			CreatedBy: user.CreatedBy,
			UpdatedAt: user.UpdatedAt,
			UpdatedBy: user.UpdatedBy,
		},
	}
}

func ToUserModel(userEntity entity.UserEntity) model.User {
	return model.User{
		ID:           userEntity.ID,
		FirstName:    userEntity.FirstName,
		LastName:     userEntity.LastName,
		Username:     userEntity.Username,
		Email:        userEntity.Email,
		PasswordHash: userEntity.PasswordHash,
		Audit: sharedModel.Audit{
			CreatedAt: userEntity.CreatedAt,
			CreatedBy: userEntity.CreatedBy,
			UpdatedAt: userEntity.UpdatedAt,
			UpdatedBy: userEntity.UpdatedBy,
		},
	}
}
