package response

import (
	"capstone/entities"
	"time"
)

type CreateComplaintResponse struct {
	ID              int       `json:"id"`
	User            User      `json:"user"`
	Category        Category  `json:"category"`
	ComplaintNumber string    `json:"complaint_number"`
	Title           string    `json:"title"`
	Location        string    `json:"location"`
	Status          string    `json:"status"`
	Description     string    `json:"description" validate:"required"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type CreateComplaintResponseWithPhoto struct {
	ID              int              `json:"id"`
	User            User             `json:"user"`
	Category        Category         `json:"category"`
	ComplaintNumber string           `json:"complaint_number"`
	Title           string           `json:"title"`
	Location        string           `json:"location"`
	Status          string           `json:"status"`
	Description     string           `json:"description" validate:"required"`
	Photos          []ComplaintPhoto `json:"photos"`
	CreatedAt       time.Time        `json:"created_at"`
	UpdatedAt       time.Time        `json:"updated_at"`
}

type CreateComplaintResponseWithReason struct {
	ID              int              `json:"id"`
	User            User             `json:"user"`
	Category        Category         `json:"category"`
	ComplaintNumber string           `json:"complaint_number"`
	Title           string           `json:"title"`
	Location        string           `json:"location"`
	Status          string           `json:"status"`
	Description     string           `json:"description" validate:"required"`
	Photos          []ComplaintPhoto `json:"photos"`
	Reason          string           `json:"reason"`
	CreatedAt       time.Time        `json:"created_at"`
	UpdatedAt       time.Time        `json:"updated_at"`
}

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone_number"`
	Email string `json:"email"`
}

type Category struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"create_at"`
	UpdatedAt   time.Time `json:"update_at"`
}

type ComplaintPhoto struct {
	ID       int    `json:"id"`
	PhotoURL string `json:"photo_url"`
}

// Fungsi untuk membuat respons dari entities.Complaint
func ComplaintFromEntitiesWithPhoto(complaint entities.Complaint, photos []entities.ComplaintPhoto) CreateComplaintResponseWithPhoto {
	// Konversi photos dari entities ke response
	var photoResponses []ComplaintPhoto
	for _, photo := range photos {
		photoResponses = append(photoResponses, ComplaintPhoto{
			ID:       photo.ID,
			PhotoURL: photo.PhotoURL,
		})
	}

	return CreateComplaintResponseWithPhoto{
		ID: complaint.ID,
		User: User{
			ID:    complaint.User.ID,
			Name:  complaint.User.Name,
			Phone: complaint.User.Phone,
			Email: complaint.User.Email,
		},
		Category: Category{
			ID:          complaint.Category.ID,
			Name:        complaint.Category.Name,
			Description: complaint.Category.Description,
			CreatedAt:   complaint.Category.CreatedAt,
			UpdatedAt:   complaint.Category.UpdatedAt,
		},
		ComplaintNumber: complaint.ComplaintNumber,
		Title:           complaint.Title,
		Location:        complaint.Location,
		Status:          complaint.Status,
		Description:     complaint.Description,
		Photos:          photoResponses,
		CreatedAt:       complaint.CreatedAt,
		UpdatedAt:       complaint.UpdatedAt,
	}
}

func ComplaintFromEntities(complaint entities.Complaint) CreateComplaintResponse {
	return CreateComplaintResponse{
		ID: complaint.ID,
		User: User{
			ID:    complaint.User.ID,
			Name:  complaint.User.Name,
			Phone: complaint.User.Phone,
			Email: complaint.User.Email,
		},
		Category: Category{
			ID:          complaint.Category.ID,
			Name:        complaint.Category.Name,
			Description: complaint.Category.Description,
			CreatedAt:   complaint.Category.CreatedAt,
			UpdatedAt:   complaint.Category.UpdatedAt,
		},
		ComplaintNumber: complaint.ComplaintNumber,
		Title:           complaint.Title,
		Location:        complaint.Location,
		Status:          complaint.Status,
		Description:     complaint.Description,
		CreatedAt:       complaint.CreatedAt,
		UpdatedAt:       complaint.UpdatedAt,
	}
}

func ComplaintsFromEntities(complaints []entities.Complaint) []CreateComplaintResponseWithPhoto {
	var responses []CreateComplaintResponseWithPhoto
	for _, complaint := range complaints {
		// Konversi daftar foto
		var photoResponses []ComplaintPhoto
		for _, photo := range complaint.Photos {
			photoResponses = append(photoResponses, ComplaintPhoto{
				ID:       photo.ID,
				PhotoURL: photo.PhotoURL,
			})
		}

		responses = append(responses, CreateComplaintResponseWithPhoto{
			ID: complaint.ID,
			User: User{
				ID:    complaint.User.ID,
				Name:  complaint.User.Name,
				Phone: complaint.User.Phone,
				Email: complaint.User.Email,
			},
			Category: Category{
				ID:          complaint.Category.ID,
				Name:        complaint.Category.Name,
				Description: complaint.Category.Description,
				CreatedAt:   complaint.Category.CreatedAt,
				UpdatedAt:   complaint.Category.UpdatedAt,
			},
			ComplaintNumber: complaint.ComplaintNumber,
			Title:           complaint.Title,
			Location:        complaint.Location,
			Status:          complaint.Status,
			Description:     complaint.Description,
			Photos:          photoResponses,
			CreatedAt:       complaint.CreatedAt,
			UpdatedAt:       complaint.UpdatedAt,
		})
	}
	return responses
}

func ComplaintFromEntitiesWithReason(complaint entities.Complaint, photos []entities.ComplaintPhoto) CreateComplaintResponseWithReason {
	var photoResponses []ComplaintPhoto
	for _, photo := range photos {
		photoResponses = append(photoResponses, ComplaintPhoto{
			ID:       photo.ID,
			PhotoURL: photo.PhotoURL,
		})
	}

	return CreateComplaintResponseWithReason{
		ID: complaint.ID,
		User: User{
			ID:    complaint.User.ID,
			Name:  complaint.User.Name,
			Phone: complaint.User.Phone,
			Email: complaint.User.Email,
		},
		Category: Category{
			ID:          complaint.Category.ID,
			Name:        complaint.Category.Name,
			Description: complaint.Category.Description,
			CreatedAt:   complaint.Category.CreatedAt,
			UpdatedAt:   complaint.Category.UpdatedAt,
		},
		ComplaintNumber: complaint.ComplaintNumber,
		Title:           complaint.Title,
		Location:        complaint.Location,
		Status:          complaint.Status,
		Description:     complaint.Description,
		Photos:          photoResponses,
		Reason:          complaint.Reason, // Alasan pembatalan
		CreatedAt:       complaint.CreatedAt,
		UpdatedAt:       complaint.UpdatedAt,
	}
}
