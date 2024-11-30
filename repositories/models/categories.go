package models

import (
	"capstone/entities"
	"time"
)

// Category struct
type Category struct {
	ID          int       `gorm:"primaryKey;autoIncrement"`
	Name        string    `gorm:"type:varchar(255);not null"`
	Description string    `gorm:"type:text"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

func FromEntitiesCategory(c entities.Category) Category {
	return Category{
		ID:          c.ID,
		Name:        c.Name,
		Description: c.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (c Category) ToEntities() entities.Category {
	return entities.Category{
		ID:          c.ID,
		Name:        c.Name,
		Description: c.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
