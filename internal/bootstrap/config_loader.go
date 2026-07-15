package bootstrap

import (
	"os"
	"talacert-api/internal/config"
	"talacert-api/internal/logger"
)

func MustLoadConfig() *config.Config {

	if err := config.Init(); err != nil {
		logger.ErrorLogger.Error(
			"Configuration initialization failed",
			"error", err,
		)

		os.Exit(1)
	}

	cfg := config.Get()

	return cfg
}
