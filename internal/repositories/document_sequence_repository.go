package repositories

import (
	"talacert-api/internal/models"

	"gorm.io/gorm"
)

type DocumentSequenceRepositoryInterface interface {
	GetLast(documentType string, year int) (int, error)
	Save(documentType string, year int, seq int) error
}

type DocumentSequenceRepository struct {
	DB *gorm.DB
}

// NewDocumentSequenceRepository creates a new instance of DocumentSequenceRepository with the provided gorm.DB connection.
func NewDocumentSequence(db *gorm.DB) *DocumentSequenceRepository {
	return &DocumentSequenceRepository{DB: db}
}

func (r *DocumentSequenceRepository) GetLast(documentType string, year int) (int, error) {
	var seq int
	err := r.DB.
		Model(&models.DocumentSequence{}).
		Where("type = ? AND year = ?", documentType, year).
		Select("COALESCE(MAX(last_seq), 0)").
		Scan(&seq).Error

	return seq, err
}

func (r *DocumentSequenceRepository) Save(documentType string, year int, seq int) error {
	return r.DB.
		Model(&models.DocumentSequence{}).
		Where("type = ? AND year = ?", documentType, year).
		Assign(models.DocumentSequence{
			Type:    documentType,
			Year:    year,
			LastSeq: seq,
		}).
		FirstOrCreate(&models.DocumentSequence{}).Error
}
