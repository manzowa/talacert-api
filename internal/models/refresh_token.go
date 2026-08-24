package models

import (
	"time"

	"gorm.io/gorm"
)

type RefreshToken struct {
	gorm.Model

	UserID    uint      `gorm:"not null;index" json:"user_id"`
	Token     string    `gorm:"type:char(64);not null;uniqueIndex"`
	ExpiresAt time.Time `gorm:"not null;" json:"expires_at"`
	Revoked   bool      `gorm:"not null;default:false" json:"revoked"`

	User User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

func (RefreshToken) TableName() string {
	return "refresh_tokens"
}
