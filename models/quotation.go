package models

import (
	"gorm.io/gorm"
	"time"
)

type Quotation struct {
	ID                 uint           `gorm:"primaryKey"`
	QuotationNumber    string         `gorm:"size:255;not null"`
	QuotationDate      time.Time      `gorm:"not null"`
	SentDate      time.Time      `gorm:"not null"`
	QuotationTotalValue float64       `gorm:"not null"`
	QuotationReason    string         `gorm:"size:255"`
	ClientID           uint           `gorm:"not null"`
	Client             Client         `gorm:"foreignKey:ClientID"`
	QuotationValidTill  time.Time      
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt `gorm:"index"`
	IsDeleted           bool           `gorm:"default:false"`
	IsApproved         bool           `gorm:"default:false"`
	QuotationScannedLink string       `gorm:"size:255"`
}
