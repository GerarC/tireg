package mapper

import (
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
		Audit:        user.Audit,
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
		Audit:        userEntity.Audit,
	}
}
