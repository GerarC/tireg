package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/gerarc/tireg/internal/user/infrastructure/out/postgres/entity"
	"github.com/gerarc/tireg/internal/user/infrastructure/out/postgres/util/constant"
)

type UserQueryRepository struct {
	pool *pgxpool.Pool
}

func NewUserQueryRepository(pool *pgxpool.Pool) *UserQueryRepository {
	return &UserQueryRepository{pool: pool}
}

func (userQueryRepository *UserQueryRepository) SelectExistsByUsername(ctx context.Context, username string) (bool, error) {
	var exists bool
	err := userQueryRepository.pool.QueryRow(ctx, constant.ExistsByUsernameQuery, username).Scan(&exists)
	return exists, err
}

func (userQueryRepository *UserQueryRepository) SelectExistsByEmail(ctx context.Context, email string) (bool, error) {
	var exists bool
	err := userQueryRepository.pool.QueryRow(ctx, constant.ExistsByEmailQuery, email).Scan(&exists)
	return exists, err
}

func (userQueryRepository *UserQueryRepository) SelectByIdentifier(ctx context.Context, identifier string) (entity.UserEntity, error) {
	row := userQueryRepository.pool.QueryRow(ctx, constant.SelectByIdentifierQuery, identifier)

	var found entity.UserEntity
	err := row.Scan(
		&found.ID,
		&found.FirstName,
		&found.LastName,
		&found.Username,
		&found.Email,
		&found.PasswordHash,
		&found.CreatedAt,
		&found.CreatedBy,
		&found.UpdatedAt,
		&found.UpdatedBy,
	)

	return found, err
}
