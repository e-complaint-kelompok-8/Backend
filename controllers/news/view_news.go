package news

import (
	"capstone/services/news"
	"net/http"

	"github.com/labstack/echo/v4"
)

type NewsController struct {
	newsService news.NewsServiceInterface
}

func NewNewsController(service news.NewsServiceInterface) *NewsController {
	return &NewsController{newsService: service}
}

func (nc *NewsController) GetAllNews(c echo.Context) error {
	news, err := nc.newsService.GetAllNews()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to retrieve news",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "News retrieved successfully",
		"news":    news,
	})
}
