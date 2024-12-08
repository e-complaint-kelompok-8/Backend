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

func (nr *NewsRepository) UpdateNewsByID(id string, updatedNews entities.News) (entities.News, error) {
	var existingNews models.News

	// Cari berita berdasarkan ID
	err := nr.db.First(&existingNews, "id = ?", id).Error
	if err != nil {
		return entities.News{}, errors.New("news not found")
	}

	// Update data berita
	existingNews.Title = updatedNews.Title
	existingNews.Content = updatedNews.Content
	existingNews.PhotoURL = updatedNews.PhotoURL
	existingNews.CategoryID = updatedNews.CategoryID
	existingNews.Date = updatedNews.Date

	// Simpan perubahan
	err = nr.db.Save(&existingNews).Error
	if err != nil {
		return entities.News{}, err
	}

	// Preload Admin dan Category untuk response
	err = nr.db.Preload("Admin").Preload("Category").Preload("Comments.User").First(&existingNews, "id = ?", existingNews.ID).Error
	if err != nil {
		return entities.News{}, err
	}

	return existingNews.ToEntitiesWithComment(), nil
}

func (nr *NewsRepository) ValidateNewsIDs(ids []int) ([]int, error) {
	var existingIDs []int
	err := nr.db.Model(&models.News{}).Where("id IN ?", ids).Pluck("id", &existingIDs).Error
	if err != nil {
		return nil, err
	}
	return existingIDs, nil
}

func (nr *NewsRepository) DeleteMultipleNews(ids []int) error {
	// Hapus berita berdasarkan id yang diterima
	err := nr.db.Where("id IN ?", ids).Delete(&models.News{}).Error
	if err != nil {
		return err
	}

	return nil
}
