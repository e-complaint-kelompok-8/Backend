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
		return entities.Feedback{}, errors.New(utils.CapitalizeErrorMessage(errors.New("tidak berwenang untuk melihat masukan ini")))
	}

	if complaint.Status != "tanggapi" {
		return entities.Feedback{}, errors.New(utils.CapitalizeErrorMessage(errors.New("masukan tidak tersedia untuk pengaduan ini")))
	}

	// Ambil feedback
	feedback, err := fs.feedbackRepo.GetFeedbackByComplaintID(complaintID)
	if err != nil {
		return entities.Feedback{}, errors.New(utils.CapitalizeErrorMessage(errors.New("masukan tidak ditemukan")))
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
