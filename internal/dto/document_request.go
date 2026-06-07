package dto

type DocumentRequest struct {
	OwnerName string `json:"owner_name" binding:"required"`
	Type      string `json:"type" binding:"required"`
	Issuer    string `json:"issuer" binding:"required"`
}
