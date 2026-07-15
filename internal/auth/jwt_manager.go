package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTManagerInterface interface {
	GenerateAccessToken(
		userID uint,
		username string,
		email string,
		role string,
	) (string, error)

	GenerateRefreshToken(
		userID uint,
		username string,
		email string,
		role string,
	) (string, error)

	VerifyAccessToken(
		tokenString string,
	) (*CustomClaims, error)

	VerifyRefreshToken(
		tokenString string,
	) (*CustomClaims, error)

	AccessDuration() time.Duration

	RefreshDuration() time.Duration

	GetClaims(
		token string,
	) (*CustomClaims, error)

	IsAccessTokenValid(
		token string,
	) bool

	IsRefreshTokenValid(
		token string,
	) bool
}

type JWTManager struct {
	accessSecret    string
	refreshSecret   string
	accessDuration  time.Duration
	refreshDuration time.Duration
}

func NewJWTManager(
	accessSecret string,
	refreshSecret string,
	accessDuration time.Duration,
	refreshDuration time.Duration,
) *JWTManager {
	return &JWTManager{
		accessSecret:    accessSecret,
		refreshSecret:   refreshSecret,
		accessDuration:  accessDuration,
		refreshDuration: refreshDuration,
	}
}

func (j *JWTManager) AccessDuration() time.Duration {
	return j.accessDuration
}

func (j *JWTManager) RefreshDuration() time.Duration {
	return j.refreshDuration
}

func (j *JWTManager) GenerateAccessToken(
	userID uint,
	username string,
	email string,
	role string,
) (string, error) {

	return j.generateToken(
		userID,
		username,
		email,
		role,
		j.accessSecret,
		j.accessDuration,
		"access",
	)
}

func (j *JWTManager) GenerateRefreshToken(
	userID uint,
	username string,
	email string,
	role string,
) (string, error) {

	return j.generateToken(
		userID,
		username,
		email,
		role,
		j.refreshSecret,
		j.refreshDuration,
		"refresh",
	)
}

func (j *JWTManager) generateToken(
	userID uint,
	username string,
	email string,
	role string,
	secret string,
	duration time.Duration,
	subject string,
) (string, error) {

	now := time.Now()

	claims := &CustomClaims{
		UserID:   userID,
		Username: username,
		Email:    email,
		Role:     role,

		RegisteredClaims: jwt.RegisteredClaims{
			Subject: subject,
			Issuer:  "talacert-api",

			IssuedAt: jwt.NewNumericDate(now),

			NotBefore: jwt.NewNumericDate(now),

			ExpiresAt: jwt.NewNumericDate(
				now.Add(duration),
			),
		},
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString([]byte(secret))
}

func (j *JWTManager) VerifyAccessToken(
	tokenString string,
) (*CustomClaims, error) {

	return j.verifyToken(
		tokenString,
		j.accessSecret,
		"access",
	)
}
func (j *JWTManager) verifyToken(
	tokenString string,
	secret string,
	expectedSubject string,
) (*CustomClaims, error) {

	token, err := jwt.ParseWithClaims(
		tokenString,
		&CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}

			return []byte(secret), nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*CustomClaims)

	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	if claims.Subject != expectedSubject {
		return nil, errors.New("invalid token subject")
	}

	return claims, nil
}
func (j *JWTManager) VerifyRefreshToken(
	tokenString string,
) (*CustomClaims, error) {

	return j.verifyToken(
		tokenString,
		j.refreshSecret,
		"refresh",
	)
}

func (j *JWTManager) GetClaims(
	token string,
) (*CustomClaims, error) {

	return j.VerifyAccessToken(token)
}

func (j *JWTManager) IsAccessTokenValid(
	token string,
) bool {

	_, err := j.VerifyAccessToken(token)

	return err == nil
}

func (j *JWTManager) IsRefreshTokenValid(
	token string,
) bool {

	_, err := j.VerifyRefreshToken(token)

	return err == nil
}
