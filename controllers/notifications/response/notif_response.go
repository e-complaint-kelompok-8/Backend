package response

import (
	"capstone/entities"
	"time"
)

// NotificationResponse adalah struktur untuk merespons data notifikasi
type NotificationResponse struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Content   string    `json:"content"`
	IsRead    bool      `json:"is_read"`
	CreatedAt time.Time `json:"created_at"`
}

// FromEntityNotification mengubah struktur entities.Notification menjadi NotificationResponse
func FromEntityNotification(notification entities.Notification) NotificationResponse {
	return NotificationResponse{
		ID:        notification.ID,
		UserID:    notification.UserID,
		Content:   notification.Content,
		IsRead:    notification.IsRead,
		CreatedAt: notification.CreatedAt,
	}
}

// FormatNotifications mengubah slice entities.Notification menjadi slice NotificationResponse
func FormatNotifications(notifications []entities.Notification) []NotificationResponse {
	var formattedResponses []NotificationResponse
	for _, notification := range notifications {
		formattedResponses = append(formattedResponses, FromEntityNotification(notification))
	}
	return formattedResponses
}