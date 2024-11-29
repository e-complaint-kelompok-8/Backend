package entities

import "time"

type Commentar struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	NewsID    int       `json:"news_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}