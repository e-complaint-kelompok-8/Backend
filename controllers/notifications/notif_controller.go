package notifications

import (
	"capstone/entities"
	"capstone/services/notifications"
	"capstone/controllers/notifications/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// NotificationController adalah struct controller untuk notifikasi
type NotificationController struct {
	notificationService notifications.NotificationServiceInterface
}

// NewNotificationController membuat instance baru NotificationController
func NewNotificationController(service notifications.NotificationServiceInterface) *NotificationController {
	return &NotificationController{
		notificationService: service,
	}
}

// CreateNotification adalah handler untuk membuat notifikasi baru
func (controller *NotificationController) CreateNotification(c echo.Context) error {
	var notificationRequest entities.Notification

	// Binding request body ke struct Notification
	if err := c.Bind(&notificationRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	// Panggil service untuk membuat notifikasi
	createdNotification, err := controller.notificationService.CreateNotification(&notificationRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create notification"})
	}

	// Format respons dan kirim ke client
	responseData := response.FromEntityNotification(*createdNotification)
	return c.JSON(http.StatusCreated, responseData)
}

// GetNotificationsByUserID adalah handler untuk mengambil semua notifikasi berdasarkan user_id
func (controller *NotificationController) GetNotificationsByUserID(c echo.Context) error {
	// Ambil userID dari parameter URL
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	// Panggil service untuk mengambil notifikasi berdasarkan userID
	notifications, err := controller.notificationService.GetNotificationsByUserID(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch notifications"})
	}

	// Format respons dan kirim ke client
	responseData := response.FormatNotifications(notifications)
	return c.JSON(http.StatusOK, responseData)
}