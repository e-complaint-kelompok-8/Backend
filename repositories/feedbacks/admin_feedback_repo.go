package feedbacks

import (
	"capstone/entities"
	"capstone/repositories/models"

	"gorm.io/gorm"
)

func (cr *FeedbackRepository) CreateFeedback(feedback *entities.Feedback) error {
	feedbackModel := models.Feedback{
		AdminID:     feedback.AdminID,
		UserID:      feedback.UserID,
		ComplaintID: feedback.ComplaintID,
		Content:     feedback.Content,
		CreatedAt:   feedback.CreatedAt,
	}

	// Simpan ke database
	if err := cr.db.Create(&feedbackModel).Error; err != nil {
		return err
	}

	// Perbarui ID feedback
	feedback.ID = feedbackModel.ID
	return nil
}

func (cr *FeedbackRepository) CheckAdminExists(adminID int) (bool, error) {
	var count int64
	err := cr.db.Model(&models.Admin{}).Where("id = ?", adminID).Count(&count).Error
	return count > 0, err
}

func (cr *FeedbackRepository) CheckUserExists(userID int) (bool, error) {
	var count int64
	err := cr.db.Model(&models.User{}).Where("id = ?", userID).Count(&count).Error
	return count > 0, err
}

func (cr *FeedbackRepository) AdminUpdateComplaintStatus(complaintID int, newStatus string, adminID int) error {
	// Perbarui status dan admin ID di tabel complaints
	return cr.db.Model(&models.Complaint{}).Where("id = ?", complaintID).Updates(map[string]interface{}{
		"status":     newStatus,
		"admin_id":   adminID,
		"updated_at": gorm.Expr("NOW()"),
	}).Error
}
