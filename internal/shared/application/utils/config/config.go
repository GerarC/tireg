package config

import (
	"os"

	"github.com/gerarc/tireg/internal/shared/application/utils/config/util/constant"
)

type Config struct {
	Port string
}

func Load() Config {
	_ = LoadEnvFile(constant.EnvFileName)

	port := os.Getenv(constant.PortEnvVar)
	if port == "" {
		port = constant.DefaultPort
	}

	return Config{Port: port}
}
