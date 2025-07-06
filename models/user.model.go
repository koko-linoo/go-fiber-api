package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	FullName  string         `json:"fullName" validate:"required"`
	Username  string         `gorm:"unique" json:"username" validate:"required"`
	Email     string         `gorm:"unique" json:"email" validate:"required"`
	Phone     *string        `gorm:"unique" json:"phone"`
	Password  string         `json:"password" validate:"required"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
