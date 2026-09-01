package postgres

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gerarc/tireg/internal/shared/application/utils/config"
	"github.com/gerarc/tireg/internal/shared/application/utils/logger"
	"github.com/gerarc/tireg/internal/shared/infrastructure/out/postgres/util/constant"
)

func GetClient(appConfig *config.Config, appLogger logger.Logger) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		constant.ConnectionStringFormat,
		appConfig.PostgresUser,
		appConfig.PostgresPassword,
		appConfig.PostgresHost,
		appConfig.PostgresPort,
		appConfig.PostgresDBName,
		appConfig.PostgresSSLMode,
	)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: NewGormLoggerAdapter(appLogger)})
}
