package entities

import "time"

type User struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	Email            string    `json:"email"`
	Password         string    `json:"password"`
	NoTelp           string    `json:"no_telp"`
	Role             string    `json:"role"`
	UrlPhotoProfile  string    `json:"url_photo_profile"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}