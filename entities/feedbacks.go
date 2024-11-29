package entities

import "time"

type Feedback struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	ComplaintID int       `json:"complaint_id"`
	Content     string    `json:"content"`
	CreatedAt   time.Time `json:"created_at"`
}