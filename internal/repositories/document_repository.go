package repositories

import (
	"context"
	"errors"

	"talacert-api/internal/logger"
	"talacert-api/internal/models"

	"gorm.io/gorm"
)

type DocumentRepositoryInterface interface {
	Create(cxt context.Context, document *models.Document) error
	Update(cxt context.Context, document *models.Document) error
	Delete(cxt context.Context, documentID string) error

	FindById(cxt context.Context, documentID string) (*models.Document, error)
	FindByHash(cxt context.Context, hash string) (*models.Document, error)
	FindAll(cxt context.Context) ([]models.Document, error)
}

type DocumentRepository struct {
	DB *gorm.DB
}

// NewDocumentRepository creates a new instance of DocumentRepository with the provided gorm.DB connection.
func NewDocumentRepository(db *gorm.DB) *DocumentRepository {
	return &DocumentRepository{DB: db}
}

func (r *DocumentRepository) Create(
	cxt context.Context,
	document *models.Document,
) error {

	if err := r.DB.WithContext(cxt).Create(document).Error; err != nil {
		logger.AccessLogger.Error("failed to create document", "error", err)
		return err
	}

	return nil
}

func (r *DocumentRepository) Update(
	cxt context.Context,
	document *models.Document,
) error {

	if err := r.DB.WithContext(cxt).Save(document).Error; err != nil {
		logger.AccessLogger.Error("failed to update document", "error", err)
		return err
	}

	return nil
}

func (r *DocumentRepository) Delete(
	cxt context.Context,
	documentID string,
) error {

	if err := r.DB.WithContext(cxt).
		Where("document_id = ?", documentID).
		Delete(&models.User{}).Error; err != nil {

		logger.AccessLogger.Error("failed to delete user", "error", err)
		return err
	}

	return nil
}

func (r *DocumentRepository) FindById(
	cxt context.Context,
	documentID string,
) (*models.Document, error) {
	var document models.Document

	err := r.DB.WithContext(cxt).
		Where("document_id = ?", documentID).
		First(&document).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &document, err
}

func (r *DocumentRepository) FindByHash(
	cxt context.Context,
	hash string,
) (*models.Document, error) {
	var document models.Document

	err := r.DB.WithContext(cxt).
		Where("hash = ?", hash).
		First(&document).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &document, err
}

func (r *DocumentRepository) FindAll(cxt context.Context) ([]models.Document, error) {
	var documents []models.Document

	err := r.DB.WithContext(cxt).
		Find(&documents).
		Limit(20).
		Error

	return documents, err
}
