package handlers

import (
	"strings"

	"github.com/gin-gonic/gin"

	"talacert-api/internal/auth"
	"talacert-api/internal/dto"
	"talacert-api/internal/services"
	"talacert-api/internal/utils"
)

type AuthHandler struct {
	Service *services.AuthService
}

func NewAuth(service *services.AuthService) *AuthHandler {
	return &AuthHandler{
		Service: service,
	}
}

func (h *AuthHandler) Login(cxt *gin.Context) {
	var req dto.LoginRequest

	if err := cxt.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(cxt, "Invalid request payload", err.Error())

		return
	}

	auth, err := h.Service.Login(
		cxt,
		req.Email,
		req.Password,
	)

	if err != nil {
		utils.Unauthorized(cxt, "Authentication failed")
		return
	}

	var response = dto.LoginResponse{
		AccessToken:  auth.AccessToken,
		RefreshToken: auth.RefreshToken,
		TokenType:    auth.TokenType,
		ExpiresIn:    auth.ExpiresIn,
	}

	utils.Ok(cxt, "Login successful", response)

}

func (h *AuthHandler) Refresh(cxt *gin.Context) {
	var req dto.LogoutRequest

	if err := cxt.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(cxt, "Invalid request payload", err.Error())
		return
	}

	response, err := h.Service.RefreshToken(
		cxt,
		req.RefreshToken,
	)

	if err != nil {
		utils.Unauthorized(cxt, "Token refresh failed")
		return
	}

	utils.Ok(cxt, "Token refreshed successfully", response)
}

func (h *AuthHandler) Logout(cxt *gin.Context) {
	var req dto.LogoutRequest

	if err := cxt.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(cxt, "Invalid request payload", err.Error())
		return
	}

	err := h.Service.Logout(
		cxt,
		req.RefreshToken,
	)
	if err != nil {

		utils.InternalServerError(cxt, "Logout failed", err.Error())
		return
	}

	utils.Ok(cxt, "Logout successful", nil)
}

func (h *AuthHandler) Me(cxt *gin.Context) {
	userValue, exists := cxt.Get(auth.ContextUser)
	if !exists {
		utils.Unauthorized(cxt, "User not found in context")
		return
	}

	user, ok := userValue.(dto.UserResponse)
	if !ok {
		utils.InternalServerError(cxt, "Invalid user data in context", nil)
		return
	}

	response := dto.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Role:     strings.ToLower(user.Role),
	}

	utils.Ok(cxt, "User retrieved successfully", response)
}
