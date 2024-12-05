package comment

import (
	"capstone/entities"
	"capstone/repositories/models"
	"capstone/utils"
	"errors"

	"gorm.io/gorm"
)

type CommentRepositoryInterface interface {
	AddComment(comment entities.Comment) (entities.Comment, error)
	GetCommentsByUserID(userID int) ([]entities.Comment, error)
	CheckCategoryExists(id int) (bool, error)
	GetNewsByID(newsID int) (models.News, error)
	GetAllComments() ([]entities.Comment, error)
	GetCommentByID(commentID string) (entities.Comment, error)
}

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

func (cr *CommentRepository) AddComment(comment entities.Comment) (entities.Comment, error) {
	commentModel := models.FromEntitiesComment(comment)

	// Simpan komentar ke database
	if err := cr.db.Create(&commentModel).Error; err != nil {
		return entities.Comment{}, errors.New(utils.CapitalizeErrorMessage(errors.New("gagal menambahkan komentar")))
	}

	// Preload relasi user dan news
	if err := cr.db.Preload("User").Preload("News").First(&commentModel, "id = ?", commentModel.ID).Error; err != nil {
		return entities.Comment{}, err
	}

	return commentModel.ToEntities(), nil
}

func (cr *CommentRepository) GetCommentsByUserID(userID int) ([]entities.Comment, error) {
	var comments []models.Comment

	// Query database untuk mengambil komentar berdasarkan user_id
	err := cr.db.Preload("User").Preload("News").Where("user_id = ?", userID).Find(&comments).Error
	if err != nil {
		return nil, err
	}

	// Konversi model ke entitas
	var result []entities.Comment
	for _, comment := range comments {
		result = append(result, comment.ToEntities())
	}

	return result, nil
}

func (ar *CommentRepository) CheckCategoryExists(categoryID int) (bool, error) {
	var count int64
	err := ar.db.Model(&models.Category{}).Where("id = ?", categoryID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil // True jika kategori ditemukan
}

func (ar *CommentRepository) GetNewsByID(newsID int) (models.News, error) {
	var news models.News
	err := ar.db.Preload("Category").First(&news, "id = ?", newsID).Error
	if err != nil {
		return models.News{}, err
	}
	return news, nil
}

func (cr *CommentRepository) GetAllComments() ([]entities.Comment, error) {
	var comments []models.Comment

	// Query database untuk mengambil semua komentar
	err := cr.db.Preload("User").Preload("News").Find(&comments).Error
	if err != nil {
		return nil, err
	}

	// Konversi model ke entitas
	var result []entities.Comment
	for _, comment := range comments {
		result = append(result, comment.ToEntities())
	}

	return result, nil
}

func (cr *CommentRepository) GetCommentByID(commentID string) (entities.Comment, error) {
	var comment models.Comment

	// Query database untuk mengambil komentar berdasarkan ID
	err := cr.db.Preload("User").Preload("News").Where("id = ?", commentID).First(&comment).Error
	if err != nil {
		return entities.Comment{}, err
	}

	// Konversi model ke entitas
	return comment.ToEntities(), nil
}
