package mapper

import (
	"talacert-api/internal/dto"
	"talacert-api/internal/models"
)

func ToModelDocument(docRequest *dto.DocumentRequest) models.Document {
	return models.Document{
		OwnerName: docRequest.OwnerName,
		Type:      docRequest.Type,
		Issuer:    docRequest.Issuer,
		Status:    models.DocumentStatusValid,
	}
}

func ToResponseDocument(doc *models.Document) dto.DocumentResponse {
	return dto.DocumentResponse{
		DocumentID:   doc.DocumentID,
		OwnerName:    doc.OwnerName,
		Type:         doc.Type,
		Issuer:       doc.Issuer,
		Hash:         doc.Hash,
		BlockchainTx: doc.BlockchainTx,
		QrCode:       doc.QrCode,
		Status:       doc.Status,
		CreatedAt:    doc.CreatedAt,
		UpdatedAt:    doc.UpdatedAt,
	}
}

func ToResponseDocuments(documents []models.Document) []dto.DocumentResponse {
	if len(documents) == 0 {
		return []dto.DocumentResponse{}
	}

	responses := make([]dto.DocumentResponse, len(documents))

	for i, doc := range documents {
		responses[i] = ToResponseDocument(&doc)
	}

	return responses
}
