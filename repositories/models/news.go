package models

import (
	"capstone/entities"
	"time"
)

// News struct
type News struct {
	ID         int       `gorm:"primaryKey;autoIncrement"`
	AdminID    int       `gorm:"not null"`
	Admin      Admin     `gorm:"foreignKey:AdminID"`
	CategoryID int       `gorm:"not null"`
	Category   Category  `gorm:"foreignKey:CategoryID"`
	Title      string    `gorm:"type:varchar(255);not null"`
	Content    string    `gorm:"type:text;not null"`
	PhotoURL   string    `gorm:"type:varchar(255);not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}

func (news News) ToEntities() entities.News {
	return entities.News{
		ID:        news.ID,
		Admin:     news.Admin.ToEntities(),
		Category:  news.Category.ToEntities(),
		Title:     news.Title,
		Content:   news.Content,
		PhotoURL:  news.PhotoURL,
		CreatedAt: news.CreatedAt,
		UpdatedAt: news.UpdatedAt,
	}
}

func FromEntitiesNews(news entities.News) News {
	return News{
		ID:        news.ID,
		Title:     news.Title,
		Content:   news.Content,
		PhotoURL:  news.PhotoURL,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
