package models

import (
	"capstone/entities"
	"time"
)

type ComplaintPhoto struct {
	ID          int       `gorm:"primaryKey;autoIncrement"`
	ComplaintID int       `gorm:"not null"`
	PhotoURL    string    `gorm:"type:varchar(255);not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}

func FromEntitiesComplaintPhoto(photo entities.ComplaintPhoto) ComplaintPhoto {
	return ComplaintPhoto{
		ID:          photo.ID,
		ComplaintID: photo.ComplaintID,
		PhotoURL:    photo.PhotoURL,
		CreatedAt:   time.Now(),
	}
}

func (cp ComplaintPhoto) ToEntities() entities.ComplaintPhoto {
	return entities.ComplaintPhoto{
		ID:          cp.ID,
		ComplaintID: cp.ComplaintID,
		PhotoURL:    cp.PhotoURL,
		CreatedAt:   cp.CreatedAt,
	}
}
