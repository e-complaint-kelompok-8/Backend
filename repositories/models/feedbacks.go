package models

import (
	"capstone/entities"
	"time"
)

// Feedback struct
type Feedback struct {
	ID          int       `gorm:"primaryKey;autoIncrement"`
	AdminID     int       `gorm:"not null"`
	Admin       Admin     `gorm:"foreignKey:AdminID"`
	UserID      int       `gorm:"not null"`
	User        User      `gorm:"foreignKey:UserID"`
	ComplaintID int       `gorm:"not null"`
	Complaint   Complaint `gorm:"foreignKey:ComplaintID"`
	Content     string    `gorm:"type:text;not null"`
	Response    string    `gorm:"type:text"` // Tambahkan kolom ini
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

// FromEntitiesFeedback converts an entity feedback to a model feedback
func FromEntitiesFeedback(feedback entities.Feedback) Feedback {
	return Feedback{
		ID:          feedback.ID,
		AdminID:     feedback.Admin.ID,
		UserID:      feedback.User.ID,
		ComplaintID: feedback.Complaint.ID,
		Content:     feedback.Content,
		CreatedAt:   feedback.CreatedAt,
	}
}

// ToEntities converts a model feedback to an entity feedback
func (feedback Feedback) ToEntities() entities.Feedback {
	return entities.Feedback{
		ID:      feedback.ID,
		AdminID: feedback.AdminID,
		Admin: entities.Admin{
			ID:        feedback.Admin.ID,
			Email:     feedback.Admin.Email,
			Role:      feedback.Admin.Role,
			CreatedAt: feedback.Admin.CreatedAt,
			UpdatedAt: feedback.Admin.UpdatedAt,
		},
		UserID: feedback.UserID,
		User: entities.User{
			ID:        feedback.User.ID,
			Name:      feedback.User.Name,
			Email:     feedback.User.Email,
			Phone:     feedback.User.Phone,
			CreatedAt: feedback.User.CreatedAt,
			UpdatedAt: feedback.User.UpdatedAt,
		},
		ComplaintID: feedback.ComplaintID,
		Complaint: entities.Complaint{
			ID:         feedback.Complaint.ID,
			CategoryID: feedback.Complaint.CategoryID,
			Category: entities.Category{
				ID:          feedback.Complaint.Category.ID,
				Name:        feedback.Complaint.Category.Name,
				Description: feedback.Complaint.Category.Description,
				CreatedAt:   feedback.Complaint.Category.CreatedAt,
				UpdatedAt:   feedback.Complaint.Category.UpdatedAt,
			},
			ComplaintNumber: feedback.Complaint.ComplaintNumber,
			Title:           feedback.Complaint.Title,
			Location:        feedback.Complaint.Location,
			Status:          feedback.Complaint.Status,
			Description:     feedback.Complaint.Description,
			Photos:          ToEntityPhotos(feedback.Complaint.Photos),
			CreatedAt:       feedback.Complaint.CreatedAt,
			UpdatedAt:       feedback.Complaint.UpdatedAt,
		},
		Content:   feedback.Content,
		Response:  feedback.Response,
		CreatedAt: feedback.CreatedAt,
	}
}

// Helper function to convert ComplaintPhoto models to entities
func ToEntityPhotos(photos []ComplaintPhoto) []entities.ComplaintPhoto {
	var entityPhotos []entities.ComplaintPhoto
	for _, photo := range photos {
		entityPhotos = append(entityPhotos, photo.ToEntities())
	}
	return entityPhotos
}
