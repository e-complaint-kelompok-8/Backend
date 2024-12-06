package entities

import "time"

type Feedback struct {
	ID          int       `json:"id"`
	AdminID     int       `json:"admin_id"`
	Admin       Admin     `json:"admin"`
	UserID      int       `json:"user_id"`
	User        User      `json:"user"`
	ComplaintID int       `json:"complaint_id"`
	Complaint   Complaint `json:"complaint"`
	Content     string    `json:"content"`
	CreatedAt   time.Time `json:"created_at"`
}
