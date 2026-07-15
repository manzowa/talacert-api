package seed

import (
	"talacert-api/internal/config"

	"gorm.io/gorm"
)

type SeedManager struct {
	Config *config.Config
	DB     *gorm.DB
}

func NewSeedManager(
	cfg *config.Config,
	db *gorm.DB,
) *SeedManager {

	return &SeedManager{
		Config: cfg,
		DB:     db,
	}
}

func (m SeedManager) Run() error {

	if err := m.SeedAdminUser(); err != nil {
		return err
	}

	// Ajouter d'autres seeds ici
	// SeedRoles(db)
	// SeedPermissions(db)
	// SeedDocumentTypes(db)

	return nil
}
