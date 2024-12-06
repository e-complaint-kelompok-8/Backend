package feedbacks

import (
	"capstone/controllers/feedbacks/response"
	feedback "capstone/services/feedbacks"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type FeedbackController struct {
	feedbackService feedback.FeedbackServiceInterface
}

func NewFeedbackController(service feedback.FeedbackServiceInterface) *FeedbackController {
	return &FeedbackController{feedbackService: service}
}

func (fc *FeedbackController) GetFeedbackByComplaint(c echo.Context) error {
	// Ambil User ID dari middleware
	userID, ok := c.Get("user_id").(int)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "User not authorized",
		})
	}

	// Ambil complaint_id dari parameter
	complaintID, err := strconv.Atoi(c.Param("complaint_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid complaint ID",
		})
	}

	// Ambil feedback melalui service
	feedback, err := fc.feedbackService.GetFeedbackByComplaint(complaintID, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "Feedback retrieved successfully",
		"feedback": response.FromEntityFeedback(feedback),
	})
}

func (fc *FeedbackController) GetFeedbacksByUser(c echo.Context) error {
	// Ambil User ID dari middleware
	userID, ok := c.Get("user_id").(int)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "User not authorized",
		})
	}

	// Ambil semua feedback dari service
	feedbacks, err := fc.feedbackService.GetFeedbacksByUser(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to retrieve feedbacks",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":   "Feedbacks retrieved successfully",
		"feedbacks": response.FromEntitiesFeedbacks(feedbacks),
	})
}
