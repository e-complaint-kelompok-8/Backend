package models

import "time"

// Feedback struct
type Feedback struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	UserID      uint      `gorm:"not null"`
	User        User      `gorm:"foreignKey:UserID"`
	ComplaintID uint      `gorm:"not null"`
	Complaint   Complaint `gorm:"foreignKey:ComplaintID"`
	Content     string    `gorm:"type:text;not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}
