package config

import (
	"os"
	"strconv"
	"time"

	"github.com/gerarc/tireg/internal/shared/application/utils/config/util/constant"
)

type Config struct {
	Port string

	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDBName   string
	PostgresSSLMode  string

	JWTSecret     string
	JWTExpiration time.Duration
}

func Load() Config {
	_ = LoadEnvFile(constant.EnvFileName)

	jwtExpirationMinutes, err := strconv.Atoi(getEnvOrDefault(constant.JWTExpirationMinutesEnvVar, constant.DefaultJWTExpirationMinutes))
	if err != nil {
		jwtExpirationMinutes, _ = strconv.Atoi(constant.DefaultJWTExpirationMinutes)
	}

	return Config{
		Port: getEnvOrDefault(constant.PortEnvVar, constant.DefaultPort),

		PostgresHost:     getEnvOrDefault(constant.PostgresHostEnvVar, constant.DefaultPostgresHost),
		PostgresPort:     getEnvOrDefault(constant.PostgresPortEnvVar, constant.DefaultPostgresPort),
		PostgresUser:     getEnvOrDefault(constant.PostgresUserEnvVar, constant.DefaultPostgresUser),
		PostgresPassword: os.Getenv(constant.PostgresPasswordEnvVar),
		PostgresDBName:   getEnvOrDefault(constant.PostgresDBNameEnvVar, constant.DefaultPostgresDBName),
		PostgresSSLMode:  getEnvOrDefault(constant.PostgresSSLModeEnvVar, constant.DefaultPostgresSSLMode),

		JWTSecret:     os.Getenv(constant.JWTSecretEnvVar),
		JWTExpiration: time.Duration(jwtExpirationMinutes) * time.Minute,
	}
}

func NewConfig() *Config {
	loadedConfig := Load()
	return &loadedConfig
}

func getEnvOrDefault(envVar string, defaultValue string) string {
	value := os.Getenv(envVar)
	if value == "" {
		return defaultValue
	}

	return value
}
