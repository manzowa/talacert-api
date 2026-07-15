package routes

import (
	"github.com/gin-gonic/gin"

	_ "talacert-api/docs"
)

func (h *APIHandler) registerSwaggerRoutes(
	router *gin.RouterGroup,
) {
	swagger := router.Group("")

	// swagger.GET(
	// 	"/doc/*any",
	// 	ginSwagger.WrapHandler(swaggerFiles.Handler),
	// )

	swagger.StaticFS("/doc", gin.Dir("./swagger-ui", true))
}
