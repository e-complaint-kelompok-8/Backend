package feedbacks

import (
	"capstone/entities"
	"capstone/utils"
	"errors"
	"time"
)

func (cs *FeedbackService) ProvideFeedback(adminID, complaintID int, content string) (entities.Feedback, error) {
	// Periksa apakah admin ID valid
	adminExists, err := cs.feedbackRepo.CheckAdminExists(adminID)
	if err != nil || !adminExists {
		return entities.Feedback{}, errors.New(utils.CapitalizeErrorMessage(errors.New("admin tidak ditemukan")))
	}

	// Periksa apakah complaint ID valid
	complaint, err := cs.feedbackRepo.GetComplaintByID(complaintID)
	if err != nil {
		return entities.Feedback{}, errors.New(utils.CapitalizeErrorMessage(errors.New("pengaduan tidak ditemukan")))
	}

	// Periksa apakah complaint sudah memiliki feedback
	hasFeedback, err := cs.feedbackRepo.ComplaintHasFeedback(complaintID)
	if err != nil {
		return entities.Feedback{}, errors.New(utils.CapitalizeErrorMessage(errors.New("gagal memeriksa feedback pengaduan")))
	}
	if hasFeedback {
		return entities.Feedback{}, errors.New(utils.CapitalizeErrorMessage(errors.New("pengaduan sudah memiliki tanggapan")))
	}

	// Pastikan user ID valid
	userExists, err := cs.feedbackRepo.CheckUserExists(complaint.UserID)
	if err != nil || !userExists {
		return entities.Feedback{}, errors.New(utils.CapitalizeErrorMessage(errors.New("pengguna tidak ditemukan")))
	}

	// Buat feedback entity
	feedback := entities.Feedback{
		AdminID:     adminID,
		ComplaintID: complaintID,
		UserID:      complaint.UserID,
		Content:     content,
		CreatedAt:   time.Now(),
	}

	// Simpan feedback di repository
	err = cs.feedbackRepo.CreateFeedback(&feedback)
	if err != nil {
		return entities.Feedback{}, err
	}

	// Perbarui status complaint menjadi "tanggapi"
	err = cs.feedbackRepo.AdminUpdateComplaintStatus(complaintID, "tanggapi", adminID)
	if err != nil {
		return entities.Feedback{}, errors.New(utils.CapitalizeErrorMessage(errors.New("gagal memperbarui status pengaduan")))
	}

	// Ambil feedback lengkap dengan relasi
	feedbackModel, err := cs.feedbackRepo.GetFeedbackByComplaintID(complaintID)
	if err != nil {
		return entities.Feedback{}, errors.New(utils.CapitalizeErrorMessage(errors.New("gagal mengambil masukan lengkap")))
	}

	// Gabungkan feedback dan complaint yang diperbarui
	feedback = feedbackModel

	return feedback, nil
}

func (cs *FeedbackService) UpdateFeedback(feedbackID int, content string) (entities.Feedback, error) {
	// Periksa apakah feedback ID valid
	feedback, err := cs.feedbackRepo.GetFeedbackByID(feedbackID)
	if err != nil {
		return entities.Feedback{}, errors.New(utils.CapitalizeErrorMessage(errors.New("tanggapan tidak ditemukan")))
	}

	// Perbarui konten
	feedback.Content = content

	// Simpan perubahan ke database
	err = cs.feedbackRepo.UpdateFeedback(feedback)
	if err != nil {
		return entities.Feedback{}, errors.New(utils.CapitalizeErrorMessage(errors.New("gagal memperbarui tanggapan")))
	}

	// Ambil feedback yang diperbarui
	updatedFeedback, err := cs.feedbackRepo.GetFeedbackByID(feedbackID)
	if err != nil {
		return entities.Feedback{}, errors.New(utils.CapitalizeErrorMessage(errors.New("gagal mengambil tanggapan yang diperbarui")))
	}

	return updatedFeedback, nil
}
