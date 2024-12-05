package comment

import (
	"capstone/entities"
	"capstone/repositories/comment"
	"errors"
)

type CommentServiceInterface interface {
	AddComment(comment entities.Comment) (entities.Comment, error)
	GetCommentsByUserID(userID int) ([]entities.Comment, error)
	GetAllComments() ([]entities.Comment, error)
}

type CommentService struct {
	commentRepo comment.CommentRepositoryInterface
}

func NewCommentService(repo comment.CommentRepositoryInterface) *CommentService {
	return &CommentService{commentRepo: repo}
}

func (cs *CommentService) AddComment(comment entities.Comment) (entities.Comment, error) {
	if comment.NewsID == 0 {
		return entities.Comment{}, errors.New("Berita tidak valid. Silakan pilih berita yang sesuai.")
	}

	if comment.Content == "" {
		return entities.Comment{}, errors.New("Konten komentar tidak boleh kosong.")
	}

	// Validasi berita
	news, err := cs.commentRepo.GetNewsByID(comment.NewsID)
	if err != nil {
		return entities.Comment{}, errors.New("Berita tidak ditemukan.")
	}

	// Validasi kategori
	exists, err := cs.commentRepo.CheckCategoryExists(news.CategoryID)
	if err != nil {
		return entities.Comment{}, errors.New("Terjadi kesalahan saat memeriksa kategori.")
	}
	if !exists {
		return entities.Comment{}, errors.New("Kategori berita tidak ditemukan.")
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

func (cs *CommentService) GetAllComments() ([]entities.Comment, error) {
	// Ambil semua komentar dari repository
	comments, err := cs.commentRepo.GetAllComments()
	if err != nil {
		return nil, err
	}
	return comments, nil
}
