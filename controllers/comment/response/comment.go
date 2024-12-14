package response

import (
	"capstone/entities"
	"time"
)

type CommentResponse struct {
	ID      int          `json:"id"`
	Content string       `json:"content"`
	User    UserResponse `json:"user"`
	News    NewsResponse `json:"news"`
}

type UserResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone_number"`
}

type NewsResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	PhotoURL  string    `json:"photo_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Comment struct {
	ID        int    `json:"id"`
	Content   string `json:"content"`
	UserName  string `json:"user_name"`
	NewsTitle string `json:"news_title"`
	CreatedAt string `json:"created_at"`
}

func FromEntityComment(comment entities.Comment) CommentResponse {
	return CommentResponse{
		ID:      comment.ID,
		Content: comment.Content,
		User: UserResponse{
			ID:    comment.User.ID,
			Name:  comment.User.Name,
			Email: comment.User.Email,
			Phone: comment.User.Phone,
		},
		News: NewsResponse{
			ID:        comment.News.ID,
			Title:     comment.News.Title,
			Content:   comment.News.Content,
			PhotoURL:  comment.News.PhotoURL,
			CreatedAt: comment.News.CreatedAt,
			UpdatedAt: comment.News.UpdatedAt,
		},
	}
}

func FromEntityComments(comments []entities.Comment) []CommentResponse {
	var responses []CommentResponse
	for _, comment := range comments {
		responses = append(responses, FromEntityComment(comment))
	}
	return responses
}
