package news

import (
	"capstone/entities"
	"capstone/repositories/models"
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
