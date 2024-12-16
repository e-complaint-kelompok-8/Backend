package feedbacks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFeedbackService_ProvideFeedback(t *testing.T) {
	setupTestService()

	t.Run("Admin tidak ditemukan", func(t *testing.T) {
		_, err := feedbackService.ProvideFeedback(99, 1, "Feedback content")
		assert.Error(t, err)
		assert.Equal(t, "Admin Tidak Ditemukan", err.Error())
	})

	t.Run("Pengaduan tidak ditemukan", func(t *testing.T) {
		repo := FeedbackRepoDummy{
			ShouldFail: true, // Simulasi feedback sudah ada
		}
		feedbackService = *NewFeedbackService(repo)

		_, err := feedbackService.ProvideFeedback(1, 999, "Feedback content")
		assert.Error(t, err)
		assert.Equal(t, "Pengaduan Tidak Ditemukan", err.Error())
	})

	t.Run("Pengaduan sudah memiliki feedback", func(t *testing.T) {
		// Ubah behavior mock agar `ComplaintHasFeedback` mengembalikan true
		repo := FeedbackRepoDummy{
			ShouldFail: true, // Simulasi feedback sudah ada
		}
		feedbackService = *NewFeedbackService(repo)

		_, err := feedbackService.ProvideFeedback(1, 3, "Feedback content")
		assert.Error(t, err)
		assert.Equal(t, "Pengaduan Sudah Memiliki Tanggapan", err.Error())
	})

	t.Run("User tidak ditemukan", func(t *testing.T) {
		_, err := feedbackService.ProvideFeedback(1, 4, "Feedback content")
		assert.Error(t, err)
		assert.Equal(t, "Pengguna Tidak Ditemukan", err.Error())
	})

	t.Run("Gagal menyimpan feedback", func(t *testing.T) {
		// Ubah mock agar `CreateFeedback` mengembalikan error
		repo := FeedbackRepoDummy{
			ShouldFailUpdate: true, // Simulasi error saat menyimpan feedback
		}
		feedbackService = *NewFeedbackService(repo)

		_, err := feedbackService.ProvideFeedback(1, 2, "Feedback content")
		assert.Error(t, err)
	})

	t.Run("Gagal memperbarui status pengaduan", func(t *testing.T) {
		// Ubah mock agar `AdminUpdateComplaintStatus` mengembalikan error
		repo := FeedbackRepoDummy{
			ShouldFailStatus: true, // Simulasi error saat memperbarui status
		}
		feedbackService = *NewFeedbackService(repo)

		_, err := feedbackService.ProvideFeedback(1, 2, "Feedback content")
		assert.Error(t, err)
		assert.Equal(t, "Gagal Memperbarui Status Pengaduan", err.Error())
	})

	t.Run("gagal memeriksa feedback pengaduan", func(t *testing.T) {
		repo := FeedbackRepoDummy{
			ShouldFailHasFeedback: true, // Simulasi error saat memperbarui status
		}
		feedbackService = *NewFeedbackService(repo)

		_, err := feedbackService.ProvideFeedback(1, 2, "Feedback content")
		assert.Error(t, err)
		assert.Equal(t, "Gagal Memeriksa Feedback Pengaduan", err.Error())
	})

	t.Run("gagal mengambil masukan lengkap", func(t *testing.T) {
		repo := FeedbackRepoDummy{
			ShouldFailGetFeedback: true, // Simulasi error saat memperbarui status
		}
		feedbackService = *NewFeedbackService(repo)

		_, err := feedbackService.ProvideFeedback(1, 2, "Feedback content")
		assert.Error(t, err)
		assert.Equal(t, "Gagal Mengambil Masukan Lengkap", err.Error())
	})

	t.Run("Sukses membuat feedback", func(t *testing.T) {
		// Ubah mock agar semua proses berhasil
		repo := FeedbackRepoDummy{}
		feedbackService = *NewFeedbackService(repo)

		feedback, err := feedbackService.ProvideFeedback(1, 2, "test content")
		assert.NoError(t, err)
		assert.Equal(t, 2, feedback.ID)
		assert.Equal(t, "test content", feedback.Content)
	})
}

func TestFeedbackService_UpdateFeedback(t *testing.T) {
	setupTestService()

	t.Run("tanggapan tidak ditemukan", func(t *testing.T) {
		_, err := feedbackService.UpdateFeedback(99, "Feedback content")
		assert.Error(t, err)
		assert.Equal(t, "Tanggapan Tidak Ditemukan", err.Error())
	})

	t.Run("gagal memperbarui tanggapan", func(t *testing.T) {
		repo := FeedbackRepoDummy{
			ShouldFailUpdate: true,
		}
		feedbackService = *NewFeedbackService(repo)

		_, err := feedbackService.UpdateFeedback(2, "Feedback content")
		assert.Error(t, err)
		assert.Equal(t, "Gagal Memperbarui Tanggapan", err.Error())
	})
}
