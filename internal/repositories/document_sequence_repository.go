package repositories

import (
	"context"
	"talacert-api/internal/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DocumentSequenceRepositoryInterface interface {
	GetLast(ctx context.Context, documentType string, year int) (int, error)
	Save(ctx context.Context, documentType string, year int, seq int) error
}

type DocumentSequenceRepository struct {
	DB *gorm.DB
}

// NewDocumentSequenceRepository creates a new instance of DocumentSequenceRepository with the provided gorm.DB connection.
func NewDocumentSequenceRepository(db *gorm.DB) *DocumentSequenceRepository {
	return &DocumentSequenceRepository{DB: db}
}

func (r *DocumentSequenceRepository) GetLast(ctx context.Context, documentType string, year int) (int, error) {
	var seq int
	err := r.DB.WithContext(ctx).
		Clauses(clause.Locking{
			Strength: "UPDATE",
		}).
		Model(&models.DocumentSequence{}).
		Where("type = ? AND year = ?", documentType, year).
		Select("COALESCE(MAX(last_seq), 0)").
		Scan(&seq).Error

	return seq, err
}

func (r *DocumentSequenceRepository) Save(ctx context.Context, documentType string, year int, seq int) error {
	return r.DB.WithContext(ctx).
		Model(&models.DocumentSequence{}).
		Where("type = ? AND year = ?", documentType, year).
		Assign(models.DocumentSequence{
			Type:    documentType,
			Year:    year,
			LastSeq: seq,
		}).
		FirstOrCreate(&models.DocumentSequence{}).Error
}
