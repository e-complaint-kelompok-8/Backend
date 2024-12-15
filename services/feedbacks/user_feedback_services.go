package feedbacks

import (
	"capstone/entities"
	feedback "capstone/repositories/feedbacks"
	"capstone/utils"
	"errors"
)

type FeedbackServiceInterface interface {
	GetFeedbackByComplaint(complaintID int, userID int) (entities.Feedback, error)
	GetFeedbacksByUser(userID int) ([]entities.Feedback, error)
	AddResponseToFeedback(feedbackID int, userID int, response string) error
	GetFeedbackByID(feedbackID int, userID int) (entities.Feedback, error)
	ProvideFeedback(adminID, complaintID int, content string) (entities.Feedback, error)
	UpdateFeedback(feedbackID int, content string) (entities.Feedback, error)
}

type FeedbackService struct {
	feedbackRepo feedback.FeedbackRepositoryInterface
}

func NewFeedbackService(repo feedback.FeedbackRepositoryInterface) *FeedbackService {
	return &FeedbackService{feedbackRepo: repo}
}

func (fs *FeedbackService) GetFeedbackByComplaint(complaintID int, userID int) (entities.Feedback, error) {
	// Validasi apakah complaint dimiliki oleh user dan statusnya 'tanggapi'
	complaint, err := fs.feedbackRepo.GetComplaintByID(complaintID)
	if err != nil {
		return entities.Feedback{}, errors.New(utils.CapitalizeErrorMessage(errors.New("pengaduan tidak ditemukan")))
	}

	if complaint.UserID != userID {
		return entities.Feedback{}, errors.New(utils.CapitalizeErrorMessage(errors.New("tidak berwenang untuk melihat tanggapan ini")))
	}

	// Ambil feedback
	feedback, err := fs.feedbackRepo.GetFeedbackByComplaintID(complaintID)
	if err != nil {
		return entities.Feedback{}, errors.New(utils.CapitalizeErrorMessage(errors.New("tanggapan tidak ditemukan")))
	}

	return feedback, nil
}

func (fs *FeedbackService) GetFeedbacksByUser(userID int) ([]entities.Feedback, error) {
	// Ambil semua feedback dari repository
	feedbacks, err := fs.feedbackRepo.GetFeedbacksByUserID(userID)
	if err != nil {
		return nil, err
	}
	return feedbacks, nil
}

func (fs *FeedbackService) AddResponseToFeedback(feedbackID int, userID int, response string) error {
	// Ambil feedback untuk validasi
	feedback, err := fs.feedbackRepo.GetFeedbackByID(feedbackID)
	if err != nil {
		return errors.New(utils.CapitalizeErrorMessage(errors.New("tanggapan tidak ditemukan")))
	}

	// Pastikan feedback milik user
	if feedback.UserID != userID {
		return errors.New(utils.CapitalizeErrorMessage(errors.New("anda tidak memiliki akses untuk memberikan balasan pada tanggapan ini")))
	}

	// Pastikan feedback belum memiliki balasan
	if feedback.Response != "" {
		return errors.New(utils.CapitalizeErrorMessage(errors.New("tanggapan ini sudah memiliki balasan")))
	}

	// Pastikan status complaint
	if feedback.Complaint.Status != "tanggapi" {
		return errors.New(utils.CapitalizeErrorMessage(errors.New("tanggapan tidak dapat diberikan komentar")))
	}

	// Tambahkan balasan
	err = fs.feedbackRepo.UpdateFeedbackResponse(feedbackID, response)
	if err != nil {
		return errors.New(utils.CapitalizeErrorMessage(errors.New("gagal memberikan balasan pada tanggapan")))
	}

	// Perbarui status complaint menjadi "selesai"
	err = fs.feedbackRepo.UpdateComplaintStatus(feedback.ComplaintID, "selesai")
	if err != nil {
		return errors.New(utils.CapitalizeErrorMessage(errors.New("gagal memperbarui status pengaduan")))
	}

	return nil
}

func (fs *FeedbackService) GetFeedbackByID(feedbackID int, userID int) (entities.Feedback, error) {
	// Ambil feedback dari repository
	feedback, err := fs.feedbackRepo.GetFeedbackByID(feedbackID)
	if err != nil {
		return entities.Feedback{}, errors.New(utils.CapitalizeErrorMessage(errors.New("feedback tidak ditemukan")))
	}

	// Pastikan feedback milik user
	if feedback.UserID != userID {
		return entities.Feedback{}, errors.New(utils.CapitalizeErrorMessage(errors.New("anda tidak memiliki akses untuk melihat feedback ini")))
	}

	return feedback, nil
}
