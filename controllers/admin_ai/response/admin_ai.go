package response

import (
	"capstone/entities"
	"time"
)

type AISuggestion struct {
	ID        int       `json:"id"`
	Admin     Admin     `json:"admin"`
	Complaint Complaint `json:"complaint"`
	Request   string    `json:"request"`
	Response  string    `json:"response"`
	CreatedAt time.Time `json:"create_at"`
}

type Admin struct {
	ID    int
	Email string
	Photo string
	Role  string
}

type Complaint struct {
	ID              int              `json:"id"`
	User            User             `json:"user"`
	Category        Category         `json:"category"`
	ComplaintNumber string           `json:"complaint_number"`
	Title           string           `json:"title"`
	Location        string           `json:"location"`
	Status          string           `json:"status"`
	Description     string           `json:"description" validate:"required"`
	Photos          []ComplaintPhoto `json:"photos"`
}

type Category struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"create_at"`
	UpdatedAt   time.Time `json:"update_at"`
}

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone_number"`
	Email string `json:"email"`
}

type ComplaintPhoto struct {
	ID       int    `json:"id"`
	PhotoURL string `json:"photo_url"`
}

func FromEntityPhotos(photos []entities.ComplaintPhoto) []ComplaintPhoto {
	var photoResponses []ComplaintPhoto
	for _, photo := range photos {
		photoResponses = append(photoResponses, ComplaintPhoto{
			ID:       photo.ID,
			PhotoURL: photo.PhotoURL,
		})
	}
	return photoResponses
}

func AISuggestionFromEntities(savedAISuggestion entities.AISuggestion, admin entities.Admin) AISuggestion {
	return AISuggestion{
		ID: savedAISuggestion.ID,
		Admin: Admin{
			ID:    admin.ID,
			Email: admin.Email,
			Photo: admin.Photo,
			Role:  admin.Role,
		},
		Complaint: Complaint{
			ID:              savedAISuggestion.Complaint.ID,
			User:            User{ID: savedAISuggestion.Complaint.User.ID, Name: savedAISuggestion.Complaint.User.Name, Phone: savedAISuggestion.Complaint.User.Phone, Email: savedAISuggestion.Complaint.User.Email},
			Category:        Category{ID: savedAISuggestion.Complaint.Category.ID, Name: savedAISuggestion.Complaint.Category.Name, Description: savedAISuggestion.Complaint.Category.Description},
			ComplaintNumber: savedAISuggestion.Complaint.ComplaintNumber,
			Title:           savedAISuggestion.Complaint.Title,
			Location:        savedAISuggestion.Complaint.Location,
			Status:          savedAISuggestion.Complaint.Status,
			Description:     savedAISuggestion.Complaint.Description,
			Photos:          FromEntityPhotos(savedAISuggestion.Complaint.Photos),
		},
		Request:   savedAISuggestion.Request,
		Response:  savedAISuggestion.Response,
		CreatedAt: savedAISuggestion.CreatedAt,
	}
}

func FormattedSuggestions(aiSuggestions []entities.AISuggestion) []AISuggestion {
	var formattedSuggestions []AISuggestion
	for _, suggestion := range aiSuggestions {
		formattedSuggestions = append(formattedSuggestions, AISuggestion{
			ID:    suggestion.ID,
			Admin: Admin{ID: suggestion.Admin.ID, Email: suggestion.Admin.Email, Photo: suggestion.Admin.Photo, Role: suggestion.Admin.Role},
			Complaint: Complaint{
				ID:              suggestion.Complaint.ID,
				User:            User{ID: suggestion.Complaint.User.ID, Name: suggestion.Complaint.User.Name, Phone: suggestion.Complaint.User.Phone, Email: suggestion.Complaint.User.Email},
				Category:        Category{ID: suggestion.Complaint.Category.ID, Name: suggestion.Complaint.Category.Name, Description: suggestion.Complaint.Category.Description},
				ComplaintNumber: suggestion.Complaint.ComplaintNumber,
				Title:           suggestion.Complaint.Title,
				Location:        suggestion.Complaint.Location,
				Status:          suggestion.Complaint.Status,
				Description:     suggestion.Complaint.Description,
				Photos:          FromEntityPhotos(suggestion.Complaint.Photos),
			},
			Request:   suggestion.Request,
			Response:  suggestion.Response,
			CreatedAt: suggestion.CreatedAt,
		})
	}
	return formattedSuggestions
}
