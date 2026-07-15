package bootstrap

import (
	"os"

	"talacert-api/internal/logger"

	"github.com/gin-gonic/gin"
)

func ConfigureHTTP(mode string) {

	switch mode {

	case gin.ReleaseMode:
		gin.SetMode(gin.ReleaseMode)

	default:
		gin.SetMode(gin.DebugMode)
	}
}

func Run(
	app *Application,
) {

	logger.AppLogger.Info(
		"Starting HTTP server",
		"port",
		app.Config.Port,
	)

	err := app.Router.Run(
		":" + app.Config.Port,
	)

	if err != nil {

		logger.ErrorLogger.Error(
			"HTTP server failed",
			"error",
			err,
		)

		os.Exit(1)
	}
}
