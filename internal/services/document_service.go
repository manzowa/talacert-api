package services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"talacert-api/internal/constants"
	"talacert-api/internal/dto"
	"talacert-api/internal/logger"
	"talacert-api/internal/models"
	"talacert-api/internal/repositories"
	"talacert-api/internal/utils"
)

var (
	ErrDocumentNotFound = errors.New("user not found")
)

type DocumentService struct {
	DocumentRepository         repositories.DocumentRepositoryInterface
	DocumentSequenceRepository repositories.DocumentSequenceRepositoryInterface
	Hash                       *utils.HashGenerator
	QRCode                     *QRService
}

func NewDocument(
	repository repositories.DocumentRepositoryInterface,
	sequenceRepository repositories.DocumentSequenceRepositoryInterface,
	hash *utils.HashGenerator,
	qrCode *QRService,
) *DocumentService {
	return &DocumentService{
		DocumentRepository:         repository,
		DocumentSequenceRepository: sequenceRepository,
		Hash:                       hash,
		QRCode:                     qrCode,
	}
}
func (s *DocumentService) GetAll(
	ctx context.Context,
) ([]dto.DocumentResponse, error) {
	documents, err := s.DocumentRepository.FindAll(ctx)

	if err != nil {
		return nil, err
	}

	response := make([]dto.DocumentResponse, 0, len(documents))

	for _, document := range documents {
		response = append(response, s.toDocumentResponse(&document))
	}
	return response, nil
}

func (s *DocumentService) Create(
	ctx context.Context,
	req *dto.CreateDocumentRequest,
) error {

	year := time.Now().Year()

	// 1. sequence
	lastSeq, err := s.DocumentSequenceRepository.GetLast(ctx, string(req.Type), year)
	if err != nil {
		return err
	}

	// 2. Générer la nouvelle séquence
	seq := lastSeq + 1

	// 3. Générer le Document ID
	var documentID = s.Hash.GenerateDocumentID(
		string(req.Type),
		year,
		seq,
	)

	// 4. Génére le hash
	content := fmt.Sprintf(
		"%s|%s|%s|%s",
		documentID,
		req.OwnerName,
		req.Type,
		req.Issuer,
	)

	documentHash := s.Hash.Generate(content)

	// 5. Générer l'URL de vérification / QR Code
	if s.QRCode == nil {
		logger.AccessLogger.Error("QRService is nil")
		return nil
	}
	qrCode := s.QRCode.GeneratePath(documentID)

	// 6. save sequence
	if err := s.DocumentSequenceRepository.Save(ctx, string(req.Type), year, seq); err != nil {
		return err
	}

	document := &models.Document{
		DocumentID: documentID,
		OwnerName:  req.OwnerName,
		Type:       req.Type,
		Issuer:     req.Issuer,
		Hash:       documentHash,
		QRCode:     qrCode,
		Status:     constants.DocumentStatus(constants.DocumentValid),
	}

	log.Printf("✅ document = %+v", document)

	return s.DocumentRepository.Create(ctx, document)

}

func (s *DocumentService) Update(
	ctx context.Context,
	documentID string,
	req dto.UpdateDocumentRequest,
) error {

	document, err := s.DocumentRepository.FindById(ctx, documentID)

	if err != nil {
		return err
	}
	if document == nil {
		return ErrDocumentNotFound
	}

	if req.OwnerName != "" {
		document.OwnerName = req.OwnerName
	}

	if req.Type != "" {
		document.Type = req.Type
	}

	if req.Issuer != "" {
		document.Issuer = req.Issuer
	}

	if req.Status != "" {
		document.Status = constants.DocumentStatus(req.Status)
	}

	return s.DocumentRepository.Update(ctx, document)
}

func (s *DocumentService) Delete(
	ctx context.Context,
	documentID string,
) error {
	document, err := s.DocumentRepository.FindById(ctx, documentID)

	if err != nil {
		return err
	}

	if document == nil {
		return ErrDocumentNotFound
	}

	return s.DocumentRepository.Delete(ctx, documentID)
}

func (s *DocumentService) GetByDocumentID(
	ctx context.Context,
	documentID string,
) (*dto.DocumentResponse, error) {

	document, err := s.DocumentRepository.FindById(ctx, documentID)

	if err != nil {
		return nil, err
	}

	if document == nil {
		return nil, ErrDocumentNotFound
	}

	response := s.toDocumentResponse(document)

	return &response, nil
}

func (s *DocumentService) GetByHash(
	ctx context.Context,
	hash string,
) (*dto.DocumentResponse, error) {
	document, err := s.DocumentRepository.FindByHash(ctx, hash)
	if err != nil {
		return nil, err
	}

	if document == nil {
		return nil, ErrDocumentNotFound
	}

	response := s.toDocumentResponse(document)

	return &response, nil
}

func (s *DocumentService) toDocumentResponse(
	document *models.Document,
) dto.DocumentResponse {

	return dto.DocumentResponse{
		DocumentID: document.DocumentID,
		OwnerName:  document.OwnerName,
		Type:       document.Type,
		Issuer:     document.Issuer,
		Hash:       document.Hash,
		CreatedAt:  document.CreatedAt,
	}
}
