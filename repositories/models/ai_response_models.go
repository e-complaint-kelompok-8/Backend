package models

import "time"

// AIResponse struct
type AIResponse struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	ComplaintID uint      `gorm:"not null"`
	Complaint   Complaint `gorm:"foreignKey:ComplaintID"`
	Response    string    `gorm:"type:text;not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}
