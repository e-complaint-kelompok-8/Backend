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
	GetCommentsByUserID(userID, offset, limit int) ([]entities.Comment, int, error)
	CheckCategoryExists(id int) (bool, error)
	GetNewsByID(newsID int) (models.News, error)
	GetAllComments(offset, limit int) ([]entities.Comment, int, error)
	GetCommentByID(commentID string) (entities.Comment, error)
	DeleteComments(commentIDs []int) error
	ValidateCommentIDs(commentIDs []int) ([]int, error)
	GetCommentsByNewsID(newsID, offset, limit int) ([]entities.Comment, int, error)
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

func (cr *CommentRepository) GetCommentsByUserID(userID, offset, limit int) ([]entities.Comment, int, error) {
	var comments []models.Comment
	var total int64

	// Hitung total data
	err := cr.db.Model(&models.Comment{}).Where("user_id = ?", userID).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// Ambil komentar berdasarkan user_id dengan pagination
	err = cr.db.Preload("User").Preload("News").
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&comments).Error
	if err != nil {
		return nil, 0, err
	}

	// Konversi model ke entitas
	var result []entities.Comment
	for _, comment := range comments {
		result = append(result, comment.ToEntities())
	}

	return result, int(total), nil
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

func (cr *CommentRepository) GetAllComments(offset, limit int) ([]entities.Comment, int, error) {
	var comments []models.Comment
	var total int64

	// Hitung total data
	err := cr.db.Model(&models.Comment{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// Ambil komentar dengan pagination
	err = cr.db.Preload("User").Preload("News").Order("created_at DESC").Offset(offset).Limit(limit).Find(&comments).Error
	if err != nil {
		return nil, 0, err
	}

	// Konversi model ke entitas
	var result []entities.Comment
	for _, comment := range comments {
		result = append(result, comment.ToEntities())
	}

	return result, int(total), nil
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
