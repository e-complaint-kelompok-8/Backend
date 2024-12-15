package news

import (
	"capstone/entities"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var newsService NewsService

type NewsRepoDummy struct {
	ShouldFail bool
}

func (repo NewsRepoDummy) GetAllNews(page int, limit int) ([]entities.News, int64, error) {
	response := []entities.News{
		{
			ID: 1,
			Admin: entities.Admin{
				ID:    1,
				Email: "admin@gmail",
				Role:  "admin",
				Photo: "example.jpg",
			},
			Category: entities.Category{
				ID:          1,
				Name:        "contoh",
				Description: "ini cuman contoh",
			},
			Title:    "test",
			Content:  "ini cuman test",
			PhotoURL: "example.jpg",
			Comments: []entities.Comment{
				{
					ID: 1,
					User: entities.User{
						ID:    1,
						Name:  "test",
						Email: "test@gmail.com",
					},
				},
			},
		},
		{
			ID: 2,
			Admin: entities.Admin{
				ID:    1,
				Email: "admin@gmail",
				Role:  "admin",
				Photo: "example.jpg",
			},
			Category: entities.Category{
				ID:          1,
				Name:        "contoh",
				Description: "ini cuman contoh",
			},
			Title:    "test 2",
			Content:  "ini cuman test",
			PhotoURL: "example.jpg",
			Comments: []entities.Comment{
				{
					ID: 1,
					User: entities.User{
						ID:    1,
						Name:  "test",
						Email: "test@gmail.com",
					},
				},
			},
		},
	}

	return response, int64(len(response)), nil
}
func (repo NewsRepoDummy) GetNewsByID(id string) (entities.News, error) {
	return entities.News{
		ID: 1,
		Admin: entities.Admin{
			ID:    1,
			Email: "admin@gmail",
			Role:  "admin",
			Photo: "example.jpg",
		},
		Category: entities.Category{
			ID:          1,
			Name:        "contoh",
			Description: "ini cuman contoh",
		},
		Title:    "test",
		Content:  "ini cuman test",
		PhotoURL: "example.jpg",
		Comments: []entities.Comment{
			{
				ID: 1,
				User: entities.User{
					ID:    1,
					Name:  "test",
					Email: "test@gmail.com",
				},
			},
		},
	}, nil
}
func (repo NewsRepoDummy) GetAllNewsWithComments(page, limit int) ([]entities.News, int64, error) {
	response := []entities.News{
		{
			ID: 1,
			Admin: entities.Admin{
				ID:    1,
				Email: "admin@gmail",
				Role:  "admin",
				Photo: "example.jpg",
			},
			Category: entities.Category{
				ID:          1,
				Name:        "contoh",
				Description: "ini cuman contoh",
			},
			Title:    "test",
			Content:  "ini cuman test",
			PhotoURL: "example.jpg",
			Comments: []entities.Comment{
				{
					ID: 1,
					User: entities.User{
						ID:    1,
						Name:  "test",
						Email: "test@gmail.com",
					},
				},
			},
		},
		{
			ID: 2,
			Admin: entities.Admin{
				ID:    1,
				Email: "admin@gmail",
				Role:  "admin",
				Photo: "example.jpg",
			},
			Category: entities.Category{
				ID:          1,
				Name:        "contoh",
				Description: "ini cuman contoh",
			},
			Title:    "test 2",
			Content:  "ini cuman test",
			PhotoURL: "example.jpg",
			Comments: []entities.Comment{
				{
					ID: 1,
					User: entities.User{
						ID:    1,
						Name:  "test",
						Email: "test@gmail.com",
					},
				},
			},
		},
	}

	return response, int64(len(response)), nil
}
func (repo NewsRepoDummy) GetNewsByIDWithComments(id string) (entities.News, error) {
	return entities.News{
		ID: 1,
		Admin: entities.Admin{
			ID:    1,
			Email: "admin@gmail",
			Role:  "admin",
			Photo: "example.jpg",
		},
		Category: entities.Category{
			ID:          1,
			Name:        "contoh",
			Description: "ini cuman contoh",
		},
		Title:    "test",
		Content:  "ini cuman test",
		PhotoURL: "example.jpg",
		Comments: []entities.Comment{
			{
				ID: 1,
				User: entities.User{
					ID:    1,
					Name:  "test",
					Email: "test@gmail.com",
				},
			},
		},
	}, nil
}
func (repo NewsRepoDummy) IsCategoryValid(categoryID int) (bool, error) {
	if repo.ShouldFail {
		return false, nil // Simulasikan kategori tidak valid
	}
	return true, nil // Kategori valid
}
func (repo NewsRepoDummy) CreateNews(news entities.News) (entities.News, error) {
	return entities.News{
		ID:      1,
		AdminID: 1,
		Admin: entities.Admin{
			ID:    1,
			Email: "admin@gmail",
			Role:  "admin",
			Photo: "example.jpg",
		},
		Category: entities.Category{
			ID:          1,
			Name:        "contoh",
			Description: "ini cuman contoh",
		},
		Title:    "test",
		Content:  "ini cuman test",
		PhotoURL: "example.jpg",
	}, nil
}
func (repo NewsRepoDummy) UpdateNewsByID(id string, updatedNews entities.News) (entities.News, error) {
	return entities.News{
		ID: 1,
		Admin: entities.Admin{
			ID:    1,
			Email: "admin@gmail",
			Role:  "admin",
			Photo: "example.jpg",
		},
		Category: entities.Category{
			ID:          1,
			Name:        "contoh",
			Description: "ini cuman contoh",
		},
		Title:    "updatetest",
		Content:  "update ini cuman test",
		PhotoURL: "updateexample.jpg",
	}, nil
}
func (repo NewsRepoDummy) DeleteMultipleNews(ids []int) error {
	return nil
}
func (repo NewsRepoDummy) ValidateNewsIDs(ids []int) ([]int, error) {
	// Mock ID yang ada di database
	existingIDs := []int{1, 2, 3}

	// Filter ID yang valid berdasarkan existingIDs
	var validIDs []int
	for _, id := range ids {
		for _, existingID := range existingIDs {
			if id == existingID {
				validIDs = append(validIDs, id)
			}
		}
	}

	// Jika ShouldFail aktif, simulasi error validasi
	if repo.ShouldFail {
		return nil, errors.New("error validasi ID berita")
	}

	return validIDs, nil
}

func setupTestService() {
	repo := NewsRepoDummy{}
	newsService = *NewNewsService(repo)
}

func TestNewsService_GetAllNews(t *testing.T) {
	setupTestService()

	t.Run("sukses", func(t *testing.T) {
		// Data dummy untuk pengujian
		news, total, err := newsService.GetAllNews(0, 2)

		// Periksa apakah error tidak terjadi
		assert.NoError(t, err)
		assert.Equal(t, int64(2), total)
		assert.Equal(t, "test", news[0].Title)
		assert.Equal(t, 1, news[0].ID)
		assert.Equal(t, 1, news[0].Comments[0].ID)
		assert.Equal(t, "test 2", news[1].Title)
		assert.Equal(t, 2, news[1].ID)
		assert.Equal(t, 1, news[1].Comments[0].ID)
	})
}

func TestNewsService_GetNewsByID(t *testing.T) {
	setupTestService()

	t.Run("sukses", func(t *testing.T) {
		// Data dummy untuk pengujian
		news, err := newsService.GetNewsByID("1")

		// Periksa apakah error tidak terjadi
		assert.NoError(t, err)
		assert.Equal(t, "test", news.Title)
		assert.Equal(t, 1, news.ID)
		assert.Equal(t, 1, news.Comments[0].ID)
	})
}

func TestNewsService_GetAllNewsWithComments(t *testing.T) {
	setupTestService()

	t.Run("sukses", func(t *testing.T) {
		// Data dummy untuk pengujian
		news, total, err := newsService.GetAllNewsWithComments(0, 2)

		// Periksa apakah error tidak terjadi
		assert.NoError(t, err)
		assert.Equal(t, int64(2), total)
		assert.Equal(t, "test", news[0].Title)
		assert.Equal(t, 1, news[0].ID)
		assert.Equal(t, 1, news[0].Comments[0].ID)
		assert.Equal(t, "test 2", news[1].Title)
		assert.Equal(t, 2, news[1].ID)
		assert.Equal(t, 1, news[1].Comments[0].ID)
	})
}
