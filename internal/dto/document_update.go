package dto

// L'utilisation de pointeurs permet de distinguer :
// - une valeur non fournie (nil) : le champ ne sera pas mis à jour
// - une valeur fournie (même vide) : le champ sera mis à jour avec cette valeur
type DocumentUpdateDTO struct {
	OwnerName    *string `json:"owner_name,omitempty"`
	Type         *string `json:"type,omitempty"`
	Issuer       *string `json:"issuer,omitempty"`
	Status       *string `json:"status,omitempty"`
	BlockchainTx *string `json:"blockchain_tx,omitempty"`
	QrCode       *string `json:"qr_code,omitempty"`
	UpdatedAt    *string `json:"updated_at,omitempty"`
}
