package notifications

import (
	"capstone/entities"
	"capstone/repositories/models"
	"gorm.io/gorm"
)

// NotificationRepositoryInterface is the interface for interacting with notifications
type NotificationRepositoryInterface interface {
	CreateNotification(notification *entities.Notification) (*entities.Notification, error)
	GetNotificationsByUserID(userID int) ([]entities.Notification, error)
}

// NotificationRepository struct to interact with the Notification model
type NotificationRepository struct {
	DB *gorm.DB
}

// NewNotificationRepository is a constructor function to create a new NotificationRepository
func NewNotificationRepository(db *gorm.DB) NotificationRepositoryInterface {
	return &NotificationRepository{
		DB: db,
	}
}

// CreateNotification function to create a new notification
func (repo *NotificationRepository) CreateNotification(notification *entities.Notification) (*entities.Notification, error) {
	// Convert entities.Notification to models.Notification
	modelNotification := models.Notification{
		UserID:  uint(notification.UserID), // Konversi int ke uint
		Content: notification.Content,
		IsRead:  notification.IsRead,
	}

	// Save the notification to the database
	if err := repo.DB.Create(&modelNotification).Error; err != nil {
		return nil, err
	}

	// Return the created notification as an entity
	notification.ID = int(modelNotification.ID)
	notification.CreatedAt = modelNotification.CreatedAt
	return notification, nil
}

// GetNotificationsByUserID function to get notifications for a specific user
func (repo *NotificationRepository) GetNotificationsByUserID(userID int) ([]entities.Notification, error) {
	var notifications []models.Notification

	// Fetch notifications from the database
	if err := repo.DB.Where("user_id = ?", uint(userID)).Order("created_at desc").Find(&notifications).Error; err != nil {
		return nil, err
	}

	// Convert model notifications to entity notifications
	var result []entities.Notification
	for _, notification := range notifications {
		result = append(result, entities.Notification{
			ID:        int(notification.ID),
			UserID:    int(notification.UserID), // Konversi uint ke int
			Content:   notification.Content,
			IsRead:    notification.IsRead,
			CreatedAt: notification.CreatedAt,
		})
	}
	return result, nil
}