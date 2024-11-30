package entities

import "time"

type ComplaintPhoto struct {
	ID          int       `json:"id"`
	ComplaintID int       `json:"complaint_id"`
	PhotoURL    string    `json:"photo_url"`
	CreatedAt   time.Time `json:"create_at"`
}
