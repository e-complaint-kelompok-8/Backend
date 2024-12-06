package feedbacks

import (
	"capstone/entities"
	"capstone/repositories/models"

	"gorm.io/gorm"
)

type FeedbackRepositoryInterface interface {
	GetComplaintByID(complaintID int) (models.Complaint, error)
	GetFeedbackByComplaintID(complaintID int) (entities.Feedback, error)
	GetFeedbacksByUserID(userID int) ([]entities.Feedback, error)
	UpdateFeedbackResponse(feedbackID int, response string) error
	UpdateComplaintStatus(complaintID int, status string) error
	GetFeedbackByID(feedbackID int) (entities.Feedback, error)
}

type FeedbackRepository struct {
	db *gorm.DB
}

func NewFeedbackRepository(db *gorm.DB) *FeedbackRepository {
	return &FeedbackRepository{db: db}
}

func (fr *FeedbackRepository) GetComplaintByID(complaintID int) (models.Complaint, error) {
	var complaint models.Complaint
	err := fr.db.First(&complaint, "id = ?", complaintID).Error
	if err != nil {
		return models.Complaint{}, err
	}
	return complaint, nil
}

func (fr *FeedbackRepository) GetFeedbackByComplaintID(complaintID int) (entities.Feedback, error) {
	var feedback models.Feedback
	err := fr.db.Preload("Admin").
		Preload("User").
		Preload("Complaint.Category").
		Preload("Complaint.Photos").
		First(&feedback, "complaint_id = ?", complaintID).Error
	if err != nil {
		return entities.Feedback{}, err
	}
	return feedback.ToEntities(), nil
}

func (fr *FeedbackRepository) GetFeedbacksByUserID(userID int) ([]entities.Feedback, error) {
	var feedbacks []models.Feedback

	// Query database untuk mendapatkan feedback berdasarkan user_id
	err := fr.db.Preload("Admin").
		Preload("User").
		Preload("Complaint.Category").
		Preload("Complaint.Photos").
		Where("user_id = ?", userID).
		Find(&feedbacks).Error
	if err != nil {
		return nil, err
	}

	// Konversi model feedback ke entities
	var result []entities.Feedback
	for _, feedback := range feedbacks {
		result = append(result, feedback.ToEntities())
	}
	return result, nil
}

func (fr *FeedbackRepository) UpdateFeedbackResponse(feedbackID int, response string) error {
	// Perbarui kolom response di tabel feedback
	return fr.db.Model(&models.Feedback{}).Where("id = ?", feedbackID).Update("response", response).Error
}

func (fr *FeedbackRepository) UpdateComplaintStatus(complaintID int, status string) error {
	// Perbarui status di tabel complaints
	return fr.db.Model(&models.Complaint{}).Where("id = ?", complaintID).Update("status", status).Error
}

func (fr *FeedbackRepository) GetFeedbackByID(feedbackID int) (entities.Feedback, error) {
	var feedback models.Feedback
	err := fr.db.Preload("Admin").
		Preload("User").
		Preload("Complaint.Category").
		Preload("Complaint.Photos").
		First(&feedback, "id = ?", feedbackID).Error
	if err != nil {
		return entities.Feedback{}, err
	}
	return feedback.ToEntities(), nil
}
