package entities

import "time"

type Comment struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	User      User      `json:"user"`
	NewsID    int       `json:"news_id"`
	News      News      `json:"news"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
