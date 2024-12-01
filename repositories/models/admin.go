package models

import (
	"capstone/entities"
	"time"
)

// Admin struct
type Admin struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	Email     string    `gorm:"type:varchar(255);unique;not null"`
	Password  string    `gorm:"type:varchar(255);not null"`
	Role      string    `gorm:"type:enum('admin', 'superadmin');not null;default:'admin'"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// FromEntitiesAdmin maps entities.Admin to models.Admin
func FromEntitiesAdmin(admin entities.Admin) Admin {
	return Admin{
		ID:        admin.ID,
		Email:     admin.Email,
		Password:  admin.Password,
		Role:      admin.Role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// ToEntities maps models.Admin to entities.Admin
func (admin Admin) ToEntities() entities.Admin {
	return entities.Admin{
		ID:        admin.ID,
		Email:     admin.Email,
		Password:  admin.Password,
		Role:      admin.Role,
		CreatedAt: admin.CreatedAt,
		UpdatedAt: admin.UpdatedAt,
	}
}