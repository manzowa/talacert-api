package config

import "time"

type Config struct {
	// Application Configuration
	AppEnv  string
	GinMode string
	Port    string

	// Database Configuration
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	// JWT Configuration
	JWTAccessSecret      string
	JWTAccessExpiration  time.Duration
	JWTRefreshSecret     string
	JWTRefreshExpiration time.Duration

	// ADMIN Configuration
	AdminDefaultUsername string
	AdminDefaultEmail    string
	AdminDefaultPassword string
}

var AppConfig *Config

// Init initializes the application configuration.
func Init() error {

	cfg, err := LoadEnv()
	if err != nil {
		return err
	}

	AppConfig = cfg
	return nil
}

// Get returns the loaded application configuration.
func Get() *Config {
	return AppConfig
}
