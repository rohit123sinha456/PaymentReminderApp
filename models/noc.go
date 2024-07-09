package models

import (
	"gorm.io/gorm"
	"time"
)

type NOC struct {
    ID            uint           `gorm:"primaryKey"`
    CreatedAt     time.Time
    UpdatedAt     time.Time
    DeletedAt     gorm.DeletedAt `gorm:"index"`
    IsDeleted     bool           `gorm:"default:false"`
    StatusNote    string         `gorm:"type:text"`
    DocumentLink  string         `gorm:"size:255"`
    ClientID      uint           `gorm:"not null"`
    Client        Client         `gorm:"foreignKey:ClientID"`
}
