package main

import (
	"github.com/gin-gonic/gin"

	"talacert-api/internal/config"
	"talacert-api/internal/handlers"
	"talacert-api/internal/logger"
	"talacert-api/internal/middleware"
	"talacert-api/internal/models"
	"talacert-api/internal/repositories"
	"talacert-api/internal/routes"
	"talacert-api/internal/services"
)

func main() {

	//-----------------------------------
	// LOGGER
	//-----------------------------------
	logger.Init()
	logger.AppLogger.Info("Starting application")

	//-----------------------------------
	// CONFIG
	//-----------------------------------
	config.InitConfig()

	//-----------------------------------
	// GIN MODE
	//-----------------------------------
	if config.AppConfig.GinMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	//-----------------------------------
	// DATABASE
	//-----------------------------------
	db, err := config.ConnectDatabase(config.GetConfig())
	if err != nil {
		logger.ErrorLogger.Error("Database connection failed", "error", err)
	}

	logger.AppLogger.Info("Database connected")

	//-----------------------------------
	// MIGRATION
	//-----------------------------------
	if err := db.AutoMigrate(
		&models.Document{},
		&models.DocumentSequence{},
	); err != nil {
		logger.ErrorLogger.Error("Failed migration", "error", err)
	}

	logger.AppLogger.Info("Database migrated")

	//-----------------------------------
	// DEPENDENCY INJECTION
	//-----------------------------------
	documentRepository := repositories.New(db)
	documentSequenceRepository := repositories.NewDocumentSequence(db)

	documentService := services.New(
		documentRepository,
		documentSequenceRepository,
	)

	documentHandler := handlers.New(documentService)

	//-----------------------------------
	// ROUTER
	//-----------------------------------
	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(middleware.AccessLog())
	router.Use(middleware.ErrorLog())

	//-----------------------------------
	// ROUTES
	//-----------------------------------
	apiHandler := &routes.APIHandler{
		DocumentHandler: documentHandler,
	}

	apiHandler.SetupRoutes(router)

	//-----------------------------------
	// START SERVER
	//-----------------------------------
	logger.AppLogger.Info(
		"Server started",
		"port", config.AppConfig.Port,
	)

	if err := router.Run(":" + config.AppConfig.Port); err != nil {
		logger.ErrorLogger.Error("Server failed", "error", err)
	}
}
