package models

import (
	"capstone/entities"
	"time"
)

type Comment struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	UserID    int       `gorm:"not null"`
	User      User      `gorm:"foreignKey:UserID"`
	NewsID    int       `gorm:"not null"`
	News      News      `gorm:"foreignKey:NewsID"`
	Content   string    `gorm:"type:text;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func FromEntitiesComment(comment entities.Comment) Comment {
	return Comment{
		ID:        comment.ID,
		UserID:    comment.UserID,
		NewsID:    comment.NewsID,
		Content:   comment.Content,
		CreatedAt: time.Now(),
	}
}

func (comment Comment) ToEntities() entities.Comment {
	return entities.Comment{
		ID:        comment.ID,
		UserID:    comment.UserID,
		User:      comment.User.ToEntities(),
		NewsID:    comment.NewsID,
		News:      comment.News.ToEntities(),
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt,
	}
}
