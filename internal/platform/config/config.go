package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	AppName                 string
	AppEnv                  string
	AppPort                 string
	DatabaseUrl             string
	ReadinessTimeoutSeconds int
}

func Load() (Config, error) {
	cfg := Config{
		AppName:                 getEnv("APP_NAME", "EstateHub API"),
		AppEnv:                  getEnv("APP_ENV", "development"),
		AppPort:                 getEnv("APP_PORT", "8080"),
		DatabaseUrl:             os.Getenv("DATABASE_URL"),
		ReadinessTimeoutSeconds: getEnvAsInt("READINESS_TIMEOUT_SECONDS", 2),
	}

	if cfg.DatabaseUrl == "" {
		return Config{}, fmt.Errorf("DATABASE_URL is required")
	}
	return cfg, nil
}

func (c Config) Addr() string {
	return ":" + c.AppPort
}

func (c Config) ReadinessTimeout() time.Duration {
	return time.Duration(c.ReadinessTimeoutSeconds) * time.Second
}

func getEnv(key string, fallBack string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallBack
}

func getEnvAsInt(key string, fallBack int) int {
	value := os.Getenv(key)

	if value == "" {
		return fallBack
	}

	parsed, err := strconv.Atoi(value)
	if err != nil {
		return fallBack
	}
	return parsed

}
