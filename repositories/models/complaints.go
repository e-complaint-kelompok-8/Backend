package models

import (
	"capstone/entities"
	"time"
)

// Complaint struct
type Complaint struct {
	ID              int              `gorm:"primaryKey;autoIncrement"`
	AdminID         *int             `gorm:"admin_id"` // Gunakan pointer untuk mendukung nilai NULL
	Admin           Admin            `gorm:"foreignKey:AdminID"`
	UserID          int              `gorm:"not null"`
	User            User             `gorm:"foreignKey:UserID"`
	CategoryID      int              `gorm:"not null"`
	Category        Category         `gorm:"foreignKey:CategoryID"`
	ComplaintNumber string           `gorm:"type:varchar(255);unique"`
	Title           string           `gorm:"type:varchar(255);not null"`
	Location        string           `gorm:"type:varchar(255);not null"`
	Status          string           `gorm:"type:enum('proses', 'tanggapi', 'batal', 'selesai');default:'proses'"`
	Description     string           `gorm:"type:text;not null"`
	Photos          []ComplaintPhoto `gorm:"foreignKey:ComplaintID;constraint:OnDelete:CASCADE"` // Tambahkan OnDelete:CASCADE
	Reason          string           `gorm:"type:text"`
	CreatedAt       time.Time        `gorm:"autoCreateTime"`
	UpdatedAt       time.Time        `gorm:"autoUpdateTime"`
}

func FromEntitiesComplaint(c entities.Complaint) Complaint {
	return Complaint{
		ID:              c.ID,
		UserID:          c.UserID,
		CategoryID:      c.CategoryID,
		ComplaintNumber: c.ComplaintNumber,
		Title:           c.Title,
		Location:        c.Location,
		Status:          c.Status,
		Description:     c.Description,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
}

func (c Complaint) ToEntities() entities.Complaint {
	var photos []entities.ComplaintPhoto
	for _, photo := range c.Photos {
		photos = append(photos, photo.ToEntities())
	}

	return entities.Complaint{
		ID:              c.ID,
		AdminID:         c.AdminID,
		Admin:           c.Admin.ToEntities(),
		UserID:          c.UserID,
		User:            c.User.ToEntities(),
		CategoryID:      c.CategoryID,
		Category:        c.Category.ToEntities(),
		ComplaintNumber: c.ComplaintNumber,
		Title:           c.Title,
		Location:        c.Location,
		Description:     c.Description,
		Status:          c.Status,
		Photos:          photos,
		CreatedAt:       c.CreatedAt,
		UpdatedAt:       c.UpdatedAt,
	}
}

func (c Complaint) ToEntitiesReason() entities.Complaint {
	var photos []entities.ComplaintPhoto
	for _, photo := range c.Photos {
		photos = append(photos, photo.ToEntities())
	}

	return entities.Complaint{
		ID:              c.ID,
		AdminID:         c.AdminID,
		Admin:           c.Admin.ToEntities(),
		UserID:          c.UserID,
		User:            c.User.ToEntities(),
		CategoryID:      c.CategoryID,
		Category:        c.Category.ToEntities(),
		ComplaintNumber: c.ComplaintNumber,
		Title:           c.Title,
		Location:        c.Location,
		Description:     c.Description,
		Status:          c.Status,
		Photos:          photos,
		Reason:          c.Reason,
		CreatedAt:       c.CreatedAt,
		UpdatedAt:       c.UpdatedAt,
	}
}
