package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/gerarc/tireg/internal/user/infrastructure/out/postgres/entity"
)

type UserQueryRepository struct {
	db *gorm.DB
}

func NewUserQueryRepository(db *gorm.DB) *UserQueryRepository {
	return &UserQueryRepository{db: db}
}

func (userQueryRepository *UserQueryRepository) SelectExistsByUsername(ctx context.Context, username string) (bool, error) {
	var count int64
	err := userQueryRepository.db.WithContext(ctx).Model(&entity.UserEntity{}).Where("username = ?", username).Count(&count).Error
	return count > 0, err
}

func (userQueryRepository *UserQueryRepository) SelectExistsByEmail(ctx context.Context, email string) (bool, error) {
	var count int64
	err := userQueryRepository.db.WithContext(ctx).Model(&entity.UserEntity{}).Where("email = ?", email).Count(&count).Error
	return count > 0, err
}

func (userQueryRepository *UserQueryRepository) SelectByIdentifier(ctx context.Context, identifier string) (entity.UserEntity, error) {
	var found entity.UserEntity
	err := userQueryRepository.db.WithContext(ctx).Where("username = ? OR email = ?", identifier, identifier).First(&found).Error
	return found, err
}
