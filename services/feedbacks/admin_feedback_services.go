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

	// Ambil feedback lengkap dari database
	feedbackModel, err := cs.feedbackRepo.GetFeedbackByID(feedback.ID)
	if err != nil {
		return entities.Feedback{}, errors.New(utils.CapitalizeErrorMessage(errors.New("gagal mengambil masukan")))
	}

	// Perbarui status complaint menjadi "tanggapi"
	err = cs.feedbackRepo.AdminUpdateComplaintStatus(complaintID, "tanggapi", adminID)
	if err != nil {
		return entities.Feedback{}, errors.New(utils.CapitalizeErrorMessage(errors.New("gagal memperbarui status pengaduan")))
	}

	return feedbackModel, nil
}
