package news

import (
	"capstone/entities"
	"capstone/repositories/models"
	"capstone/utils"
	"errors"

	"gorm.io/gorm"
)

type NewsRepositoryInterface interface {
	GetAllNews() ([]entities.News, error)
	GetNewsByID(id string) (entities.News, error)
	GetAllNewsWithComments() ([]entities.News, error)
	GetNewsByIDWithComments(id string) (entities.News, error)
}

type NewsRepository struct {
	db *gorm.DB
}

func NewNewsRepository(db *gorm.DB) *NewsRepository {
	return &NewsRepository{db: db}
}

func (nr *NewsRepository) GetAllNews() ([]entities.News, error) {
	var news []models.News

	// Query semua berita dengan Preload admin dan category
	err := nr.db.Preload("Admin").Preload("Category").Find(&news).Error
	if err != nil {
		return nil, err
	}

	// Konversi model ke entitas
	var result []entities.News
	for _, n := range news {
		result = append(result, n.ToEntities())
	}

	return result, nil
}

func (nr *NewsRepository) GetNewsByID(id string) (entities.News, error) {
	var news models.News

	// Query berita berdasarkan ID dengan Preload admin dan category
	err := nr.db.Preload("Admin").Preload("Category").First(&news, "id = ?", id).Error
	if err != nil {
		return entities.News{}, errors.New(utils.CapitalizeErrorMessage(errors.New("berita tidak ditemukan")))
	}

	// Konversi model ke entitas
	return news.ToEntities(), nil
}
