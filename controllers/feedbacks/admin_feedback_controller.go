package feedbacks

import (
	"capstone/controllers/feedbacks/request"
	"capstone/controllers/feedbacks/response"
	"capstone/middlewares"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (cc *FeedbackController) ProvideFeedback(c echo.Context) error {
	// Validasi role admin
	role, err := middlewares.ExtractAdminRole(c)
	if err != nil || role != "admin" {
		return c.JSON(http.StatusForbidden, map[string]string{"message": "Access denied"})
	}

	// Ambil admin ID dari JWT
	adminID, ok := c.Get("admin_id").(int)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid admin token"})
	}

	// Bind request
	req := request.FeedbackRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request payload"})
	}

	// Panggil service untuk memberikan feedback
	feedback, err := cc.feedbackService.ProvideFeedback(adminID, req.ComplaintID, req.Content)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	// Gunakan response untuk menampilkan data lengkap
	response := response.FromEntityFeedback(feedback)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "Feedback provided successfully",
		"feedback": response,
	})
}

func (cc *FeedbackController) UpdateFeedback(c echo.Context) error {
	// Validasi role admin
	role, err := middlewares.ExtractAdminRole(c)
	if err != nil || role != "admin" {
		return c.JSON(http.StatusForbidden, map[string]string{"message": "Access denied"})
	}

	// Ambil feedback ID dari parameter
	feedbackID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid feedback ID"})
	}

	// Bind request
	req := request.FeedbackRequesContent{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request payload"})
	}

	// Panggil service untuk memperbarui feedback
	updatedFeedback, err := cc.feedbackService.UpdateFeedback(feedbackID, req.Content)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	// Gunakan response untuk menampilkan data yang diperbarui
	response := response.FromEntityFeedback(updatedFeedback)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "Feedback updated successfully",
		"feedback": response,
	})
}
