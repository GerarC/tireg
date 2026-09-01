package adapter

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgconn"

	"github.com/gerarc/tireg/internal/user/domain/exception"
	"github.com/gerarc/tireg/internal/user/domain/model"
	"github.com/gerarc/tireg/internal/user/domain/spi"
	domainConstant "github.com/gerarc/tireg/internal/user/domain/util/constant"
	"github.com/gerarc/tireg/internal/user/infrastructure/out/postgres/mapper"
	"github.com/gerarc/tireg/internal/user/infrastructure/out/postgres/repository"
	pgConstant "github.com/gerarc/tireg/internal/user/infrastructure/out/postgres/util/constant"
)

const uniqueViolationErrorCode = "23505"

type UserCommandAdapter struct {
	userCommandRepository *repository.UserCommandRepository
}

func NewUserCommandAdapter(userCommandRepository *repository.UserCommandRepository) spi.UserCommandRepository {
	return &UserCommandAdapter{userCommandRepository: userCommandRepository}
}

func (userCommandAdapter *UserCommandAdapter) Save(user model.User) (model.User, error) {
	inserted, err := userCommandAdapter.userCommandRepository.Insert(context.Background(), mapper.ToUserEntity(user))
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == uniqueViolationErrorCode {
			switch pgErr.ConstraintName {
			case pgConstant.UsernameUniqueConstraint:
				return model.User{}, exception.NewAlreadyTakenError(domainConstant.DetailUsernameAlreadyTaken)
			case pgConstant.EmailUniqueConstraint:
				return model.User{}, exception.NewAlreadyTakenError(domainConstant.DetailEmailAlreadyTaken)
			}
		}

		return model.User{}, err
	}

	return mapper.ToUserModel(inserted), nil
}
