package models

import (
	"capstone/entities"
	"time"
)

// AIResponse struct
type AIResponse struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	UserID    int       `gorm:"not null"`
	User      User      `gorm:"foreignKey:UserID"`
	Request   string    `gorm:"type:text;not null"`
	Response  string    `gorm:"type:text;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func (r AIResponse) ToEntities() entities.AIResponse {
	return entities.AIResponse{
		ID:     int(r.ID),
		UserID: r.UserID,
		User: entities.User{
			ID:       r.User.ID,
			Name:     r.User.Name,
			Email:    r.User.Email,
			Phone:    r.User.Phone,
			PhotoURL: r.User.PhotoURL,
		},
		Request:   r.Request,
		Response:  r.Response,
		CreatedAt: r.CreatedAt,
	}
}
