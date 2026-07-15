package bootstrap

import (
	"talacert-api/internal/config"
	"talacert-api/internal/logger"

	"gorm.io/gorm"
)

type Application struct {
	Config *config.Config
	DB     *gorm.DB

	Dependencies *Dependencies
	Router       Router
}

func Initialize() *Application {

	logger.AppLogger.Info("Initializing application")

	cfg := MustLoadConfig()

	ConfigureHTTP(cfg.GinMode)

	db := MustConnectDatabase(cfg)

	MigrateDatabse(db)

	SeedDatabase(cfg, db)

	deps := BuildDependencyContainer(cfg, db)

	router := NewRouter(deps.APIHandler)

	return &Application{

		Config: cfg,

		DB: db,

		Dependencies: deps,

		Router: router,
	}
}
