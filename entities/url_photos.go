package entities

import "time"

type UrlPhoto struct {
	ID        int       `json:"id"`
	UrlPhoto  string    `json:"url_photo"`
	CreatedAt time.Time `json:"created_at"`
}