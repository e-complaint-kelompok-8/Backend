package feedbacks

import (
	"capstone/controllers/feedbacks/request"
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

func (fc *FeedbackController) AddResponseToFeedback(c echo.Context) error {
	// Ambil User ID dari middleware
	userID, ok := c.Get("user_id").(int)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "User tidak memiliki otorisasi",
		})
	}

	// Ambil feedback ID dari parameter
	feedbackID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "ID feedback tidak valid",
		})
	}

	// Ambil data balasan dari body
	request := request.Feedbackrequest{}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Data Balasan Tidak Valid",
		})
	}

	// Tambahkan balasan melalui service
	err = fc.feedbackService.AddResponseToFeedback(feedbackID, userID, request.Response)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	// Ambil feedback terbaru setelah balasan berhasil ditambahkan
	feedback, err := fc.feedbackService.GetFeedbackByID(feedbackID, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Gagal mengambil data feedback terbaru",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "Balasan Berhasil Ditambahkan",
		"feedback": response.FromEntityFeedbackWithResponse(feedback),
	})
}
