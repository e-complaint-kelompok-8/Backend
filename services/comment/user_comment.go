package comment

import (
	"capstone/entities"
	"capstone/repositories/comment"
	"capstone/utils"
	"errors"
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
	if comment.NewsID == 0 {
		return entities.Comment{}, errors.New(utils.CapitalizeErrorMessage(errors.New("pilih berita")))
	}

	if comment.Content == "" {
		return entities.Comment{}, errors.New(utils.CapitalizeErrorMessage(errors.New("tambahkan komentar")))
	}
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
