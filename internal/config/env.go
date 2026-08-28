package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"talacert-api/internal/logger"

	"github.com/joho/godotenv"
)

// LoadEnv charge le fichier .env
func LoadEnv() (*Config, error) {

	if err := godotenv.Load(); err != nil {
		logger.ErrorLogger.Warn("No .env file found", "error", err)
		return nil, err
	}

	accessExp, err := durationEnv("JWT_ACCESS_EXPIRATION")
	if err != nil {
		return nil, err
	}

	refreshExp, err := durationEnv("JWT_REFRESH_EXPIRATION")
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		AppEnv:     mustEnv("APP_ENV"),
		GinMode:    mustEnv("GIN_MODE"),
		AppPort:    mustEnv("APP_PORT"),
		AppURL:     mustEnv("APP_URL"),
		DBHost:     mustEnv("DB_HOST"),
		DBPort:     mustEnv("DB_PORT"),
		DBUser:     mustEnv("DB_USER"),
		DBPassword: mustEnv("DB_PASSWORD"),
		DBName:     mustEnv("DB_NAME"),

		JWTAccessSecret:      mustEnv("JWT_ACCESS_SECRET"),
		JWTAccessExpiration:  accessExp,
		JWTRefreshSecret:     mustEnv("JWT_REFRESH_SECRET"),
		JWTRefreshExpiration: refreshExp,

		AdminDefaultUsername: mustEnv("ADMIN_DEFAULT_USERNAME"),
		AdminDefaultEmail:    mustEnv("ADMIN_DEFAULT_EMAIL"),
		AdminDefaultPassword: mustEnv("ADMIN_DEFAULT_PASSWORD"),
	}

	return cfg, nil
}

// durationEnv parses an environment variable as seconds.
func durationEnv(key string) (time.Duration, error) {
	value := mustEnv(key)

	seconds, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("%s must be a valid integer: %w", key, err)
	}

	return time.Duration(seconds) * time.Second, nil
}

// mustEnv returns the environment variable or panics if missing.
func mustEnv(key string) string {
	value := os.Getenv(key)

	if value == "" {
		panic(fmt.Errorf("%s environment variable is required", key))
	}

	return value
}

// GetEnv retourne une variable d'environnement ou une valeur par défaut.
func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}
