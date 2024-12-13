package entities

import "time"

type AISuggestion struct {
	ID          int       `json:"id"`
	AdminID     int       `json:"admin_id"`
	Admin       Admin     `json:"admin"`
	ComplaintID int       `json:"complaint_id"`
	Complaint   Complaint `json:"complaint"`
	Request     string    `json:"request"`
	Response    string    `json:"response"`
	CreatedAt   time.Time `json:"create_at"`
}
