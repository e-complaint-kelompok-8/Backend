package entities

import "time"

type User struct {
	ID         int         `json:"id"`
	Name       string      `json:"name"`
	Email      string      `json:"email"`
	Password   string      `json:"password"`
	Phone      string      `json:"phone_number"`
	PhotoURL   string      `json:"photo"`
	Verified   bool        `gorm:"default:false"`
	OTP        string      `gorm:"type:varchar(6)"` // Menyimpan kode OTP
	OTPExpiry  time.Time   // Menyimpan waktu kedaluwarsa OTP
	Token      string      `json:"token"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
	Complaints []Complaint `json:"complaints"`
}
