package postgres

import (
	"net"
	"net/url"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gerarc/tireg/internal/shared/application/utils/config"
	"github.com/gerarc/tireg/internal/shared/application/utils/logger"
	"github.com/gerarc/tireg/internal/shared/infrastructure/out/postgres/util/constant"
)

func BuildDSN(appConfig *config.Config) string {
	dsn := url.URL{
		Scheme: constant.Scheme,
		User:   url.UserPassword(appConfig.PostgresUser, appConfig.PostgresPassword),
		Host:   net.JoinHostPort(appConfig.PostgresHost, appConfig.PostgresPort),
		Path:   "/" + appConfig.PostgresDBName,
		RawQuery: url.Values{
			constant.SSLModeQueryParam: {appConfig.PostgresSSLMode},
		}.Encode(),
	}

	return dsn.String()
}

func GetClient(appConfig *config.Config, appLogger logger.Logger) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(BuildDSN(appConfig)), &gorm.Config{Logger: NewGormLoggerAdapter(appLogger)})
}
