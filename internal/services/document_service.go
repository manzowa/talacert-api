package services

import (
	"fmt"
	"talacert-api/internal/models"
	"talacert-api/internal/repositories"
	"talacert-api/internal/utils"
	"talacert-api/internal/utils/hash"
	"time"
)

type DocumentService struct {
	DocumentRepository         *repositories.DocumentRepository
	DocumentSequenceRepository *repositories.DocumentSequenceRepository
	Hash                       *hash.Hash
}

func New(repository *repositories.DocumentRepository, sequenceRepository *repositories.DocumentSequenceRepository) *DocumentService {
	return &DocumentService{
		DocumentRepository:         repository,
		DocumentSequenceRepository: sequenceRepository,
		Hash:                       &hash.Hash{},
	}
}

func (s *DocumentService) GetDocuments() ([]models.Document, error) {
	return s.DocumentRepository.FindAll()
}

func (s *DocumentService) CreateDocument(document *models.Document) error {
	year := time.Now().Year()

	// 1. sequence
	lastSeq, err := s.DocumentSequenceRepository.GetLast(string(document.Type), year)
	if err != nil {
		//return fmt.Errorf("failed to get last sequence: %w", err)
		return err
	}
	seq := lastSeq + 1
	document.DocumentID = utils.GenerateDocumentID(
		string(document.Type),
		year,
		seq,
	)

	content := fmt.Sprintf(
		"%s|%s|%s|%s",
		document.DocumentID,
		document.OwnerName,
		document.Type,
		document.Issuer,
	)
	document.Hash = s.Hash.GenerateHash(content)

	// 4. save sequence
	if err := s.DocumentSequenceRepository.Save(string(document.Type), year, seq); err != nil {
		return err
	}

	return s.DocumentRepository.Create(document)
}

func (s *DocumentService) GetByDocumentID(documentID string) (*models.Document, error) {
	return s.DocumentRepository.FindByDocumentId(documentID)
}

func (s *DocumentService) GetByDocumentHash(hash string) (*models.Document, error) {
	return s.DocumentRepository.FindByHash(hash)
}

func (s *DocumentService) UpdateByDocumentID(documentID string, document *models.Document) (*models.Document, error) {

	data := map[string]any{
		"owner_name": document.OwnerName,
		"type":       document.Type,
		"issuer":     document.Issuer,
		"status":     document.Status,
	}

	return s.DocumentRepository.Update(documentID, data)
}

func (s *DocumentService) DeleteByDocumentID(documentID string) error {
	return s.DocumentRepository.Delete(documentID)
}
