package news

import (
	"capstone/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewsService_GetNewsByIDWithComments(t *testing.T) {
	setupTestService()

	t.Run("sukses", func(t *testing.T) {
		// Data dummy untuk pengujian
		news, err := newsService.GetNewsByIDWithComments("1")

		// Periksa apakah error tidak terjadi
		assert.NoError(t, err)
		assert.Equal(t, "test", news.Title)
		assert.Equal(t, 1, news.ID)
		assert.Equal(t, 1, news.Comments[0].ID)
	})
}

func TestNewsService_AddNews(t *testing.T) {
	setupTestService()

	t.Run("sukses", func(t *testing.T) {
		// Mock sukses
		newsRepo := NewsRepoDummy{
			ShouldFail: false,
		}
		newsService := NewNewsService(newsRepo)

		// Data dummy untuk pengujian
		news, err := newsService.AddNews(entities.News{
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
		})

		// Periksa apakah error tidak terjadi
		assert.NoError(t, err)
		assert.Equal(t, "test", news.Title)
		assert.Equal(t, 1, news.ID)
	})

	t.Run("gagal - kategori tidak valid", func(t *testing.T) {
		// Mock gagal
		newsRepo := NewsRepoDummy{
			ShouldFail: true,
		}
		newsService := NewNewsService(newsRepo)

		// Data dummy untuk pengujian
		_, err := newsService.AddNews(entities.News{
			ID: 1,
			Admin: entities.Admin{
				ID:    1,
				Email: "admin@gmail",
				Role:  "admin",
				Photo: "example.jpg",
			},
			Category: entities.Category{
				ID:          99, // ID kategori yang tidak valid
				Name:        "invalid",
				Description: "invalid category",
			},
			Title:    "test",
			Content:  "ini cuman test",
			PhotoURL: "example.jpg",
		})

		// Periksa apakah error terjadi dengan pesan yang sesuai
		assert.Error(t, err)
		assert.Equal(t, "ID Kategori Tidak Valid", err.Error())
	})

	t.Run("gagal - admin tidak ditemukan", func(t *testing.T) {
		// Mock sukses
		newsRepo := NewsRepoDummy{
			ShouldFail: false,
		}
		newsService := NewNewsService(newsRepo)

		// Data dummy untuk pengujian
		_, err := newsService.AddNews(entities.News{
			ID: 1,
			Category: entities.Category{
				ID:          1,
				Name:        "contoh",
				Description: "ini cuman contoh",
			},
			Title:    "test",
			Content:  "ini cuman test",
			PhotoURL: "example.jpg",
		})

		// Periksa apakah error terjadi dengan pesan yang sesuai
		assert.Error(t, err)
		assert.Equal(t, "Admin Tidak Ditemukan", err.Error())
	})
}

func TestNewsService_UpdateNewsByID(t *testing.T) {
	setupTestService()

	t.Run("sukses", func(t *testing.T) {
		// Mock sukses
		newsRepo := NewsRepoDummy{
			ShouldFail: false,
		}
		newsService := NewNewsService(newsRepo)

		// Data dummy untuk pengujian
		news, err := newsService.UpdateNewsByID("1", entities.News{
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
		})

		// Periksa apakah error tidak terjadi
		assert.NoError(t, err)
		assert.Equal(t, "updatetest", news.Title)
		assert.Equal(t, 1, news.ID)
	})

	t.Run("gagal - kategori tidak valid", func(t *testing.T) {
		// Mock gagal
		newsRepo := NewsRepoDummy{
			ShouldFail: true,
		}
		newsService := NewNewsService(newsRepo)

		// Data dummy untuk pengujian
		_, err := newsService.UpdateNewsByID("1", entities.News{
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
		})

		// Periksa apakah error terjadi dengan pesan yang sesuai
		assert.Error(t, err)
		assert.Equal(t, "ID Kategori Tidak Benar", err.Error())
	})
}

func TestNewsService_DeleteMultipleNews(t *testing.T) {
	setupTestService()

	t.Run("sukses", func(t *testing.T) {
		// Mock sukses
		newsRepo := NewsRepoDummy{
			ShouldFail: false,
		}
		newsService := NewNewsService(newsRepo)

		// Data dummy untuk pengujian
		var ids = []int{1, 2, 3}
		err := newsService.DeleteMultipleNews(ids)

		// Periksa apakah error tidak terjadi
		assert.NoError(t, err)
	})

	t.Run("gagal - tidak ada berita yang dipilih untuk dihapus", func(t *testing.T) {
		// Mock gagal
		newsRepo := NewsRepoDummy{
			ShouldFail: false,
		}
		newsService := NewNewsService(newsRepo)

		// Data dummy untuk pengujian
		var ids = []int{}
		err := newsService.DeleteMultipleNews(ids)

		// Periksa apakah error terjadi dengan pesan yang sesuai
		assert.Error(t, err)
		assert.Equal(t, "Tidak Ada Berita Yang Dipilih Untuk Dihapus", err.Error())
	})

	t.Run("gagal - berita tidak ditemukan", func(t *testing.T) {
		// Mock gagal
		newsRepo := NewsRepoDummy{
			ShouldFail: false,
		}
		newsService := NewNewsService(newsRepo)

		// Data dummy untuk pengujian
		var ids = []int{10, 11, 12} // ID berita yang tidak valid
		err := newsService.DeleteMultipleNews(ids)

		// Periksa apakah error terjadi dengan pesan yang sesuai
		assert.Error(t, err)
		assert.Equal(t, "Berita Tidak Ditemukan", err.Error()) // Sesuai dengan pesan error di service
	})

	t.Run("gagal - beberapa ID berita tidak ditemukan", func(t *testing.T) {
		// Mock gagal
		newsRepo := NewsRepoDummy{
			ShouldFail: false,
		}
		newsService := NewNewsService(newsRepo)

		// Data dummy untuk pengujian
		var ids = []int{1, 4, 5} // Beberapa ID valid, beberapa ID tidak valid
		err := newsService.DeleteMultipleNews(ids)

		// Periksa apakah error terjadi dengan pesan yang sesuai
		assert.Error(t, err)
		assert.Equal(t, "Beberapa ID Berita Tidak Ditemukan", err.Error()) // Sesuai dengan pesan error di service
	})
}
