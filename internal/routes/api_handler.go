package routes

import (
	"talacert-api/internal/auth"
	"talacert-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

type APIHandler struct {
	AuthHandler     *handlers.AuthHandler
	DocumentHandler *handlers.DocumentHandler
	UserHandler     *handlers.UserHandler
	HealthHandler   *handlers.HealthHandler

	JWTManager *auth.JWTManager
}

func (h *APIHandler) Register(router *gin.Engine) {

	// Route Static
	//router.Static("/docs", "./docs")

	api := router.Group("/api/v1")
	apiSwagger := router.Group("/swagger")

	// Health check public
	router.GET("/health", h.HealthHandler.HealthCheck)

	// Routes Publiques
	h.registerAuthRoutes(api)
	h.registerSwaggerRoutes(apiSwagger)

	// Routes protégées
	protected := api.Group("")
	protected.Use(auth.JWTMiddleware(h.JWTManager))

	h.registerProtectedAuthRoutes(protected)
	h.registerUserRoutes(protected)
	h.registerDocumentRoutes(protected)

}
