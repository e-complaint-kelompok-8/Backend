package news

import (
	"capstone/controllers/news/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (nc *NewsController) GetAllNewsWithComments(c echo.Context) error {
	newsList, err := nc.newsService.GetAllNewsWithComments()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to fetch news",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"news":    response.NewsFromEntities(newsList),
	})
}

func (nc *NewsController) GetNewsDetailByAdmin(c echo.Context) error {
	id := c.Param("id")

	news, err := nc.newsService.GetNewsByIDWithComments(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"news":    response.NewFromEntities(news),
	})
}
