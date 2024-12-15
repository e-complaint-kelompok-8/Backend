package feedbacks

import (
	"capstone/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

var feedbackService FeedbackService

type FeedbackRepoDummy struct {
	ShouldFail bool
}

func (repo FeedbackRepoDummy) GetComplaintByID(complaintID int) (entities.Complaint, error) {

}
func (repo FeedbackRepoDummy) GetFeedbackByComplaintID(complaintID int) (entities.Feedback, error) {

}
func (repo FeedbackRepoDummy) GetFeedbacksByUserID(userID int) ([]entities.Feedback, error) {

}
func (repo FeedbackRepoDummy) UpdateFeedbackResponse(feedbackID int, response string) error {

}
func (repo FeedbackRepoDummy) UpdateComplaintStatus(complaintID int, status string) error {

}
func (repo FeedbackRepoDummy) GetFeedbackByID(feedbackID int) (entities.Feedback, error) {

}
func (repo FeedbackRepoDummy) CreateFeedback(feedback *entities.Feedback) error {

}
func (repo FeedbackRepoDummy) CheckAdminExists(adminID int) (bool, error) {

}
func (repo FeedbackRepoDummy) CheckUserExists(userID int) (bool, error) {

}
func (repo FeedbackRepoDummy) AdminUpdateComplaintStatus(complaintID int, newStatus string, adminID int) error {
}
func (repo FeedbackRepoDummy) ComplaintHasFeedback(complaintID int) (bool, error) {

}
func (repo FeedbackRepoDummy) UpdateFeedback(feedback entities.Feedback) error {

}

func setupTestService() {
	repo := FeedbackRepoDummy{}
	feedbackService = *NewFeedbackService(repo)
}

func TestCustomerService_GetAllUsers(t *testing.T) {
	setupTestService()

	t.Run("sukses", func(t *testing.T) {
		// Data dummy untuk pengujian
		feedback, err := feedbackService.GetFeedbackByComplaint(1, 1)

		// Periksa apakah error tidak terjadi
		assert.NoError(t, err)
	})

}
