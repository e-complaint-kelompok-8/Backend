package comment

import (
	"capstone/entities"
	"capstone/repositories/comment"
	"capstone/utils"
	"errors"
)

type CommentServiceInterface interface {
	AddComment(comment entities.Comment) (entities.Comment, error)
	GetCommentsByUserID(userID, page, limit int) ([]entities.Comment, int, error)
	GetAllComments(page, limit int) ([]entities.Comment, int, error)
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

func (cs *CommentService) GetCommentsByUserID(userID, page, limit int) ([]entities.Comment, int, error) {
	// Hitung offset
	offset := (page - 1) * limit

	// Ambil komentar dari repository dengan pagination
	comments, total, err := cs.commentRepo.GetCommentsByUserID(userID, offset, limit)
	if err != nil {
		return nil, 0, err
	}
	return comments, total, nil
}

func (cs *CommentService) GetAllComments(page, limit int) ([]entities.Comment, int, error) {
	// Hitung offset
	offset := (page - 1) * limit

	// Ambil komentar dari repository dengan pagination
	comments, total, err := cs.commentRepo.GetAllComments(offset, limit)
	if err != nil {
		return nil, 0, err
	}
	return comments, total, nil
}

func (cs *CommentService) GetCommentByID(commentID string) (entities.Comment, error) {
	// Ambil detail komentar dari repository
	comment, err := cs.commentRepo.GetCommentByID(commentID)
	if err != nil {
		return entities.Comment{}, err
	}
	return comment, nil
}
