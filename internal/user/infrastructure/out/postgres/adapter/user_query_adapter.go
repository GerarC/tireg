package adapter

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/gerarc/tireg/internal/user/domain/exception"
	"github.com/gerarc/tireg/internal/user/domain/model"
	"github.com/gerarc/tireg/internal/user/domain/spi"
	"github.com/gerarc/tireg/internal/user/infrastructure/out/postgres/mapper"
	"github.com/gerarc/tireg/internal/user/infrastructure/out/postgres/repository"
)

type UserQueryAdapter struct {
	userQueryRepository *repository.UserQueryRepository
}

func NewUserQueryAdapter(userQueryRepository *repository.UserQueryRepository) spi.UserQueryRepository {
	return &UserQueryAdapter{userQueryRepository: userQueryRepository}
}

func (userQueryAdapter *UserQueryAdapter) ExistsByUsername(username string) (bool, error) {
	return userQueryAdapter.userQueryRepository.SelectExistsByUsername(context.Background(), username)
}

func (userQueryAdapter *UserQueryAdapter) ExistsByEmail(email string) (bool, error) {
	return userQueryAdapter.userQueryRepository.SelectExistsByEmail(context.Background(), email)
}

func (userQueryAdapter *UserQueryAdapter) FindByIdentifier(identifier string) (model.User, error) {
	found, err := userQueryAdapter.userQueryRepository.SelectByIdentifier(context.Background(), identifier)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.User{}, exception.NewUserNotFoundError()
		}

		return model.User{}, err
	}

	return mapper.ToUserModel(found), nil
}
