package models

import "time"

// AIResponse struct
type AIResponse struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	UserID    int       `gorm:"not null"`
	User      User      `gorm:"foreignKey:UserID"`
	Request   string    `gorm:"type:text;not null"`
	Response  string    `gorm:"type:text;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
