package entities

import "time"

type AIResponse struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	User      User      `json:"user"`
	Request   string    `json:"request"`
	Response  string    `json:"response"`
	CreatedAt time.Time `json:"created_at"`
}
