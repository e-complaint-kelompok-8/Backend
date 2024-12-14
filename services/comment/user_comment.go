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
	GetAllComments() ([]entities.Comment, error)
	GetCommentByID(commentID string) (entities.Comment, error)
	DeleteComments(commentIDs []int) error
}

type CommentService struct {
	commentRepo comment.CommentRepositoryInterface
}

func NewCommentService(repo comment.CommentRepositoryInterface) *CommentService {
	return &CommentService{commentRepo: repo}
}

func (cs *CommentService) AddComment(comment entities.Comment) (entities.Comment, error) {
	if comment.NewsID == 0 {
		return entities.Comment{}, errors.New(utils.CapitalizeErrorMessage(errors.New("berita tidak valid. silakan pilih berita yang sesuai")))
	}

	if comment.Content == "" {
		return entities.Comment{}, errors.New(utils.CapitalizeErrorMessage(errors.New("konten komentar tidak boleh kosong")))
	}

	// Validasi berita
	news, err := cs.commentRepo.GetNewsByID(comment.NewsID)
	if err != nil {
		return entities.Comment{}, errors.New(utils.CapitalizeErrorMessage(errors.New("berita tidak ditemukan")))
	}

	// Validasi kategori
	exists, err := cs.commentRepo.CheckCategoryExists(news.CategoryID)
	if err != nil {
		return entities.Comment{}, errors.New(utils.CapitalizeErrorMessage(errors.New("terjadi kesalahan saat memeriksa kategori")))
	}
	if !exists {
		return entities.Comment{}, errors.New(utils.CapitalizeErrorMessage(errors.New("kategori berita tidak ditemukan")))
	}

	return cs.commentRepo.AddComment(comment)
}

func (cs *CommentService) GetCommentsByUserID(userID int) ([]entities.Comment, error) {
	// Ambil data komentar berdasarkan user_id dari repository
	comments, err := cs.commentRepo.GetCommentsByUserID(userID)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (cs *CommentService) GetAllComments() ([]entities.Comment, error) {
	// Ambil semua komentar dari repository
	comments, err := cs.commentRepo.GetAllComments()
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (cs *CommentService) GetCommentByID(commentID string) (entities.Comment, error) {
	// Ambil detail komentar dari repository
	comment, err := cs.commentRepo.GetCommentByID(commentID)
	if err != nil {
		return entities.Comment{}, err
	}
	return comment, nil
}
