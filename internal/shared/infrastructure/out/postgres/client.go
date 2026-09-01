package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/gerarc/tireg/internal/shared/application/utils/config"
	"github.com/gerarc/tireg/internal/shared/infrastructure/out/postgres/util/constant"
)

func GetClient(appConfig *config.Config) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf(
		constant.ConnectionStringFormat,
		appConfig.PostgresUser,
		appConfig.PostgresPassword,
		appConfig.PostgresHost,
		appConfig.PostgresPort,
		appConfig.PostgresDBName,
		appConfig.PostgresSSLMode,
	)

	return pgxpool.New(context.Background(), dsn)
}
