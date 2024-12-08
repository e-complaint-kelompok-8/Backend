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
	Comments  []Comment `json:"comments"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Comment struct {
	ID        int       `json:"id"`
	User      User      `json:"user"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
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

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone_number"`
	Email string `json:"email"`
}

func NewFromEntities(news entities.News) NewsResponse {
	var comments []Comment
	for _, comment := range news.Comments {
		comments = append(comments, Comment{
			ID:        comment.ID,
			User:      User{ID: comment.User.ID, Name: comment.User.Name, Email: comment.User.Email},
			Content:   comment.Content,
			CreatedAt: comment.CreatedAt,
		})
	}

	return NewsResponse{
		ID:        news.ID,
		Admin:     Admin{ID: news.Admin.ID, Email: news.Admin.Email, Role: news.Admin.Role, CreatedAt: news.Admin.CreatedAt, UpdatedAt: news.Admin.UpdatedAt},
		Category:  Category{ID: news.Category.ID, Name: news.Category.Name, Description: news.Category.Description, CreatedAt: news.Category.CreatedAt, UpdatedAt: news.Category.UpdatedAt},
		Title:     news.Title,
		Content:   news.Content,
		PhotoURL:  news.PhotoURL,
		Comments:  comments,
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
