package models

type DocumentSequence struct {
	Type    string `gorm:"primaryKey" json:"type"`
	Year    int    `gorm:"primaryKey" json:"year"`
	LastSeq int    `gorm:"not null" json:"last_seq"`
}
