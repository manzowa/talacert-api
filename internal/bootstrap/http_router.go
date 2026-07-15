package bootstrap

import (
	"github.com/gin-gonic/gin"

	"talacert-api/internal/middleware"
	"talacert-api/internal/routes"
)

type Router interface {
	Run(addr ...string) error
}

func NewRouter(
	apiHandler *routes.APIHandler,
) *gin.Engine {

	router := gin.New()

	// Middlerware
	router.Use(gin.Recovery())
	router.Use(middleware.AccessLog())
	router.Use(middleware.ErrorLog())

	apiHandler.Register(router)

	return router
}
