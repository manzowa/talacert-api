package repositories

import (
	"context"
	"errors"
	"time"

	"talacert-api/internal/logger"
	"talacert-api/internal/models"

	"gorm.io/gorm"
)

type AuthRepositoryInterface interface {
	FindUserByEmail(cxt context.Context, email string) (*models.User, error)
	SaveRefreshToken(cxt context.Context, token *models.RefreshToken) error
	GetRefreshToken(cxt context.Context, token string) (*models.RefreshToken, error)
	RevokeRefreshToken(cxt context.Context, token string) error
}

type AuthRepository struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		DB: db,
	}
}

func (r *AuthRepository) CreateRefreshToken(
	cxt context.Context,
	refreshToken *models.RefreshToken,
) error {

	return r.DB.
		WithContext(cxt).
		Create(refreshToken).
		Error
}

func (r *AuthRepository) FindUserByEmail(
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

func (r *AuthRepository) SaveRefreshToken(
	cxt context.Context,
	token *models.RefreshToken,
) error {

	if err := r.DB.WithContext(cxt).Create(token).Error; err != nil {
		logger.AccessLogger.Error("failed to create Token", "error", err)
		return err
	}

	return nil
}

func (r *AuthRepository) GetRefreshToken(
	cxt context.Context,
	token string,
) (*models.RefreshToken, error) {

	var refreshToken models.RefreshToken

	err := r.DB.
		WithContext(cxt).
		Where(
			"token = ? AND revoked = ? AND expires_at > ?",
			token,
			false,
			time.Now(),
		).
		First(&refreshToken).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &refreshToken, nil
}

func (r *AuthRepository) RevokeRefreshToken(
	cxt context.Context,
	token string,
) error {

	result := r.DB.
		WithContext(cxt).
		Model(&models.RefreshToken{}).
		Where("token = ? AND revoked = ?", token, false).
		Update("revoked", true)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *AuthRepository) DeleteRefreshToken(
	cxt context.Context,
	token string,
) error {

	result := r.DB.
		WithContext(cxt).
		Where("token = ?", token).
		Delete(&models.RefreshToken{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
