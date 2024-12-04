package news

import (
	"capstone/controllers/news/response"
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
		"news":    response.NewsFromEntities(news),
	})
}

func (nc *NewsController) GetNewsByID(c echo.Context) error {
	// Ambil ID dari parameter URL
	id := c.Param("id")

	// Panggil service untuk mendapatkan detail berita berdasarkan ID
	news, err := nc.newsService.GetNewsByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "News not found",
		})
	}

	// Format response
	response := response.NewFromEntities(news)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "News retrieved successfully",
		"news":    response,
	})
}
