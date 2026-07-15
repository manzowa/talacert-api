package seed

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"talacert-api/internal/constants"
	"talacert-api/internal/logger"
	"talacert-api/internal/models"
)

func (s SeedManager) SeedAdminUser() error {

	var count int64

	err := s.DB.Model(&models.User{}).
		Where("email = ?", s.Config.AdminDefaultEmail).
		Count(&count).Error

	if err != nil {
		return err
	}

	if count > 0 {

		logger.AppLogger.Info(
			"Admin user already exists",
		)

		return nil
	}

	password, err := bcrypt.GenerateFromPassword(
		[]byte(s.Config.AdminDefaultPassword),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	admin := models.User{
		Username: s.Config.AdminDefaultUsername,
		Email:    s.Config.AdminDefaultEmail,
		Password: string(password),
		Role:     constants.RoleAdmin,
	}

	err = s.DB.Create(&admin).Error

	if err != nil {
		return errors.New("unable to create admin user")
	}

	logger.AppLogger.Info(
		"Default administrator created",
		"email", admin.Email,
	)

	return nil
}
