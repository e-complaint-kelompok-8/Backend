package models

import (
	"capstone/entities"
	"time"
)

type AISuggestion struct {
	ID          int       `gorm:"primaryKey;autoIncrement"`
	AdminID     int       `gorm:"not null"`
	Admin       Admin     `gorm:"foreignKey:AdminID"`
	ComplaintID int       `gorm:"not null"`
	Complaint   Complaint `gorm:"foreignKey:ComplaintID"`
	Request     string    `gorm:"type:text"`
	Response    string    `gorm:"type:text;not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}

func (r AISuggestion) ToEntities() entities.AISuggestion {
	return entities.AISuggestion{
		ID:          r.ID,
		AdminID:     r.AdminID,
		Admin:       r.Admin.ToEntities(),
		ComplaintID: r.ComplaintID,
		Complaint:   r.Complaint.ToEntities(),
		Request:     r.Request,
		Response:    r.Response,
		CreatedAt:   r.CreatedAt,
	}
}
