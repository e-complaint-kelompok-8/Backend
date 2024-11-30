package entities

import "time"

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Phone     string    `json:"no_telp"`
	Verified  bool      `gorm:"default:false"`
	OTP       string    `gorm:"type:varchar(6)"` // Menyimpan kode OTP
	OTPExpiry time.Time // Menyimpan waktu kedaluwarsa OTP
	Role      string    `json:"role"`
	Token     string
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
