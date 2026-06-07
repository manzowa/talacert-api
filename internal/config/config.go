package config

import (
	"os"
	"talacert-api/internal/logger"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv  string
	GinMode string
	Port    string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

var AppConfig *Config

func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		logger.ErrorLogger.Error("Error loading .env file", "error", err)
	}

	AppConfig = &Config{
		AppEnv:  os.Getenv("APP_ENV"),
		GinMode: os.Getenv("GIN_MODE"),
		Port:    os.Getenv("PORT"),

		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
	}
}

func GetConfig() *Config {
	return AppConfig
}
