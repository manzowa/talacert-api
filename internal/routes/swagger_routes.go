package routes

import (
	"github.com/gin-gonic/gin"

	_ "talacert-api/docs"
)

func (h *APIHandler) registerSwaggerRoutes(
	router *gin.RouterGroup,
) {
	router.StaticFS("/docs", gin.Dir("./docs", true))
	router.StaticFS("/swagger-ui", gin.Dir("./swagger-ui", true))
}
