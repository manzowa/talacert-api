package models

import (
	"talacert-api/internal/constants"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Username string         `gorm:"type:varchar(191);not null;uniqueIndex" json:"username"`
	Email    string         `gorm:"type:varchar(191);not null;uniqueIndex" json:"email"`
	Password string         `gorm:"type:varchar(255);not null" json:"-"`
	Role     constants.Role `gorm:"type:varchar(20);not null;default:'USER'" json:"role"`
}

func (u *User) TableName() string {
	return "users"
}
