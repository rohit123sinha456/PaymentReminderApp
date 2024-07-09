package models

import (
	"gorm.io/gorm"
	"time"
)

type Payment struct {
	ID                uint           `gorm:"primaryKey"`
	InvoiceDate       time.Time      `gorm:"not null"`
	SentDate      time.Time      `gorm:"not null"`
	InvoiceNumber     string         `gorm:"size:255;not null"`
	InvoiceAmount     float64        `gorm:"not null"`
	InvoiceFor        string         `gorm:"size:255"`
	NumberOfReminders int            `gorm:"default:0"`
	ReminderDates     []time.Time    `gorm:"type:json"`
	IsPaid            bool           `gorm:"default:false"`
	InvoiceLink       string         `gorm:"size:255"`
	ClientID          uint           `gorm:"not null"`
	Client            Client         `gorm:"foreignKey:ClientID"`
	Categories        []string       `gorm:"type:json"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
	IsDeleted           bool           `gorm:"default:false"`
}
