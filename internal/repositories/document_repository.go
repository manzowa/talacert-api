package repositories

import (
	"talacert-api/internal/models"
	"time"

	"gorm.io/gorm"
)

type DocumentRepositoryInterface interface {
	Create(document *models.Document) error
	FindAll() ([]models.Document, error)
	FindById(documentID string) (*models.Document, error)
	FindByHash(hash string) (*models.Document, error)
	Update(documentID string, data map[string]any) (*models.Document, error)
	Delete(documentID string) error
	Patch(documentID string, data map[string]any) error
}

type DocumentRepository struct {
	DB *gorm.DB
}

// NewDocumentRepository creates a new instance of DocumentRepository with the provided gorm.DB connection.
func New(db *gorm.DB) *DocumentRepository {
	return &DocumentRepository{DB: db}
}

func (r *DocumentRepository) FindAll() ([]models.Document, error) {
	var documents []models.Document

	err := r.DB.
		Select("id, document_id, owner_name, type, issuer, hash, status, created_at").
		Find(&documents).Limit(20).Error

	return documents, err
}

func (r *DocumentRepository) Create(document *models.Document) error {
	return r.DB.Create(document).Error
}

func (r *DocumentRepository) FindByDocumentId(documentID string) (*models.Document, error) {
	var document models.Document

	err := r.DB.
		Select("id, document_id, owner_name, type, issuer, hash, status, created_at").
		Where("document_id = ?", documentID).
		First(&document).Error

	return &document, err
}

func (r *DocumentRepository) FindByHash(hash string) (*models.Document, error) {
	var document models.Document
	err := r.DB.
		Select("id, document_id, owner_name, type, issuer, hash, status, created_at").
		Where("hash = ?", hash).
		First(&document).Error

	return &document, err
}

func (r *DocumentRepository) Update(documentID string, data map[string]any) (*models.Document, error) {

	// 🔥 FORCER updated_at (IMPORTANT avec map)
	data["updated_at"] = time.Now()

	if err := r.DB.
		Model(&models.Document{}).
		Where("document_id = ?", documentID).
		Updates(data).Error; err != nil {
		return nil, err
	}

	// 🔥 re-fetch propre (source of truth DB)
	var document models.Document
	if err := r.DB.
		Where("document_id = ?", documentID).
		First(&document).Error; err != nil {
		return nil, err
	}

	return &document, nil
}

func (r *DocumentRepository) Delete(documentID string) error {
	result := r.DB.Where("document_id = ?", documentID).Delete(&models.Document{})

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return result.Error
}
