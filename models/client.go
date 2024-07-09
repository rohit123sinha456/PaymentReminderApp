package models

import (
	"gorm.io/gorm"
	"time"
	. "github.com/rohit123sinha456/payredapp/types"
)

type Client struct {
	ID        uint           `gorm:"primaryKey"`
	Name      string         `gorm:"size:255;not null"`
	Emails    StringArray       `gorm:"type:json"`
	PhoneNumbers StringArray    `gorm:"type:json"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	IsDeleted           bool           `gorm:"default:false"`
	SiteName  string         `gorm:"size:255;not null"`
	Address   string         `gorm:"size:255"`
	LoginEmail string         `gorm:"size:255;not null"`
	Password    string         `gorm:"size:255;not null"`
}