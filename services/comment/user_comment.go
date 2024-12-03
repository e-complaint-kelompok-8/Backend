package comment

import (
	"capstone/entities"
	"capstone/repositories/comment"
)

type CommentServiceInterface interface {
	AddComment(comment entities.Comment) (entities.Comment, error)
	GetCommentsByUserID(userID int) ([]entities.Comment, error)
}

type CommentService struct {
	commentRepo comment.CommentRepositoryInterface
}

func NewCommentService(repo comment.CommentRepositoryInterface) *CommentService {
	return &CommentService{commentRepo: repo}
}

func (cs *CommentService) AddComment(comment entities.Comment) (entities.Comment, error) {
	return cs.commentRepo.AddComment(comment)
}

func (cs CommentService) GetCommentsByUserID(userID int) ([]entities.Comment, error) {
	// Ambil data komentar berdasarkan user_id dari repository
	comments, err := cs.commentRepo.GetCommentsByUserID(userID)
	if err != nil {
		return nil, err
	}
	return comments, nil
}
