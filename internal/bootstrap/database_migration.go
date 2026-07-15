package bootstrap

import (
	"os"

	"gorm.io/gorm"

	"talacert-api/internal/logger"
	"talacert-api/internal/models"
)

func MigrateDatabse(
	db *gorm.DB,
) {

	err := db.
		Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&models.User{},
			&models.RefreshToken{},
			&models.Document{},
			&models.DocumentSequence{},
		)

	if err != nil {

		logger.ErrorLogger.Error(
			"Database migration failed",
			"error", err,
		)

		os.Exit(1)
	}

	logger.AppLogger.Info(
		"Database migration completed",
	)
}
