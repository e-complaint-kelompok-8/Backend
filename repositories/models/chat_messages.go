package models

import "time"

// ChatMessage struct
type ChatMessage struct {
	ID         uint      `gorm:"primaryKey;autoIncrement"`
	UserID     uint      `gorm:"not null"`
	User       User      `gorm:"foreignKey:UserID"`
	AdminID    uint      `gorm:"not null"`
	Message    string    `gorm:"type:text;not null"`
	IsFromUser bool      `gorm:"not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
}
