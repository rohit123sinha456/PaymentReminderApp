package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID          uint           `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	IsDeleted   bool           `gorm:"default:false"`
	Name        string         `gorm:"size:255;not null"`
	Email       string         `gorm:"size:255;unique;not null"`
	PhoneNumber string         `gorm:"size:20;not null"`
	Password    string         `gorm:"size:255;not null"`
}

