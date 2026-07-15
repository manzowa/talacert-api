package models

import (
	"gorm.io/gorm"
)

type DocumentSequence struct {
	gorm.Model

	Type    string `gorm:"primaryKey" json:"type"`
	Year    int    `gorm:"primaryKey" json:"year"`
	LastSeq int    `gorm:"not null" json:"last_seq"`
}

func (u *DocumentSequence) TableName() string {
	return "document_sequences"
}
