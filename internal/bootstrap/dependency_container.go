package bootstrap

import (
	"gorm.io/gorm"

	"talacert-api/internal/auth"
	"talacert-api/internal/config"
	"talacert-api/internal/handlers"
	"talacert-api/internal/repositories"
	"talacert-api/internal/routes"
	"talacert-api/internal/services"
	"talacert-api/internal/utils"
)

type Dependencies struct {
	APIHandler *routes.APIHandler
	JWTManager *auth.JWTManager
}

func BuildDependencyContainer(
	cfg *config.Config,
	db *gorm.DB,
) *Dependencies {

	jwtManager := auth.NewJWTManager(
		cfg.JWTAccessSecret,
		cfg.JWTRefreshSecret,
		cfg.JWTAccessExpiration,
		cfg.JWTRefreshExpiration,
	)

	// Repositories
	documentRepository := repositories.NewDocumentRepository(db)
	documentSequenceRepository := repositories.NewDocumentSequenceRepository(db)
	authRepository := repositories.NewAuthRepository(db)
	userRepository := repositories.NewUserRepository(db)

	// Services
	qr_service := services.NewQRService(cfg.AppURL)

	documentService := services.NewDocument(
		documentRepository,
		documentSequenceRepository,
		&utils.HashGenerator{},
		qr_service, // Replace with actual QRService initialization if needed
	)
	authService := services.NewAuthService(authRepository, jwtManager)
	userService := services.NewUserService(userRepository)

	// Handlers
	authHandler := handlers.NewAuth(authService)
	documentHandler := handlers.NewDocument(documentService)
	userHandler := handlers.NewUser(userService)
	healthHandler := handlers.NewHealth()

	apiHandler := &routes.APIHandler{
		AuthHandler:     authHandler,
		DocumentHandler: documentHandler,
		UserHandler:     userHandler,
		HealthHandler:   healthHandler,
		JWTManager:      jwtManager,
	}

	return &Dependencies{

		APIHandler: apiHandler,
		JWTManager: jwtManager,
	}
}
