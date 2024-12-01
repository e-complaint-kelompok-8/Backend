package models

import (
	"capstone/entities"
	"time"
)

// User struct
type User struct {
	ID        int    `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"type:varchar(255);not null"`
	Phone     string `gorm:"type:varchar(255);not null"`
	Email     string `gorm:"type:varchar(255);unique;not null"`
	Password  string `gorm:"type:varchar(255);not null"`
	Verified  bool   `gorm:"default:false"`
	OTP       string `gorm:"type:varchar(6)"`
	OTPExpiry time.Time
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func FromEntitiesUser(user entities.User) User {
	return User{
		ID:        user.ID,
		Name:      user.Name,
		Phone:     user.Phone,
		Email:     user.Email,
		Password:  user.Password,
		Verified:  user.Verified,
		OTP:       user.OTP,
		OTPExpiry: user.OTPExpiry,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (user User) ToEntities() entities.User {
	return entities.User{
		ID:        user.ID,
		Name:      user.Name,
		Phone:     user.Phone,
		Email:     user.Email,
		Password:  user.Password,
		Verified:  user.Verified,
		OTP:       user.OTP,
		OTPExpiry: user.OTPExpiry,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
