package entities

import "time"

type ChatMessage struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	AdminID   int       `json:"admin_id"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}