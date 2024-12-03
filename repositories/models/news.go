package models

import (
	"capstone/entities"
	"time"
)

// News struct
type News struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	Title     string    `gorm:"type:varchar(255);not null"`
	Content   string    `gorm:"type:text;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func FromEntitiesNews(news entities.News) News {
	return News{
		ID:        news.ID,
		Title:     news.Title,
		Content:   news.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (news News) ToEntities() entities.News {
	return entities.News{
		ID:        news.ID,
		Title:     news.Title,
		Content:   news.Content,
		CreatedAt: news.CreatedAt,
		UpdatedAt: news.UpdatedAt,
	}
}
