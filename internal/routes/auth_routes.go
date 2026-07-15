package routes

import (
	"github.com/gin-gonic/gin"
)

func (h *APIHandler) registerAuthRoutes(
	router *gin.RouterGroup,
) {

	auth := router.Group("/auth")

	auth.POST("/login", h.AuthHandler.Login)
	auth.POST("/refresh", h.AuthHandler.Refresh)
	auth.POST("/logout", h.AuthHandler.Logout)

}
