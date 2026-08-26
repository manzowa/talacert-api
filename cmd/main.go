package main

import (
	"talacert-api/internal/bootstrap"
	"talacert-api/internal/logger"
)

// @title TalaCert API
// @version 1.0
// @description API de vérification de documents officiels.
// @description TalaCert permet de créer, gérer et vérifier l'authenticité des certificats, diplômes et attestations.

// @host 127.0.0.1:8080
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {

	logger.Init()

	app := bootstrap.Initialize()

	bootstrap.Run(app)
}
