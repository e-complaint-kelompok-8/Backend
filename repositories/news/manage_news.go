package news

import (
	"capstone/entities"
	"capstone/repositories/models"
	"errors"
)

func (nr *NewsRepository) GetAllNewsWithComments() ([]entities.News, error) {
	var newsList []models.News

	err := nr.db.Preload("Admin").
		Preload("Category").
		Preload("Comments.User").
		Find(&newsList).Error

	if err != nil {
		return nil, err
	}

	// Konversi model ke entitas
	var result []entities.News
	for _, news := range newsList {
		newsEntity := news.ToEntitiesWithComment()
		result = append(result, newsEntity)
	}
	return result, nil
}

func (nr *NewsRepository) GetNewsByIDWithComments(id string) (entities.News, error) {
	var news models.News

	err := nr.db.Preload("Admin").
		Preload("Category").
		Preload("Comments.User").
		First(&news, "id = ?", id).Error

	if err != nil {
		return entities.News{}, errors.New("news not found")
	}

	return news.ToEntitiesWithComment(), nil
}

func (nr *NewsRepository) IsCategoryValid(categoryID int) (bool, error) {
	var count int64
	err := nr.db.Model(&models.Category{}).Where("id = ?", categoryID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (nr *NewsRepository) CreateNews(news entities.News) (entities.News, error) {
	newsModel := models.FromEntitiesNews(news)

	// Simpan berita baru
	err := nr.db.Create(&newsModel).Error
	if err != nil {
		return entities.News{}, err
	}

	// Ambil berita yang baru saja disimpan, dengan memuat Admin dan Category terkait
	err = nr.db.Preload("Admin").Preload("Category").First(&newsModel, "id = ?", newsModel.ID).Error
	if err != nil {
		return entities.News{}, err
	}

	// Kembalikan data dalam bentuk entitas
	return newsModel.ToEntities(), nil
}
