package dto

import "time"

type DocumentResponse struct {
	DocumentID   string     `json:"document_id"`
	OwnerName    string     `json:"owner_name,omitempty"`
	Type         string     `json:"type,omitempty"`
	Issuer       string     `json:"issuer,omitempty"`
	Hash         string     `json:"hash,omitempty"`
	BlockchainTx string     `json:"blockchain_tx,omitempty"`
	QrCode       string     `json:"qr_code,omitempty"`
	Status       string     `json:"status,omitempty"`
	CreatedAt    time.Time  `json:"created_at,omitempty"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
}
