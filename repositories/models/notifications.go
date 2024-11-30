package models

import "time"

// Notification struct
type Notification struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	UserID    uint      `gorm:"not null"`
	User      User      `gorm:"foreignKey:UserID"`
	Content   string    `gorm:"type:text;not null"`
	IsRead    bool      `gorm:"default:false"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
