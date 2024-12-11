package request

import (
	"capstone/entities"
	"time"
)

// CategoryRequest is the struct for handling incoming request data
type CategoryRequest struct {
	Name        string `json:"name" validate:"required,min=3,max=255"`
	Description string `json:"description" validate:"max=500"`
}

// ToEntity converts CategoryRequest to Category entity
func (req *CategoryRequest) ToEntity() entities.Category {
	return entities.Category{
		Name:        req.Name,
		Description: req.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

// UpdateCategoryRequest is the struct for handling update request data
type UpdateCategoryRequest struct {
	ID          int    `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required,min=3,max=255"`
	Description string `json:"description" validate:"max=500"`
}

// ToEntity converts UpdateCategoryRequest to Category entity
func (req *UpdateCategoryRequest) ToEntity() entities.Category {
	return entities.Category{
		ID:          req.ID,
		Name:        req.Name,
		Description: req.Description,
		UpdatedAt:   time.Now(),
	}
}