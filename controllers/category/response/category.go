package response

import "capstone/entities"

// CategoryResponse is the struct for sending category data as a response
type CategoryResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// FromEntity converts a Category entity into a CategoryResponse
func FromEntity(entity entities.Category) CategoryResponse {
	return CategoryResponse{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
		CreatedAt:   entity.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   entity.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

// FromEntities converts a slice of Category entities into a slice of CategoryResponse
func FromEntities(entities []entities.Category) []CategoryResponse {
	var responses []CategoryResponse
	for _, entity := range entities {
		responses = append(responses, FromEntity(entity))
	}
	return responses
}