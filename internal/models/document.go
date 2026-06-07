package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	DocumentStatusValid   = "valid"
	DocumentStatusInvalid = "invalid"
)

type Document struct {
	ID           uint       `gorm:"primaryKey" json:"id"`
	DocumentID   string     `gorm:"size:255;not null;uniqueIndex" json:"document_id"`
	OwnerName    string     `gorm:"size:255;not null" json:"owner_name"`
	Type         string     `gorm:"size:255;not null" json:"type"`
	Issuer       string     `gorm:"size:255;not null" json:"issuer"`
	Hash         string     `gorm:"size:255;not null" json:"hash"`
	BlockchainTx string     `gorm:"size:255;not null" json:"blockchain_tx"`
	QrCode       string     `gorm:"size:255;not null" json:"qr_code"`
	Status       string     `gorm:"type:enum('valid','invalid');default:'valid';index" json:"status"`
	CreatedAt    time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    *time.Time `gorm:"autoUpdateTime" json:"updated_at,omitempty"`
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
