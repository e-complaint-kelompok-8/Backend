package notifications

import (
	"capstone/entities"
	notifications "capstone/repositories/notifications"
)

// NotificationServiceInterface defines the methods for the notification service
type NotificationServiceInterface interface {
	CreateNotification(notification *entities.Notification) (*entities.Notification, error)
	GetNotificationsByUserID(userID int) ([]entities.Notification, error)
}

// NotificationService implements NotificationServiceInterface
type NotificationService struct {
	notificationRepo notifications.NotificationRepositoryInterface
}

// NewNotificationService creates a new instance of NotificationService
func NewNotificationService(repo notifications.NotificationRepositoryInterface) NotificationServiceInterface {
	return &NotificationService{
		notificationRepo: repo,
	}
}

// CreateNotification creates a new notification
func (service *NotificationService) CreateNotification(notification *entities.Notification) (*entities.Notification, error) {
	// Add any additional business logic here, if needed
	createdNotification, err := service.notificationRepo.CreateNotification(notification)
	if err != nil {
		return nil, err
	}
	return createdNotification, nil
}

// GetNotificationsByUserID retrieves notifications for a specific user
func (service *NotificationService) GetNotificationsByUserID(userID int) ([]entities.Notification, error) {
	// Add any additional business logic here, such as filtering or validation
	notifications, err := service.notificationRepo.GetNotificationsByUserID(userID)
	if err != nil {
		return nil, err
	}
	return notifications, nil
}