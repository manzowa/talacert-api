package routes

import "github.com/gin-gonic/gin"

func (h *APIHandler) registerProtectedAuthRoutes(router *gin.RouterGroup) {
	auth := router.Group("/auth")

	auth.GET("/me", h.AuthHandler.Me)
}
