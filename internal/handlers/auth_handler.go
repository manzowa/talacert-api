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

// Login godoc
// @Summary      Connexion
// @Description  Authentifie un utilisateur et retourne les tokens JWT.
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        credentials  body  dto.LoginRequest  true  "Identifiants"
// @Success      200  {object}  dto.LoginResponse
// @Failure      401  {object} interface{} "Authentication failed"
// @Router       /api/v1/auth/login [post]
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

// Refresh godoc
// @Summary      Refresh access token
// @Description  Génère un nouvel access token à partir d'un refresh token valide.
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        request body dto.LogoutRequest true "Refresh token"
// @Success      200 {object} interface{} "Token refreshed successfully"
// @Failure      400 {object} interface{} "Invalid request payload"
// @Failure      401 {object} interface{} "Token refresh failed"
// @Router       /api/v1/auth/refresh [post]
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

// Logout godoc
// @Summary      Logout
// @Description  Invalide le refresh token de l'utilisateur et termine sa session.
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        request body dto.LogoutRequest true "Refresh token"
// @Success      200 {object} interface{} "Logout successful"
// @Failure      400 {object} interface{} "Invalid request payload"
// @Failure      500 {object} interface{} "Logout failed"
// @Router       /api/v1/auth/logout [post]
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

// Me godoc
// @Summary      Get current user
// @Description  Retourne les informations de l'utilisateur actuellement authentifié.
// @Tags         Authentication
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} dto.UserResponse "User retrieved successfully"
// @Failure      401 {object} interface{} "User not found in context"
// @Failure      500 {object} interface{} "Invalid user data in context"
// @Router       /api/v1/auth/me [get]
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
