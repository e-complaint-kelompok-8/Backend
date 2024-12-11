package news

import (
	"capstone/entities"
	"capstone/repositories/models"
	"capstone/utils"
	"errors"

	"gorm.io/gorm"
)

type NewsRepositoryInterface interface {
	GetAllNews(page int, limit int) ([]entities.News, int64, error)
	GetNewsByID(id string) (entities.News, error)
	GetAllNewsWithComments(page, limit int) ([]entities.News, int64, error)
	GetNewsByIDWithComments(id string) (entities.News, error)
	IsCategoryValid(categoryID int) (bool, error)
	CreateNews(news entities.News) (entities.News, error)
	UpdateNewsByID(id string, updatedNews entities.News) (entities.News, error)
	DeleteMultipleNews(ids []int) error
	ValidateNewsIDs(ids []int) ([]int, error)
}

type NewsRepository struct {
	db *gorm.DB
}

func NewNewsRepository(db *gorm.DB) *NewsRepository {
	return &NewsRepository{db: db}
}

func (nr *NewsRepository) GetAllNews(page int, limit int) ([]entities.News, int64, error) {
	var news []models.News
	var total int64

	// Hitung total data
	err := nr.db.Model(&models.News{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// Query berita dengan pagination
	offset := (page - 1) * limit
	err = nr.db.Preload("Admin").Preload("Category").Preload("Comments.User").
		Order("created_at DESC").
		Limit(limit).Offset(offset).
		Find(&news).Error
	if err != nil {
		return nil, 0, err
	}

	// Konversi model ke entitas
	var result []entities.News
	for _, n := range news {
		result = append(result, n.ToEntitiesWithComment())
	}

	return result, total, nil
}

func (nr *NewsRepository) GetNewsByID(id string) (entities.News, error) {
	var news models.News

	// Query berita berdasarkan ID dengan Preload admin dan category
	err := nr.db.Preload("Admin").Preload("Category").Preload("Comments.User").First(&news, "id = ?", id).Error
	if err != nil {
		return entities.News{}, errors.New(utils.CapitalizeErrorMessage(errors.New("berita tidak ditemukan")))
	}

	// Konversi model ke entitas
	return news.ToEntitiesWithComment(), nil
}
