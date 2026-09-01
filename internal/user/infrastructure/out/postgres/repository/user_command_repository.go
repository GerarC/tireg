package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/gerarc/tireg/internal/user/infrastructure/out/postgres/entity"
	"github.com/gerarc/tireg/internal/user/infrastructure/out/postgres/util/constant"
)

type UserCommandRepository struct {
	pool *pgxpool.Pool
}

func NewUserCommandRepository(pool *pgxpool.Pool) *UserCommandRepository {
	return &UserCommandRepository{pool: pool}
}

func (userCommandRepository *UserCommandRepository) Insert(ctx context.Context, userEntity entity.UserEntity) (entity.UserEntity, error) {
	row := userCommandRepository.pool.QueryRow(
		ctx,
		constant.InsertUserQuery,
		userEntity.FirstName,
		userEntity.LastName,
		userEntity.Username,
		userEntity.Email,
		userEntity.PasswordHash,
		userEntity.CreatedBy,
		userEntity.UpdatedBy,
	)

	var inserted entity.UserEntity
	err := row.Scan(
		&inserted.ID,
		&inserted.FirstName,
		&inserted.LastName,
		&inserted.Username,
		&inserted.Email,
		&inserted.PasswordHash,
		&inserted.CreatedAt,
		&inserted.CreatedBy,
		&inserted.UpdatedAt,
		&inserted.UpdatedBy,
	)

	return inserted, err
}
