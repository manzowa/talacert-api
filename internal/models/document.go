package models

import (
	"talacert-api/internal/constants"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	DocumentStatusValid   = "valid"
	DocumentStatusInvalid = "invalid"
)

type Document struct {
	gorm.Model

	DocumentID   string                   `gorm:"type:varchar(191);not null;uniqueIndex" json:"document_id"`
	OwnerName    string                   `gorm:"type:varchar(191);not null" json:"owner_name"`
	Type         string                   `gorm:"type:varchar(100);not null" json:"type"`
	Issuer       string                   `gorm:"type:varchar(191);not null" json:"issuer"`
	Hash         string                   `gorm:"type:char(64);not null;uniqueIndex" json:"hash"`
	BlockchainTx string                   `gorm:"type:text;not null" json:"blockchain_tx"`
	QrCode       string                   `gorm:"type:text;not null" json:"qr_code"`
	Status       constants.DocumentStatus `gorm:"type:varchar(20);not null;default:'valid'" json:"status"`
}

func NewDocument(ownerName, docType, issuer, hash string) *Document {
	return &Document{
		DocumentID: uuid.New().String(),
		OwnerName:  ownerName,
		Type:       docType,
		Issuer:     issuer,
		Hash:       hash,
		Status:     DocumentStatusValid,
	}
}

func (d *Document) BeforeCreate(tx *gorm.DB) (err error) {
	if d.DocumentID == "" {
		d.DocumentID = uuid.New().String()
	}
	return nil
}

func (Document) TableName() string {
	return "documents"
}
