package models

import "time"

// ImportLog struct
type ImportLog struct {
	ID           uint      `gorm:"primaryKey;autoIncrement"`
	FileName     string    `gorm:"type:varchar(255);not null"`
	ImportedBy   uint      `gorm:"not null"`
	User         User      `gorm:"foreignKey:ImportedBy"`
	SuccessCount int       `gorm:"not null"`
	FailureCount int       `gorm:"not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
}
