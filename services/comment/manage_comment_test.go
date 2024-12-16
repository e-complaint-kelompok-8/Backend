package comment

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommentService_DeleteComments(t *testing.T) {
	setupTestService()

	t.Run("sukses", func(t *testing.T) {
		// Data dummy untuk pengujian
		ids := []int{1, 2, 3}
		err := commentService.DeleteComments(ids)

		assert.NoError(t, err)
	})

	t.Run("comment tidak ditemukan", func(t *testing.T) {
		repo := CommentRepoDummy{
			ShouldFail: true,
		}
		commentService := NewCommentService(repo)

		// Data dummy untuk pengujian
		ids := []int{1, 2, 3}
		err := commentService.DeleteComments(ids)

		assert.Error(t, err)
		assert.Equal(t, "comment tidak ditemukan", err.Error())
	})

	t.Run("some comment IDs do not exist", func(t *testing.T) {
		repo := CommentRepoDummy{
			ShouldFail: false,
		}
		commentService := NewCommentService(repo)

		// Data dummy untuk pengujian
		ids := []int{1, 2}
		err := commentService.DeleteComments(ids)

		assert.Error(t, err)
		assert.Equal(t, "some comment IDs do not exist", err.Error())
	})

	t.Run("gagal menghapus komentar", func(t *testing.T) {
		repo := CommentRepoDummy{
			ShouldFail:       false,
			ShouldFailDelete: true,
		}
		commentService := NewCommentService(repo)

		// Data dummy untuk pengujian
		ids := []int{1, 2, 3}
		err := commentService.DeleteComments(ids)

		assert.Error(t, err)
		assert.Equal(t, "gagal menghapus komentar", err.Error())
	})
}

func TestCommentService_GetCommentsByNewsID(t *testing.T) {
	setupTestService()

	t.Run("sukses", func(t *testing.T) {
		// Data dummy untuk pengujian
		comment, total, err := commentService.GetCommentsByNewsID(1, 0, 1)

		assert.NoError(t, err)
		assert.Equal(t, 1, total)
		assert.Equal(t, "test content", comment[0].Content)
		assert.Equal(t, 1, comment[0].NewsID)
		assert.Equal(t, 1, comment[0].UserID)
	})

	t.Run("gagal menemukan komentar", func(t *testing.T) {
		repo := CommentRepoDummy{
			ShouldFail: true,
		}
		commentService := NewCommentService(repo)
		// Data dummy untuk pengujian
		_, _, err := commentService.GetCommentsByNewsID(1, 0, 1)

		assert.Error(t, err)
		assert.Equal(t, "gagal menemukan komentar", err.Error())
	})
}
