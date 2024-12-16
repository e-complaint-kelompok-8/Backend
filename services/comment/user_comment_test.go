package comment

import (
	"capstone/entities"
	"capstone/repositories/models"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var commentService CommentService

type CommentRepoDummy struct {
	ShouldFail         bool
	ShouldFailGetNews  bool
	ShouldFailCategory bool
	ShouldFailDelete   bool
}

func (repo CommentRepoDummy) AddComment(comment entities.Comment) (entities.Comment, error) {
	return entities.Comment{
		ID:     1,
		UserID: 1,
		User: entities.User{
			ID:    1,
			Name:  "test",
			Email: "test@gmail.com",
		},
		NewsID: 1,
		News: entities.News{
			ID:      1,
			Title:   "test berita",
			Content: "test content",
		},
		Content: "test content",
	}, nil
}
func (repo CommentRepoDummy) GetCommentsByUserID(userID, offset, limit int) ([]entities.Comment, int, error) {
	responses := []entities.Comment{
		{
			ID:     1,
			UserID: 1,
			User: entities.User{
				ID:    1,
				Name:  "test",
				Email: "test@gmail.com",
			},
			NewsID: 1,
			News: entities.News{
				ID:      1,
				Title:   "test berita",
				Content: "test content",
			},
			Content: "test content",
		},
	}

	if repo.ShouldFail {
		return []entities.Comment{}, 0, errors.New("user tidak ditemukan")
	}
	return responses, len(responses), nil

}
func (repo CommentRepoDummy) CheckCategoryExists(id int) (bool, error) {
	if repo.ShouldFail {
		return false, errors.New("terjadi kesalahan saat memeriksa kategori")
	}
	if id == 2 {
		return false, nil
	}

	return true, nil

}
func (repo CommentRepoDummy) GetNewsByID(newsID int) (models.News, error) {
	if newsID == 1 {
		return models.News{
			ID:      1,
			AdminID: 1,
			Admin: models.Admin{
				ID:    1,
				Email: "test@gmail.com",
				Role:  "admin",
			},
			CategoryID: 1,
			Category: models.Category{
				ID:   1,
				Name: "test",
			},
			Title:   "title test",
			Content: "test content",
		}, nil
	}

	return models.News{}, errors.New("berita tidak ditemukan")

}
func (repo CommentRepoDummy) GetAllComments(offset, limit int) ([]entities.Comment, int, error) {
	responses := []entities.Comment{
		{
			ID:     1,
			UserID: 1,
			User: entities.User{
				ID:    1,
				Name:  "test",
				Email: "test@gmail.com",
			},
			NewsID: 1,
			News: entities.News{
				ID:      1,
				Title:   "test berita",
				Content: "test content",
			},
			Content: "test content",
		},
	}

	if repo.ShouldFail {
		return []entities.Comment{}, 0, errors.New("comment tidak ditemukan")
	}

	return responses, len(responses), nil
}
func (repo CommentRepoDummy) GetCommentByID(commentID string) (entities.Comment, error) {
	if repo.ShouldFail {
		return entities.Comment{}, errors.New("comment tidak ditemukan")
	}

	return entities.Comment{
		ID:     1,
		UserID: 1,
		User: entities.User{
			ID:    1,
			Name:  "test",
			Email: "test@gmail.com",
		},
		NewsID: 1,
		News: entities.News{
			ID:      1,
			Title:   "test berita",
			Content: "test content",
		},
		Content: "test content",
	}, nil
}
func (repo CommentRepoDummy) DeleteComments(commentIDs []int) error {
	if repo.ShouldFailDelete {
		return errors.New("gagal menghapus komentar")
	}

	return nil
}
func (repo CommentRepoDummy) ValidateCommentIDs(commentIDs []int) ([]int, error) {
	if repo.ShouldFail {
		return []int{}, errors.New("comment tidak ditemukan")
	}

	ids := []int{1, 2, 3}
	return ids, nil
}
func (repo CommentRepoDummy) GetCommentsByNewsID(newsID, offset, limit int) ([]entities.Comment, int, error) {
	responses := []entities.Comment{
		{
			ID:     1,
			UserID: 1,
			User: entities.User{
				ID:    1,
				Name:  "test",
				Email: "test@gmail.com",
			},
			NewsID: 1,
			News: entities.News{
				ID:      1,
				Title:   "test berita",
				Content: "test content",
			},
			Content: "test content",
		},
	}
	if repo.ShouldFail {
		return []entities.Comment{}, 0, errors.New("gagal menemukan komentar")
	}

	return responses, len(responses), nil
}

func setupTestService() {
	repo := CommentRepoDummy{}
	commentService = *NewCommentService(repo)
}

func TestCommentService_AddComment(t *testing.T) {
	setupTestService()

	t.Run("sukses", func(t *testing.T) {
		// Data dummy untuk pengujian
		comment, err := commentService.AddComment(entities.Comment{
			ID:     1,
			UserID: 1,
			User: entities.User{
				ID:    1,
				Name:  "test",
				Email: "test@gmail.com",
			},
			NewsID: 1,
			News: entities.News{
				ID:      1,
				Title:   "test berita",
				Content: "test content",
			},
			Content: "test content",
		})

		// Periksa apakah error tidak terjadi
		assert.NoError(t, err)
		assert.Equal(t, 1, comment.ID)
		assert.Equal(t, "test content", comment.Content)
		assert.Equal(t, 1, comment.NewsID)
		assert.Equal(t, 1, comment.UserID)
	})

	t.Run("berita tidak valid", func(t *testing.T) {
		// Data dummy untuk pengujian
		_, err := commentService.AddComment(entities.Comment{
			ID:     1,
			UserID: 1,
			User: entities.User{
				ID:    1,
				Name:  "test",
				Email: "test@gmail.com",
			},
			News: entities.News{
				ID:      1,
				Title:   "test berita",
				Content: "test content",
			},
			Content: "test content",
		})

		// Periksa apakah error tidak terjadi
		assert.Error(t, err)
		assert.Equal(t, "Berita Tidak Valid. Silakan Pilih Berita Yang Sesuai", err.Error())
	})

	t.Run("konten komentar tidak boleh kosong", func(t *testing.T) {
		// Data dummy untuk pengujian
		_, err := commentService.AddComment(entities.Comment{
			ID:     1,
			UserID: 1,
			User: entities.User{
				ID:    1,
				Name:  "test",
				Email: "test@gmail.com",
			},
			NewsID: 1,
			News: entities.News{
				ID:      1,
				Title:   "test berita",
				Content: "test content",
			},
		})

		// Periksa apakah error tidak terjadi
		assert.Error(t, err)
		assert.Equal(t, "Konten Komentar Tidak Boleh Kosong", err.Error())
	})

	t.Run("terjadi kesalahan saat memeriksa kategori", func(t *testing.T) {
		CommentRepo := CommentRepoDummy{
			ShouldFail: true,
		}
		commentService := NewCommentService(CommentRepo)

		// Data dummy untuk pengujian
		_, err := commentService.AddComment(entities.Comment{
			ID:     1,
			UserID: 1,
			User: entities.User{
				ID:    1,
				Name:  "test",
				Email: "test@gmail.com",
			},
			NewsID: 1,
			News: entities.News{
				ID:      1,
				Title:   "test berita",
				Content: "test content",
			},
			Content: "test content",
		})

		// Periksa apakah error tidak terjadi
		assert.Error(t, err)
		assert.Equal(t, "Terjadi Kesalahan Saat Memeriksa Kategori", err.Error())
	})

	t.Run("terjadi kesalahan saat memeriksa kategori", func(t *testing.T) {
		CommentRepo := CommentRepoDummy{
			ShouldFail: true,
		}
		commentService := NewCommentService(CommentRepo)

		// Data dummy untuk pengujian
		_, err := commentService.AddComment(entities.Comment{
			ID:     1,
			UserID: 1,
			User: entities.User{
				ID:    1,
				Name:  "test",
				Email: "test@gmail.com",
			},
			NewsID: 1,
			News: entities.News{
				ID:      1,
				Title:   "test berita",
				Content: "test content",
			},
			Content: "test content",
		})

		// Periksa apakah error tidak terjadi
		assert.Error(t, err)
		assert.Equal(t, "Terjadi Kesalahan Saat Memeriksa Kategori", err.Error())
	})

	t.Run("berita tidak ditemukan", func(t *testing.T) {
		CommentRepo := CommentRepoDummy{
			ShouldFail: false,
		}
		commentService := NewCommentService(CommentRepo)

		// Data dummy untuk pengujian
		_, err := commentService.AddComment(entities.Comment{
			ID:     1,
			UserID: 1,
			User: entities.User{
				ID:    1,
				Name:  "test",
				Email: "test@gmail.com",
			},
			NewsID: 2,
			News: entities.News{
				ID:      1,
				Title:   "test berita",
				Content: "test content",
			},
			Content: "test content",
		})

		// Periksa apakah error tidak terjadi
		assert.Error(t, err)
		assert.Equal(t, "Berita Tidak Ditemukan", err.Error())
	})
}

func TestCommentService_GetCommentsByUserID(t *testing.T) {
	setupTestService()

	t.Run("sukses", func(t *testing.T) {
		// Data dummy untuk pengujian
		comment, total, err := commentService.GetCommentsByUserID(1, 0, 1)

		assert.NoError(t, err)
		assert.Equal(t, 1, total)
		assert.Equal(t, "test content", comment[0].Content)
		assert.Equal(t, 1, comment[0].NewsID)
		assert.Equal(t, 1, comment[0].UserID)
	})

	t.Run("user tidak ditemukan", func(t *testing.T) {
		repo := CommentRepoDummy{
			ShouldFail: true,
		}
		commentService := NewCommentService(repo)

		// Data dummy untuk pengujian
		_, _, err := commentService.GetCommentsByUserID(1, 0, 1)

		assert.Error(t, err)
		assert.Equal(t, "user tidak ditemukan", err.Error())
	})
}

func TestCommentService_GetAllComments(t *testing.T) {
	setupTestService()

	t.Run("sukses", func(t *testing.T) {
		// Data dummy untuk pengujian
		comment, total, err := commentService.GetAllComments(0, 1)

		assert.NoError(t, err)
		assert.Equal(t, 1, total)
		assert.Equal(t, "test content", comment[0].Content)
		assert.Equal(t, 1, comment[0].NewsID)
		assert.Equal(t, 1, comment[0].UserID)
	})

	t.Run("comment tidak ditemukan", func(t *testing.T) {
		repo := CommentRepoDummy{
			ShouldFail: true,
		}
		commentService := NewCommentService(repo)

		// Data dummy untuk pengujian
		_, _, err := commentService.GetAllComments(0, 1)

		assert.Error(t, err)
		assert.Equal(t, "comment tidak ditemukan", err.Error())
	})
}

func TestCommentService_GetCommentByID(t *testing.T) {
	setupTestService()

	t.Run("sukses", func(t *testing.T) {
		// Data dummy untuk pengujian
		comment, err := commentService.GetCommentByID("1")

		assert.NoError(t, err)
		assert.Equal(t, "test content", comment.Content)
		assert.Equal(t, 1, comment.NewsID)
		assert.Equal(t, 1, comment.UserID)
	})

	t.Run("comment tidak ditemukan", func(t *testing.T) {
		repo := CommentRepoDummy{
			ShouldFail: true,
		}
		commentService := NewCommentService(repo)

		// Data dummy untuk pengujian
		_, err := commentService.GetCommentByID("1")

		assert.Error(t, err)
		assert.Equal(t, "comment tidak ditemukan", err.Error())
	})
}
