package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/gerarc/tireg/internal/user/infrastructure/out/postgres/entity"
)

type UserCommandRepository struct {
	db *gorm.DB
}

func NewUserCommandRepository(db *gorm.DB) *UserCommandRepository {
	return &UserCommandRepository{db: db}
}

func (userCommandRepository *UserCommandRepository) Insert(ctx context.Context, userEntity entity.UserEntity) (entity.UserEntity, error) {
	err := userCommandRepository.db.WithContext(ctx).Create(&userEntity).Error
	return userEntity, err
}
