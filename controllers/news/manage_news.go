package news

import (
	"capstone/controllers/news/request"
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

func (nc *NewsController) AddNews(c echo.Context) error {
	req := request.AddNewsRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request payload"})
	}

	// Konversi request ke entitas
	newsEntity := req.ToEntity()

	// Simpan berita baru dan dapatkan data lengkapnya
	news, err := nc.newsService.AddNews(newsEntity)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	// Kirimkan data berita lengkap dalam response
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "News created successfully",
		"data":    response.NewFromEntities(news),
	})
}

func (nc *NewsController) UpdateNewsByAdmin(c echo.Context) error {
	id := c.Param("id")

	req := request.AddNewsRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request payload"})
	}

	// Konversi request ke entitas
	newsEntity := req.ToEntity()

	// Panggil service untuk update berita
	updatedNews, err := nc.newsService.UpdateNewsByID(id, newsEntity)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "News updated successfully",
		"data":    response.NewFromEntities(updatedNews),
	})
}
