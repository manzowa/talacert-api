package bootstrap

import (
	"os"

	"gorm.io/gorm"

	"talacert-api/internal/config"
	"talacert-api/internal/logger"
)

func MustConnectDatabase(
	cfg *config.Config,
) *gorm.DB {

	db, err := config.ConnectDatabase(cfg)

	if err != nil {

		logger.ErrorLogger.Error(
			"Database connection failed",
			"error", err,
		)

		os.Exit(1)
	}

	logger.AppLogger.Info(
		"Database connected",
	)

	return db
}
