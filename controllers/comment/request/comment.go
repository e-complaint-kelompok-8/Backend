package request

import "capstone/entities"

type AddCommentRequest struct {
	NewsID  int    `json:"news_id" validate:"required"`
	Content string `json:"content" validate:"required"`
}

// Konversi ke entitas
func (req AddCommentRequest) ToEntity(userID int) entities.Comment {
	return entities.Comment{
		UserID:  userID,
		NewsID:  req.NewsID,
		Content: req.Content,
	}
}
