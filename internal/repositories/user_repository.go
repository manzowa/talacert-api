package repositories

import (
	"context"
	"errors"

	"talacert-api/internal/logger"
	"talacert-api/internal/models"

	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	Create(cxt context.Context, user *models.User) error
	Update(cxt context.Context, user *models.User) error
	Delete(cxt context.Context, id uint) error

	FindByID(cxt context.Context, id uint) (*models.User, error)
	FindByEmail(cxt context.Context, email string) (*models.User, error)
	FindByUsername(cxt context.Context, username string) (*models.User, error)
	FindAll(cxt context.Context) ([]models.User, error)
}

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) Create(
	cxt context.Context,
	user *models.User,
) error {
	if err := r.DB.WithContext(cxt).Create(user).Error; err != nil {
		logger.AccessLogger.Error("failed to create user", "error", err)
		return err
	}

	return nil
}

func (r *UserRepository) Update(
	cxt context.Context,
	user *models.User,
) error {
	if err := r.DB.WithContext(cxt).Save(user).Error; err != nil {
		logger.AccessLogger.Error("failed to update user", "error", err)
		return err
	}

	return nil
}

func (r *UserRepository) Delete(
	cxt context.Context,
	id uint,
) error {
	if err := r.DB.WithContext(cxt).
		Delete(&models.User{}, id).Error; err != nil {
		logger.AccessLogger.Error("failed to delete user", "error", err)
		return err
	}

	return nil
}

func (r *UserRepository) FindByID(
	cxt context.Context,
	id uint,
) (*models.User, error) {
	var user models.User

	err := r.DB.WithContext(cxt).
		First(&user, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &user, err
}

func (r *UserRepository) FindByEmail(
	cxt context.Context,
	email string,
) (*models.User, error) {
	var user models.User

	err := r.DB.WithContext(cxt).
		Where("email = ?", email).
		First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &user, err
}

func (r *UserRepository) FindByUsername(
	cxt context.Context,
	username string,
) (*models.User, error) {
	var user models.User

	err := r.DB.WithContext(cxt).
		Where("username = ?", username).
		First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &user, err
}

func (r *UserRepository) FindAll(cxt context.Context) ([]models.User, error) {
	var users []models.User

	err := r.DB.WithContext(cxt).
		Find(&users).Error

	return users, err
}
