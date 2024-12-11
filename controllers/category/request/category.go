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
func (req *UpdateCategoryRequest) ToEntity(existing entities.Category) entities.Category {
	// Use existing data as the base, only overwrite fields from the request
	return entities.Category{
		ID:          existing.ID, // Keep original ID
		Name:        req.Name,    // Update the Name
		Description: req.Description,
		CreatedAt:   existing.CreatedAt, // Preserve original CreatedAt
		UpdatedAt:   time.Now(),         // Set new UpdatedAt
	}
}