package routes

import (
	"talacert-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

type APIHandler struct {
	DocumentHandler *handlers.DocumentHandler
}

// SetupRoutes sets up the API routes for the Gin router
func (apiHandler *APIHandler) SetupRoutes(router *gin.Engine) {
	// Group routes under /api/v1
	api := router.Group("/api/v1")
	{
		// Define your API routes here
		api.GET("/documents", apiHandler.DocumentHandler.GetDocumentsHandler)
		api.POST("/documents", apiHandler.DocumentHandler.PostDocumentHandler)
		api.GET("/documents/:id", apiHandler.DocumentHandler.GetDocumentHandler)
		api.PUT("/documents/:id", apiHandler.DocumentHandler.PutDocumentHandler)
		api.DELETE("/documents/:id", apiHandler.DocumentHandler.DeleteDocumentHandler)

		api.GET("/documents/hash/:hash", apiHandler.DocumentHandler.GetDocumentByHashHandler)
	}
}
