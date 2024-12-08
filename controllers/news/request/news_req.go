package request

import (
	"capstone/entities"
	"time"
)

type AddNewsRequest struct {
	AdminID    int    `json:"admin_id"`
	CategoryID int    `json:"category_id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	PhotoURL   string `json:"photo_url"`
	Date       string `json:"date"`
}

func (req AddNewsRequest) ToEntity() entities.News {
	parsedDate, _ := time.Parse("2006-01-02", req.Date)
	return entities.News{
		AdminID:    req.AdminID,
		CategoryID: req.CategoryID,
		Title:      req.Title,
		Content:    req.Content,
		PhotoURL:   req.PhotoURL,
		Date:       parsedDate,
	}
}
