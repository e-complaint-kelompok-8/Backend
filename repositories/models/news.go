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
	Date       time.Time `gorm:"type:date;not null"`
	Comments   []Comment `gorm:"foreignKey:NewsID;constraint:OnDelete:CASCADE"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}

func (news News) ToEntities() entities.News {
	return entities.News{
		ID:         news.ID,
		AdminID:    news.AdminID,
		Admin:      news.Admin.ToEntities(),
		CategoryID: news.CategoryID,
		Category:   news.Category.ToEntities(),
		Title:      news.Title,
		Content:    news.Content,
		PhotoURL:   news.PhotoURL,
		Date:       news.Date,
		CreatedAt:  news.CreatedAt,
		UpdatedAt:  news.UpdatedAt,
	}
}

func FromEntitiesNews(news entities.News) News {
	return News{
		ID:         news.ID,
		AdminID:    news.AdminID,
		CategoryID: news.CategoryID,
		Title:      news.Title,
		Content:    news.Content,
		PhotoURL:   news.PhotoURL,
		Date:       news.Date,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}

func (news News) ToEntitiesWithComment() entities.News {
	var comments []entities.Comment
	for _, comment := range news.Comments {
		comments = append(comments, comment.ToEntities())
	}

	return entities.News{
		ID:        news.ID,
		Admin:     news.Admin.ToEntities(),
		Category:  news.Category.ToEntities(),
		Title:     news.Title,
		Content:   news.Content,
		PhotoURL:  news.PhotoURL,
		Date:      news.Date,
		Comments:  comments,
		CreatedAt: news.CreatedAt,
		UpdatedAt: news.UpdatedAt,
	}
}
