package feedbacks

import (
	"capstone/controllers/feedbacks/response"
	"capstone/middlewares"
	"net/http"

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
	var req struct {
		ComplaintID int    `json:"complaint_id" validate:"required"`
		Content     string `json:"content" validate:"required"`
	}
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
