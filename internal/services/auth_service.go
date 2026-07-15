package services

import (
	"context"
	"errors"
	"time"

	"talacert-api/internal/auth"
	"talacert-api/internal/dto"
	"talacert-api/internal/models"
	"talacert-api/internal/repositories"
	"talacert-api/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials  = errors.New("invalid credentials")
	ErrInvalidRefreshToken = errors.New("invalid refresh token")
	ErrRefreshTokenRevoked = errors.New("refresh token revoked")
)

type AuthService struct {
	AuthRepository *repositories.AuthRepository
	JWTManager     *auth.JWTManager
}

func NewAuthService(
	authRepository *repositories.AuthRepository,
	jwtManager *auth.JWTManager,
) *AuthService {
	return &AuthService{
		AuthRepository: authRepository,
		JWTManager:     jwtManager,
	}
}

func (s *AuthService) Login(
	cxt context.Context,
	email string,
	password string,
) (*dto.LoginResponse, error) {

	user, err := s.AuthRepository.FindUserByEmail(
		cxt,
		email,
	)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	); err != nil {

		return nil, ErrInvalidCredentials
	}

	accessToken, err := s.JWTManager.GenerateAccessToken(
		user.ID,
		user.Username,
		user.Email,
		string(user.Role),
	)

	if err != nil {
		return nil, err
	}

	refreshToken, err := s.JWTManager.GenerateRefreshToken(
		user.ID,
		user.Username,
		user.Email,
		string(user.Role),
	)

	if err != nil {
		return nil, err
	}
	hash := utils.NewHash()

	err = s.AuthRepository.SaveRefreshToken(
		cxt,
		&models.RefreshToken{
			UserID: user.ID,
			Token:  hash.GenerateHash(refreshToken),
			ExpiresAt: time.Now().Add(
				s.JWTManager.RefreshDuration(),
			),
		},
	)

	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    int64(s.JWTManager.AccessDuration().Seconds()),
	}, nil
}

func (s *AuthService) RefreshToken(
	cxt context.Context,
	token string,
) (*dto.LoginResponse, error) {

	hash := utils.NewHash()
	hashToken := hash.GenerateHash(token)

	stored, err := s.AuthRepository.GetRefreshToken(
		cxt,
		hashToken,
	)

	if err != nil || stored == nil {
		return nil, ErrInvalidRefreshToken
	}

	if stored.Revoked {
		return nil, ErrRefreshTokenRevoked
	}

	if stored.ExpiresAt.Before(time.Now()) {
		return nil, ErrInvalidRefreshToken
	}

	claims, err := s.JWTManager.VerifyRefreshToken(token)
	if err != nil {
		return nil, ErrInvalidRefreshToken
	}

	accessToken, err := s.JWTManager.GenerateAccessToken(
		claims.UserID,
		claims.Username,
		claims.Email,
		claims.Role,
	)

	if err != nil {
		return nil, err
	}

	newRefreshToken, err := s.JWTManager.GenerateRefreshToken(
		claims.UserID,
		claims.Username,
		claims.Email,
		claims.Role,
	)

	if err != nil {
		return nil, err
	}

	if err := s.AuthRepository.RevokeRefreshToken(
		cxt,
		hashToken,
	); err != nil {
		return nil, err
	}

	// Sauvegarde du nouveau refresh token hashé
	if err := s.AuthRepository.SaveRefreshToken(
		cxt,
		&models.RefreshToken{
			UserID: claims.UserID,
			Token:  hash.GenerateHash(newRefreshToken),
			ExpiresAt: time.Now().Add(
				s.JWTManager.RefreshDuration(),
			),
		},
	); err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    int64(s.JWTManager.AccessDuration().Seconds()),
	}, nil
}

func (s *AuthService) ValidateToken(
	token string,
) (bool, error) {

	_, err := s.JWTManager.VerifyAccessToken(
		token,
	)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *AuthService) GetUserFromToken(
	token string,
) (*auth.CustomClaims, error) {

	return s.JWTManager.VerifyAccessToken(
		token,
	)
}

func (s *AuthService) Logout(
	cxt context.Context,
	refreshToken string,
) error {

	hash := utils.NewHash()

	return s.AuthRepository.RevokeRefreshToken(
		cxt,
		hash.GenerateHash(refreshToken),
	)
}
