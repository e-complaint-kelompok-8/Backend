package news

import (
	"capstone/entities"
	"capstone/repositories/models"

	"gorm.io/gorm"
)

type NewsRepositoryInterface interface {
	GetAllNews() ([]entities.News, error)
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
