package routes

import (
	"talacert-api/internal/auth"
	"talacert-api/internal/constants"

	"github.com/gin-gonic/gin"
)

func (h *APIHandler) registerUserRoutes(
	router *gin.RouterGroup,
) {

	users := router.Group("/users")

	users.GET(
		"",
		auth.RequireRole(
			constants.RoleManager.String(),
			constants.RoleAdmin.String(),
		),
		h.UserHandler.GetAll)

	users.GET(
		"/:id",
		auth.RequireRole(
			constants.RoleManager.String(),
			constants.RoleAdmin.String(),
		),
		h.UserHandler.GetByID)

	users.POST(
		"",
		auth.RequireRole(
			constants.RoleAdmin.String(),
		),
		h.UserHandler.Create)

	users.PUT(
		"/:id",
		auth.RequireRole(
			constants.RoleAdmin.String(),
		),
		h.UserHandler.Update)

	users.DELETE(
		"/:id",
		auth.RequireRole(
			constants.RoleAdmin.String(),
		),
		h.UserHandler.Delete)

	users.PUT(
		"/:id/role",
		auth.RequireRole(
			constants.RoleAdmin.String(),
		),
		h.UserHandler.ChangeRole)

	users.PUT(
		"/:id/password",
		auth.RequireRole(
			constants.RoleAdmin.String(),
		),
		h.UserHandler.ChangePassword)
}
