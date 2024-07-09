package models

import (
	"gorm.io/gorm"
	"time"
)

type ServiceReport struct {
	ID                  uint           `gorm:"primaryKey"`
	ServiceReportNumber string         `gorm:"size:255;not null"`
	ServiceReportDate   time.Time      `gorm:"not null"`
	SentDate            time.Time
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt `gorm:"index"`
	IsDeleted           bool           `gorm:"default:false"`
	ServiceReportLink   string         `gorm:"size:255"`
	ClientID            uint           `gorm:"not null"`
	Client              Client         `gorm:"foreignKey:ClientID"`
	ServiceReportNotes  string         `gorm:"type:text"`
	AttentionRequired   bool           `gorm:"size:255"`
	SystemStatus        string         `gorm:"size:255"`
}
