package entities

import "time"

type Complaint struct {
	ID              int              `json:"id"`
	UserID          int              `json:"user_id"`
	User            User             `json:"user"`
	CategoryID      int              `json:"category_id"`
	Category        Category         `json:"category"`
	ComplaintNumber string           `json:"complaint_number"`
	Title           string           `json:"title"`
	Location        string           `json:"location"`
	Status          string           `json:"status"`
	Description     string           `json:"description"`
	Photos          []ComplaintPhoto `json:"photos"`
	CreatedAt       time.Time        `json:"create_at"`
	UpdatedAt       time.Time        `json:"update_at"`
}
