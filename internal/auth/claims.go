package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

// Claims represents the JWT claims used for authentication.
type CustomClaims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`

	jwt.RegisteredClaims
}
