package response

import (
	"capstone/entities"
	"time"
)

type FeedbackResponse struct {
	ID        int       `json:"id"`
	Admin     Admin     `json:"admin"`
	User      User      `json:"user"`
	Complaint Complaint `json:"complaint"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type Admin struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone_number"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Complaint struct {
	ID              int              `json:"id"`
	Category        Category         `json:"category"`
	ComplaintNumber string           `json:"complaint_number"`
	Title           string           `json:"title"`
	Location        string           `json:"location"`
	Status          string           `json:"status"`
	Description     string           `json:"description"`
	Photos          []ComplaintPhoto `json:"photos"`
	CreatedAt       time.Time        `json:"created_at"`
	UpdatedAt       time.Time        `json:"updated_at"`
}

type Category struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ComplaintPhoto struct {
	ID        int       `json:"id"`
	PhotoURL  string    `json:"photo_url"`
	CreatedAt time.Time `json:"created_at"`
}

type FeedbackResponseWithResponse struct {
	ID        int       `json:"id"`
	Admin     Admin     `json:"admin"`
	User      User      `json:"user"`
	Complaint Complaint `json:"complaint"`
	Content   string    `json:"content"`
	Response  string    `json:"response"` // Tambahkan kolom ini
	CreatedAt time.Time `json:"created_at"`
}

func FromEntityFeedback(feedback entities.Feedback) FeedbackResponse {
	return FeedbackResponse{
		ID: feedback.ID,
		Admin: Admin{
			ID:        feedback.Admin.ID,
			Email:     feedback.Admin.Email,
			Role:      feedback.Admin.Role,
			CreatedAt: feedback.Admin.CreatedAt,
			UpdatedAt: feedback.Admin.UpdatedAt,
		},
		User: User{
			ID:        feedback.User.ID,
			Name:      feedback.User.Name,
			Email:     feedback.User.Email,
			Phone:     feedback.User.Phone,
			CreatedAt: feedback.User.CreatedAt,
			UpdatedAt: feedback.User.UpdatedAt,
		},
		Complaint: Complaint{
			ID:              feedback.Complaint.ID,
			Category:        Category(feedback.Complaint.Category),
			ComplaintNumber: feedback.Complaint.ComplaintNumber,
			Title:           feedback.Complaint.Title,
			Location:        feedback.Complaint.Location,
			Status:          feedback.Complaint.Status,
			Description:     feedback.Complaint.Description,
			Photos:          FromEntityPhotos(feedback.Complaint.Photos),
			CreatedAt:       feedback.Complaint.CreatedAt,
			UpdatedAt:       feedback.Complaint.UpdatedAt,
		},
		Content:   feedback.Content,
		CreatedAt: feedback.CreatedAt,
	}
}

func FromEntityPhotos(photos []entities.ComplaintPhoto) []ComplaintPhoto {
	var photoResponses []ComplaintPhoto
	for _, photo := range photos {
		photoResponses = append(photoResponses, ComplaintPhoto{
			ID:        photo.ID,
			PhotoURL:  photo.PhotoURL,
			CreatedAt: photo.CreatedAt,
		})
	}
	return photoResponses
}

func FromEntitiesFeedbacks(feedbacks []entities.Feedback) []FeedbackResponse {
	var responses []FeedbackResponse
	for _, feedback := range feedbacks {
		responses = append(responses, FromEntityFeedback(feedback))
	}
	return responses
}

func FromEntityFeedbackWithResponse(feedback entities.Feedback) FeedbackResponseWithResponse {
	return FeedbackResponseWithResponse{
		ID: feedback.ID,
		Admin: Admin{
			ID:        feedback.Admin.ID,
			Email:     feedback.Admin.Email,
			Role:      feedback.Admin.Role,
			CreatedAt: feedback.Admin.CreatedAt,
			UpdatedAt: feedback.Admin.UpdatedAt,
		},
		User: User{
			ID:        feedback.User.ID,
			Name:      feedback.User.Name,
			Email:     feedback.User.Email,
			Phone:     feedback.User.Phone,
			CreatedAt: feedback.User.CreatedAt,
			UpdatedAt: feedback.User.UpdatedAt,
		},
		Complaint: Complaint{
			ID:              feedback.Complaint.ID,
			Category:        Category(feedback.Complaint.Category),
			ComplaintNumber: feedback.Complaint.ComplaintNumber,
			Title:           feedback.Complaint.Title,
			Location:        feedback.Complaint.Location,
			Status:          feedback.Complaint.Status,
			Description:     feedback.Complaint.Description,
			Photos:          FromEntityPhotos(feedback.Complaint.Photos),
			CreatedAt:       feedback.Complaint.CreatedAt,
			UpdatedAt:       feedback.Complaint.UpdatedAt,
		},
		Content:   feedback.Content,
		Response:  feedback.Response, // Tampilkan responsenya
		CreatedAt: feedback.CreatedAt,
	}
}
