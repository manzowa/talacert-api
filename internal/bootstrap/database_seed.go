package bootstrap

import (
	"os"

	"gorm.io/gorm"

	"talacert-api/internal/config"
	"talacert-api/internal/logger"
	"talacert-api/internal/seed"
)

func SeedDatabase(
	cfg *config.Config,
	db *gorm.DB,
) {

	managerSeed := seed.NewSeedManager(cfg, db)

	err := managerSeed.Run()

	if err != nil {
		logger.ErrorLogger.Error(
			"Database seed failed",
			"error", err,
		)

		os.Exit(1)
	}

	logger.AppLogger.Info(
		"Database seed completed",
	)

}
