package models

import "time"

type Document struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	DocumentID   string    `json:"document_id"`
	OwnerName    string    `json:"owner_name"`
	Type         string    `json:"type"`
	Issuer       string    `json:"issuer"`
	Hash         string    `json:"hash"`
	BlockchainTx string    `json:"blockchain_tx"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
