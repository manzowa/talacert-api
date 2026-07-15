package main

import (
	"talacert-api/internal/bootstrap"
	"talacert-api/internal/logger"
)

func main() {

	logger.Init()

	app := bootstrap.Initialize()

	bootstrap.Run(app)

}
