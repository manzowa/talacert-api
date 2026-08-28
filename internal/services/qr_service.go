package services

import (
	"fmt"

	"github.com/skip2/go-qrcode"
)

type QRService struct {
	BaseURL string
}

func NewQRService(baseURL string) *QRService {
	return &QRService{
		BaseURL: baseURL,
	}
}

func (s *QRService) GeneratePath(documentID string) string {
	return fmt.Sprintf(
		"/verify/%s",
		documentID,
	)
}

func (s *QRService) GenerateUrl(documentID string) string {
	qrCodePath := s.GeneratePath(documentID)
	return s.BaseURL + qrCodePath
}

func (s *QRService) GeneratePng(documentID string) ([]byte, error) {
	url := s.GenerateUrl(documentID)

	qrCode, err := qrcode.Encode(url, qrcode.Medium, 256)
	if err != nil {
		return nil, err
	}
	return qrCode, nil
}
