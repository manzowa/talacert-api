package routes

import (
	"talacert-api/internal/auth"
	"talacert-api/internal/constants"

	"github.com/gin-gonic/gin"
)

func (h *APIHandler) registerDocumentRoutes(
	router *gin.RouterGroup,
) {

	documents := router.Group("/documents")

	documents.GET(
		"",
		auth.RequireRole(
			constants.RoleUser.String(),
			constants.RoleManager.String(),
			constants.RoleAdmin.String(),
		),
		h.DocumentHandler.GetAll,
	)

	documents.GET(
		"/:document_id",
		auth.RequireRole(
			constants.RoleUser.String(),
			constants.RoleManager.String(),
			constants.RoleAdmin.String(),
		),
		h.DocumentHandler.GetByDocument,
	)

	documents.POST(
		"",
		auth.RequireRole(
			constants.RoleManager.String(),
			constants.RoleAdmin.String(),
		),
		h.DocumentHandler.Create,
	)

	documents.PUT(
		"/:document_id",
		auth.RequireRole(
			constants.RoleManager.String(),
			constants.RoleAdmin.String(),
		),
		h.DocumentHandler.Update,
	)

	documents.DELETE(
		"/:document_id",
		auth.RequireRole(
			constants.RoleAdmin.String(),
		),
		h.DocumentHandler.Delete,
	)

	documents.GET(
		"/by-hash/:hash",
		auth.RequireRole(
			constants.RoleUser.String(),
			constants.RoleManager.String(),
			constants.RoleAdmin.String(),
		),
		h.DocumentHandler.GetByHash,
	)

	documents.GET(
		"/verify",
		auth.RequireRole(
			constants.RoleUser.String(),
			constants.RoleManager.String(),
			constants.RoleAdmin.String(),
		),
		h.DocumentHandler.Check,
	)
}
