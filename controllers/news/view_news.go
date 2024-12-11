package news

import (
	"capstone/controllers/news/response"
	"capstone/services/news"
	"math"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type NewsController struct {
	newsService news.NewsServiceInterface
}

func NewNewsController(service news.NewsServiceInterface) *NewsController {
	return &NewsController{newsService: service}
}

func (nc *NewsController) GetAllNews(c echo.Context) error {
	// Ambil parameter pagination
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page <= 0 {
		page = 1 // Default page
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit <= 0 {
		limit = 8 // Default limit
	}

	// Panggil service dengan pagination
	news, total, err := nc.newsService.GetAllNews(page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to retrieve news",
		})
	}

	// Hitung total halaman
	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "News retrieved successfully",
		"news":    response.NewsFromEntities(news),
		"pagination": map[string]interface{}{
			"page":       page,
			"limit":      limit,
			"total":      total,
			"totalPages": totalPages,
		},
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
