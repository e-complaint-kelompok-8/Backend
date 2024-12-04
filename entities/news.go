package entities

import "time"

type News struct {
	ID        int       `json:"id"`
	Admin     Admin     `json:"admin"`
	Category  Category  `json:"category"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	PhotoURL  string    `json:"photo_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
