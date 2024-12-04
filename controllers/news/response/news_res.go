package response

import (
	"capstone/entities"
	"time"
)

type NewsResponse struct {
	ID        int       `json:"id"`
	Admin     Admin     `json:"admin"`
	Category  Category  `json:"category"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	PhotoURL  string    `json:"photo_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Admin struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Category struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewFromEntities(news entities.News) NewsResponse {
	return NewsResponse{
		ID: news.ID,
		Admin: Admin{
			ID:        news.Admin.ID,
			Email:     news.Admin.Email,
			Role:      news.Admin.Role,
			CreatedAt: news.Admin.CreatedAt,
			UpdatedAt: news.Admin.UpdatedAt,
		},
		Category: Category{
			ID:          news.Category.ID,
			Name:        news.Category.Name,
			Description: news.Category.Description,
			CreatedAt:   news.Category.CreatedAt,
			UpdatedAt:   news.Category.UpdatedAt,
		},
		Title:     news.Title,
		Content:   news.Content,
		PhotoURL:  news.PhotoURL,
		CreatedAt: news.CreatedAt,
		UpdatedAt: news.UpdatedAt,
	}
}

func NewsFromEntities(news []entities.News) []NewsResponse {
	var responses []NewsResponse
	for _, new := range news {
		responses = append(responses, NewFromEntities(new))
	}
	return responses
}
