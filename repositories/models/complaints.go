package models

import (
	"capstone/entities"
	"time"
)

// Complaint struct
type Complaint struct {
	ID          int              `gorm:"primaryKey;autoIncrement"`
	UserID      int              `gorm:"not null"`
	User        User             `gorm:"foreignKey:UserID"`
	CategoryID  int              `gorm:"not null"`
	Category    Category         `gorm:"foreignKey:CategoryID"`
	Status      string           `gorm:"type:enum('pending', 'resolved', 'rejected');default:'pending'"`
	Description string           `gorm:"type:text;not null"`
	Photos      []ComplaintPhoto `gorm:"foreignKey:ComplaintID"`
	CreatedAt   time.Time        `gorm:"autoCreateTime"`
	UpdatedAt   time.Time        `gorm:"autoUpdateTime"`
}

func FromEntitiesComplaint(c entities.Complaint) Complaint {
	return Complaint{
		ID:          c.ID,
		UserID:      c.UserID,
		CategoryID:  c.CategoryID,
		Status:      c.Status,
		Description: c.Description,
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
	}
}

func (c Complaint) ToEntities() entities.Complaint {
	var photos []entities.ComplaintPhoto
	for _, photo := range c.Photos {
		photos = append(photos, photo.ToEntities())
	}

	return entities.Complaint{
		ID:          c.ID,
		UserID:      c.UserID,
		User:        c.User.ToEntities(),
		CategoryID:  c.CategoryID,
		Category:    c.Category.ToEntities(),
		Description: c.Description,
		Status:      c.Status,
		Photos:      photos,
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
	}
}
