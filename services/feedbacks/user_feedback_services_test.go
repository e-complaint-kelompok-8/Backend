package feedbacks

import (
	"capstone/entities"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var feedbackService FeedbackService

type FeedbackRepoDummy struct {
	ShouldFail            bool
	ShouldFailUpdate      bool
	ShouldFailStatus      bool
	ShouldFailHasFeedback bool
	ShouldFailGetFeedback bool
	MockedFeedback        entities.Feedback
}

func (repo FeedbackRepoDummy) GetComplaintByID(complaintID int) (entities.Complaint, error) {
	// Return mock complaint sesuai dengan complaintID
	if complaintID == 1 {
		return entities.Complaint{
			ID:     1,
			UserID: 1,
			User: entities.User{
				ID:    1, // UserID sesuai dengan test case
				Name:  "test",
				Email: "test@gmail.com",
			},
			Category: entities.Category{
				ID:          1,
				Name:        "testkategori",
				Description: "ini cuman contoh",
			},
			ComplaintNumber: "KEE123",
			Title:           "test",
			Location:        "medan",
			Status:          "tanggapi",
			Description:     "cuman contoh",
			Reason:          "",
		}, nil
	}

	if complaintID == 2 {
		return entities.Complaint{
			ID:     2,
			UserID: 1,
			User: entities.User{
				ID:    1, // UserID sesuai dengan test case
				Name:  "test",
				Email: "test@gmail.com",
			},
			Category: entities.Category{
				ID:          1,
				Name:        "testkategori",
				Description: "ini cuman contoh",
			},
			ComplaintNumber: "KEE123",
			Title:           "test",
			Location:        "medan",
			Status:          "proses",
			Description:     "cuman contoh",
			Reason:          "",
		}, nil
	}

	if complaintID == 3 {
		return entities.Complaint{
			ID:     3,
			UserID: 1,
			User: entities.User{
				ID:    1, // UserID sesuai dengan test case
				Name:  "test",
				Email: "test@gmail.com",
			},
			Category: entities.Category{
				ID:          1,
				Name:        "testkategori",
				Description: "ini cuman contoh",
			},
			ComplaintNumber: "KEE123",
			Title:           "test",
			Location:        "medan",
			Status:          "proses",
			Description:     "cuman contoh",
			Reason:          "",
			Feedbacks: []entities.Feedback{
				{
					ID: 2,
					Admin: entities.Admin{
						ID:    1,
						Email: "test@gmail.com",
						Role:  "admin",
					},
					UserID: 1,
					User: entities.User{
						ID:    1,
						Name:  "test",
						Email: "test@gmail.com",
					},
					Complaint: entities.Complaint{
						ID: 1,
						User: entities.User{
							ID:    1,
							Name:  "test",
							Email: "test@gmail.com",
						},
						Category: entities.Category{
							ID:          1,
							Name:        "testkategori",
							Description: "ini cuman contoh",
						},
						ComplaintNumber: "KEE123",
						Title:           "test",
						Location:        "medan",
						Status:          "proses",
						Description:     "cuman contoh",
						Reason:          "",
					},
					Content:  "test content",
					Response: "test response",
				},
			},
		}, nil
	}

	if complaintID == 4 {
		return entities.Complaint{
			ID:     4,
			UserID: 999,
			User: entities.User{
				ID:    999, // UserID sesuai dengan test case
				Name:  "test",
				Email: "test@gmail.com",
			},
			Category: entities.Category{
				ID:          1,
				Name:        "testkategori",
				Description: "ini cuman contoh",
			},
			ComplaintNumber: "KEE123",
			Title:           "test",
			Location:        "medan",
			Status:          "proses",
			Description:     "cuman contoh",
			Reason:          "",
			Feedbacks: []entities.Feedback{
				{
					ID: 2,
					Admin: entities.Admin{
						ID:    1,
						Email: "test@gmail.com",
						Role:  "admin",
					},
					UserID: 1,
					User: entities.User{
						ID:    1,
						Name:  "test",
						Email: "test@gmail.com",
					},
					Complaint: entities.Complaint{
						ID: 1,
						User: entities.User{
							ID:    1,
							Name:  "test",
							Email: "test@gmail.com",
						},
						Category: entities.Category{
							ID:          1,
							Name:        "testkategori",
							Description: "ini cuman contoh",
						},
						ComplaintNumber: "KEE123",
						Title:           "test",
						Location:        "medan",
						Status:          "proses",
						Description:     "cuman contoh",
						Reason:          "",
					},
				},
			},
		}, nil
	}
	return entities.Complaint{}, errors.New("pengaduan tidak ditemukan")
}
func (repo FeedbackRepoDummy) GetFeedbackByComplaintID(complaintID int) (entities.Feedback, error) {
	if repo.ShouldFailGetFeedback {
		return entities.Feedback{}, errors.New("Gagal Mengambil Masukan Lengkap")
	}

	// Return feedback hanya jika complaintID valid
	if complaintID == 1 {
		return entities.Feedback{
			ID: 1,
			Admin: entities.Admin{
				ID:    1,
				Email: "test@gmail.com",
				Role:  "admin",
			},
			User: entities.User{
				ID:    1,
				Name:  "test",
				Email: "test@gmail.com",
			},
			Complaint: entities.Complaint{
				ID: 1,
				User: entities.User{
					ID:    1,
					Name:  "test",
					Email: "test@gmail.com",
				},
				Category: entities.Category{
					ID:          1,
					Name:        "testkategori",
					Description: "ini cuman contoh",
				},
				ComplaintNumber: "KEE123",
				Title:           "test",
				Location:        "medan",
				Status:          "tanggapi",
				Description:     "cuman contoh",
				Reason:          "",
			},
			Content:  "test content",
			Response: "test response",
		}, nil
	}

	if complaintID == 2 {
		return entities.Feedback{
			ID: 2,
			Admin: entities.Admin{
				ID:    1,
				Email: "test@gmail.com",
				Role:  "admin",
			},
			User: entities.User{
				ID:    1,
				Name:  "test",
				Email: "test@gmail.com",
			},
			Complaint: entities.Complaint{
				ID: 1,
				User: entities.User{
					ID:    1,
					Name:  "test",
					Email: "test@gmail.com",
				},
				Category: entities.Category{
					ID:          1,
					Name:        "testkategori",
					Description: "ini cuman contoh",
				},
				ComplaintNumber: "KEE123",
				Title:           "test",
				Location:        "medan",
				Status:          "tanggapi",
				Description:     "cuman contoh",
				Reason:          "",
			},
			Content:  "test content",
			Response: "test response",
		}, nil
	}

	if complaintID == 3 {
		return entities.Feedback{
			ID: 3,
			Admin: entities.Admin{
				ID:    1,
				Email: "test@gmail.com",
				Role:  "admin",
			},
			User: entities.User{
				ID:    1,
				Name:  "test",
				Email: "test@gmail.com",
			},
			Complaint: entities.Complaint{
				ID: 1,
				User: entities.User{
					ID:    1,
					Name:  "test",
					Email: "test@gmail.com",
				},
				Category: entities.Category{
					ID:          1,
					Name:        "testkategori",
					Description: "ini cuman contoh",
				},
				ComplaintNumber: "KEE123",
				Title:           "test",
				Location:        "medan",
				Status:          "tanggapi",
				Description:     "cuman contoh",
				Reason:          "",
			},
		}, errors.New("Tanggapan Tidak Ditemukan")
	}
	return entities.Feedback{}, errors.New("tanggapan tidak ditemukan") // Mock error
}
func (repo FeedbackRepoDummy) GetFeedbacksByUserID(userID int) ([]entities.Feedback, error) {
	responses := []entities.Feedback{
		{
			ID: 1,
			Admin: entities.Admin{
				ID:    1,
				Email: "test@gmail.com",
				Role:  "admin",
			},
			User: entities.User{
				ID:    1,
				Name:  "test",
				Email: "test@gmail.com",
			},
			Complaint: entities.Complaint{
				ID: 1,
				User: entities.User{
					ID:    1,
					Name:  "test",
					Email: "test@gmail.com",
				},
				Category: entities.Category{
					ID:          1,
					Name:        "testkategori",
					Description: "ini cuman contoh",
				},
				ComplaintNumber: "KEE123",
				Title:           "test",
				Location:        "medan",
				Status:          "proses",
				Description:     "cuman contoh",
				Reason:          "",
			},
			Content:  "test content",
			Response: "test response",
		}, {ID: 2,
			Admin: entities.Admin{
				ID:    1,
				Email: "test@gmail.com",
				Role:  "admin",
			},
			User: entities.User{
				ID:    1,
				Name:  "test",
				Email: "test@gmail.com",
			},
			Complaint: entities.Complaint{
				ID: 2,
				User: entities.User{
					ID:    1,
					Name:  "test 2",
					Email: "test@gmail.com",
				},
				Category: entities.Category{
					ID:          1,
					Name:        "testkategori",
					Description: "ini cuman contoh",
				},
				ComplaintNumber: "KEE123",
				Title:           "test",
				Location:        "medan",
				Status:          "proses",
				Description:     "cuman contoh",
				Reason:          "",
			},
			Content:  "test content 2",
			Response: "test response",
		},
	}

	if repo.ShouldFail {
		return []entities.Feedback{}, errors.New("feedback tidak ditemukan")
	}

	return responses, nil
}
func (repo FeedbackRepoDummy) UpdateFeedbackResponse(feedbackID int, response string) error {
	if repo.ShouldFailUpdate {
		return errors.New("gagal memberikan balasan pada tanggapan")
	}

	return nil
}
func (repo FeedbackRepoDummy) UpdateComplaintStatus(complaintID int, status string) error {
	if repo.ShouldFailStatus {
		return errors.New("gagal memperbarui status pengaduan")
	}

	return nil
}
func (repo FeedbackRepoDummy) GetFeedbackByID(feedbackID int) (entities.Feedback, error) {
	if feedbackID == 1 {
		return entities.Feedback{
			ID: 1,
			Admin: entities.Admin{
				ID:    1,
				Email: "test@gmail.com",
				Role:  "admin",
			},
			UserID: 1,
			User: entities.User{
				ID:    1,
				Name:  "test",
				Email: "test@gmail.com",
			},
			Complaint: entities.Complaint{
				ID: 1,
				User: entities.User{
					ID:    1,
					Name:  "test",
					Email: "test@gmail.com",
				},
				Category: entities.Category{
					ID:          1,
					Name:        "testkategori",
					Description: "ini cuman contoh",
				},
				ComplaintNumber: "KEE123",
				Title:           "test",
				Location:        "medan",
				Status:          "tanggapi",
				Description:     "cuman contoh",
				Reason:          "",
			},
		}, nil
	}

	if feedbackID == 2 {
		return entities.Feedback{
			ID: 2,
			Admin: entities.Admin{
				ID:    1,
				Email: "test@gmail.com",
				Role:  "admin",
			},
			UserID: 1,
			User: entities.User{
				ID:    1,
				Name:  "test",
				Email: "test@gmail.com",
			},
			Complaint: entities.Complaint{
				ID: 1,
				User: entities.User{
					ID:    1,
					Name:  "test",
					Email: "test@gmail.com",
				},
				Category: entities.Category{
					ID:          1,
					Name:        "testkategori",
					Description: "ini cuman contoh",
				},
				ComplaintNumber: "KEE123",
				Title:           "test",
				Location:        "medan",
				Status:          "proses",
				Description:     "cuman contoh",
				Reason:          "",
			},
		}, nil
	}

	if feedbackID == 3 {
		return entities.Feedback{
			ID: 2,
			Admin: entities.Admin{
				ID:    1,
				Email: "test@gmail.com",
				Role:  "admin",
			},
			UserID: 1,
			User: entities.User{
				ID:    1,
				Name:  "test",
				Email: "test@gmail.com",
			},
			Complaint: entities.Complaint{
				ID: 1,
				User: entities.User{
					ID:    1,
					Name:  "test",
					Email: "test@gmail.com",
				},
				Category: entities.Category{
					ID:          1,
					Name:        "testkategori",
					Description: "ini cuman contoh",
				},
				ComplaintNumber: "KEE123",
				Title:           "test",
				Location:        "medan",
				Status:          "tanggapi",
				Description:     "cuman contoh",
				Reason:          "",
			},
			Content:  "test content",
			Response: "test response",
		}, nil
	}

	if feedbackID == 4 {
		return entities.Feedback{
			ID: 4,
			Admin: entities.Admin{
				ID:    1,
				Email: "test@gmail.com",
				Role:  "admin",
			},
			UserID: 1,
			User: entities.User{
				ID:    1,
				Name:  "test",
				Email: "test@gmail.com",
			},
			Complaint: entities.Complaint{
				ID: 1,
				User: entities.User{
					ID:    1,
					Name:  "test",
					Email: "test@gmail.com",
				},
				Category: entities.Category{
					ID:          1,
					Name:        "testkategori",
					Description: "ini cuman contoh",
				},
				ComplaintNumber: "KEE123",
				Title:           "test",
				Location:        "medan",
				Status:          "tanggapi",
				Description:     "cuman contoh",
				Reason:          "",
			},
			Content: "test content",
		}, nil
	}

	return entities.Feedback{}, errors.New("tanggapan tidak ditemukan")
}
func (repo FeedbackRepoDummy) CreateFeedback(feedback *entities.Feedback) error {
	if repo.ShouldFailUpdate {
		return errors.New("Gagal menyimpan feedback")
	}

	return nil
}
func (repo FeedbackRepoDummy) CheckAdminExists(adminID int) (bool, error) {
	if adminID == 1 {
		return true, nil
	}
	return false, nil
}

func (repo FeedbackRepoDummy) CheckUserExists(userID int) (bool, error) {
	if userID == 1 {
		return true, nil
	}
	return false, errors.New("Pengguna Tidak Ditemukan")
}
func (repo FeedbackRepoDummy) AdminUpdateComplaintStatus(complaintID int, newStatus string, adminID int) error {
	if repo.ShouldFailStatus {
		return errors.New("gagal memperbarui status pengaduan")
	}
	return nil
}
func (repo FeedbackRepoDummy) ComplaintHasFeedback(complaintID int) (bool, error) {
	if repo.ShouldFailHasFeedback {
		return false, errors.New("gagal memeriksa feedback pengaduan")
	}

	if complaintID == 2 {
		return false, nil
	}

	if complaintID == 4 {
		return false, nil
	}
	return true, nil
}
func (repo FeedbackRepoDummy) UpdateFeedback(feedback entities.Feedback) error {
	if repo.ShouldFailUpdate {
		return errors.New("gagal memperbarui tanggapan")
	}

	return nil
}

func setupTestService() {
	repo := FeedbackRepoDummy{}
	feedbackService = *NewFeedbackService(repo)
}

func TestFeedbackService_GetFeedbackByComplaint(t *testing.T) {
	setupTestService()

	t.Run("sukses", func(t *testing.T) {
		// Data dummy untuk pengujian
		feedback, err := feedbackService.GetFeedbackByComplaint(1, 1)

		// Periksa apakah error tidak terjadi
		assert.NoError(t, err)
		assert.Equal(t, 1, feedback.ID)
		assert.Equal(t, "test content", feedback.Content)
	})

	t.Run("gagal - pengaduan tidak ditemukan", func(t *testing.T) {
		_, err := feedbackService.GetFeedbackByComplaint(999, 1) // Invalid complaintID

		assert.Error(t, err)
		assert.Equal(t, "Pengaduan Tidak Ditemukan", err.Error())
	})

	t.Run("gagal - tidak berwenang", func(t *testing.T) {
		_, err := feedbackService.GetFeedbackByComplaint(1, 999) // Valid complaintID tetapi userID tidak cocok

		assert.Error(t, err)
		assert.Equal(t, "Tidak Berwenang Untuk Melihat Tanggapan Ini", err.Error())
	})

	t.Run("gagal - tanggapan tidak ditemukan", func(t *testing.T) {
		_, err := feedbackService.GetFeedbackByComplaint(3, 1) // Valid complaintID tetapi feedback tidak ditemukan

		assert.Error(t, err)
		assert.Equal(t, "Tanggapan Tidak Ditemukan", err.Error())
	})
}

func TestFeedbackService_GetFeedbacksByUser(t *testing.T) {
	setupTestService()

	t.Run("sukses", func(t *testing.T) {
		// Mock sukses
		newsRepo := FeedbackRepoDummy{
			ShouldFail: false,
		}

		feedbackService := NewFeedbackService(newsRepo)
		// Data dummy untuk pengujian
		feedback, err := feedbackService.GetFeedbacksByUser(1)

		// Periksa apakah error tidak terjadi
		assert.NoError(t, err)
		assert.Equal(t, 1, feedback[0].ID)
		assert.Equal(t, "test content", feedback[0].Content)
		assert.Equal(t, 2, feedback[1].ID)
		assert.Equal(t, "test content 2", feedback[1].Content)
	})

	t.Run("gagal - feedback tidak ditemukan", func(t *testing.T) {
		// Mock sukses
		feedbackRepo := FeedbackRepoDummy{
			ShouldFail: true,
		}

		feedbackService := NewFeedbackService(feedbackRepo)
		// Data dummy untuk pengujian
		_, err := feedbackService.GetFeedbacksByUser(1)

		// Periksa apakah error tidak terjadi
		assert.Error(t, err)
		assert.Equal(t, "feedback tidak ditemukan", err.Error())
	})
}

func TestFeedbackService_AddResponseToFeedback(t *testing.T) {
	setupTestService()

	t.Run("sukses", func(t *testing.T) {
		// Data dummy untuk pengujian
		err := feedbackService.AddResponseToFeedback(1, 1, "terima kasih")

		// Periksa apakah error tidak terjadi
		assert.NoError(t, err)
	})

	t.Run("gagal - tanggapan tidak ditemukan", func(t *testing.T) {
		// Data dummy untuk pengujian
		err := feedbackService.AddResponseToFeedback(999, 1, "terima kasih")

		// Periksa apakah error tidak terjadi
		assert.Error(t, err)
		assert.Equal(t, "Tanggapan Tidak Ditemukan", err.Error())
	})

	t.Run("gagal - anda tidak memiliki akses untuk memberikan balasan pada tanggapan ini", func(t *testing.T) {
		// Data dummy untuk pengujian
		err := feedbackService.AddResponseToFeedback(1, 2, "terima kasih")

		// Periksa apakah error tidak terjadi
		assert.Error(t, err)
		assert.Equal(t, "Anda Tidak Memiliki Akses Untuk Memberikan Balasan Pada Tanggapan Ini", err.Error())
	})

	t.Run("gagal - tanggapan ini sudah memiliki balasan", func(t *testing.T) {
		// Data dummy untuk pengujian
		err := feedbackService.AddResponseToFeedback(3, 1, "terima kasih")

		// Periksa apakah error tidak terjadi
		assert.Error(t, err)
		assert.Equal(t, "Tanggapan Ini Sudah Memiliki Balasan", err.Error())
	})

	t.Run("gagal - tanggapan tidak dapat diberikan komentar", func(t *testing.T) {
		// Data dummy untuk pengujian
		err := feedbackService.AddResponseToFeedback(2, 1, "terima kasih")

		// Periksa apakah error tidak terjadi
		assert.Error(t, err)
		assert.Equal(t, "Tanggapan Tidak Dapat Diberikan Komentar", err.Error())
	})

	t.Run("gagal - gagal memberikan balasan pada tanggapan", func(t *testing.T) {
		// Mock sukses
		feedbackRepo := FeedbackRepoDummy{
			ShouldFailUpdate: true,
		}
		feedbackService := NewFeedbackService(feedbackRepo)

		// Data dummy untuk pengujian
		err := feedbackService.AddResponseToFeedback(1, 1, "terima kasih")

		// Periksa apakah error tidak terjadi
		assert.Error(t, err)
		assert.Equal(t, "Gagal Memberikan Balasan Pada Tanggapan", err.Error())
	})

	t.Run("gagal - gagal memperbarui status pengaduan", func(t *testing.T) {
		// Mock sukses
		feedbackRepo := FeedbackRepoDummy{
			ShouldFailStatus: true,
		}
		feedbackService := NewFeedbackService(feedbackRepo)

		// Data dummy untuk pengujian
		err := feedbackService.AddResponseToFeedback(1, 1, "terima kasih")

		// Periksa apakah error tidak terjadi
		assert.Error(t, err)
		assert.Equal(t, "Gagal Memperbarui Status Pengaduan", err.Error())
	})
}

func TestFeedbackService_GetFeedbackByID(t *testing.T) {
	setupTestService()

	t.Run("sukses", func(t *testing.T) {
		// Data dummy untuk pengujian
		feedback, err := feedbackService.GetFeedbackByID(4, 1)

		// Periksa apakah error tidak terjadi
		assert.NoError(t, err)
		assert.Equal(t, 4, feedback.ID)
		assert.Equal(t, "test content", feedback.Content)
	})

	t.Run("gagal - Feedback tidak ditemukan", func(t *testing.T) {
		_, err := feedbackService.GetFeedbackByID(999, 1) // Invalid complaintID

		assert.Error(t, err)
		assert.Equal(t, "Feedback Tidak Ditemukan", err.Error())
	})

	t.Run("gagal - tidak berwenang", func(t *testing.T) {
		_, err := feedbackService.GetFeedbackByID(1, 999) // Valid complaintID tetapi userID tidak cocok

		assert.Error(t, err)
		assert.Equal(t, "Anda Tidak Memiliki Akses Untuk Melihat Feedback Ini", err.Error())
	})
}
