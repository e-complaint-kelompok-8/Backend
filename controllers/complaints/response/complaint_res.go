package response

import (
	"capstone/entities"
	"time"
)

type CreateComplaintResponse struct {
	ID              int        `json:"id"`
	User            User       `json:"user"`
	Category        Category   `json:"category"`
	ComplaintNumber string     `json:"complaint_number"`
	Title           string     `json:"title"`
	Location        string     `json:"location"`
	Status          string     `json:"status"`
	Feedbacks       []Feedback `json:"feedback"`
	Description     string     `json:"description" validate:"required"`
	Reason          string     `json:"reason"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
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
	Feedbacks       []Feedback       `json:"feedback"`
	Reason          string           `json:"reason"`
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

type CreateComplaintResponseWithAdmin struct {
	ID              int              `json:"id"`
	Admin           Admin            `json:"Admin"`
	User            User             `json:"user"`
	Category        Category         `json:"category"`
	ComplaintNumber string           `json:"complaint_number"`
	Title           string           `json:"title"`
	Location        string           `json:"location"`
	Status          string           `json:"status"`
	Description     string           `json:"description" validate:"required"`
	Photos          []ComplaintPhoto `json:"photos"`
	Feedbacks       []Feedback       `json:"feedback"`
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

type Admin struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type Feedback struct {
	ID        int       `json:"id"`
	Admin     Admin     `json:"admin"`
	Content   string    `json:"content"`
	Response  string    `json:"response"`
	CreatedAt time.Time `json:"created_at"`
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

	// Konversi feedbacks
	var feedbackResponses []Feedback
	for _, feedback := range complaint.Feedbacks {
		feedbackResponses = append(feedbackResponses, Feedback{
			ID: feedback.ID,
			Admin: Admin{
				ID:    feedback.Admin.ID,
				Email: feedback.Admin.Email,
				Role:  feedback.Admin.Role,
			},
			Content:   feedback.Content,
			Response:  feedback.Response,
			CreatedAt: feedback.CreatedAt,
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
		Feedbacks:       feedbackResponses,
		Reason:          complaint.Reason,
		CreatedAt:       complaint.CreatedAt,
		UpdatedAt:       complaint.UpdatedAt,
	}
}

func ComplaintFromEntities(complaint entities.Complaint) CreateComplaintResponse {

	// Konversi feedbacks
	var feedbackResponses []Feedback
	for _, feedback := range complaint.Feedbacks {
		feedbackResponses = append(feedbackResponses, Feedback{
			ID: feedback.ID,
			Admin: Admin{
				ID:    feedback.Admin.ID,
				Email: feedback.Admin.Email,
				Role:  feedback.Admin.Role,
			},
			Content:   feedback.Content,
			Response:  feedback.Response,
			CreatedAt: feedback.CreatedAt,
		})
	}

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
		Feedbacks:       feedbackResponses,
		Reason:          complaint.Reason,
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

		// Konversi feedbacks
		var feedbackResponses []Feedback
		for _, feedback := range complaint.Feedbacks {
			feedbackResponses = append(feedbackResponses, Feedback{
				ID: feedback.ID,
				Admin: Admin{
					ID:    feedback.Admin.ID,
					Email: feedback.Admin.Email,
					Role:  feedback.Admin.Role,
				},
				Content:   feedback.Content,
				Response:  feedback.Response,
				CreatedAt: feedback.CreatedAt,
			})
		}

		// Periksa apakah kategori tersedia
		var categoryResponse Category
		if complaint.Category.ID > 0 {
			categoryResponse = Category{
				ID:          complaint.Category.ID,
				Name:        complaint.Category.Name,
				Description: complaint.Category.Description,
				CreatedAt:   complaint.Category.CreatedAt,
				UpdatedAt:   complaint.Category.UpdatedAt,
			}
		}

		// Tambahkan respons
		responses = append(responses, CreateComplaintResponseWithPhoto{
			ID: complaint.ID,
			User: User{
				ID:    complaint.User.ID,
				Name:  complaint.User.Name,
				Phone: complaint.User.Phone,
				Email: complaint.User.Email,
			},
			Category:        categoryResponse,
			ComplaintNumber: complaint.ComplaintNumber,
			Title:           complaint.Title,
			Location:        complaint.Location,
			Status:          complaint.Status,
			Description:     complaint.Description,
			Photos:          photoResponses,
			Feedbacks:       feedbackResponses,
			Reason:          complaint.Reason,
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

func ComplaintFromEntitiesWithAdmin(complaint entities.Complaint) CreateComplaintResponseWithAdmin {
	// Konversi daftar foto
	var photoResponses []ComplaintPhoto
	for _, photo := range complaint.Photos {
		photoResponses = append(photoResponses, ComplaintPhoto{
			ID:       photo.ID,
			PhotoURL: photo.PhotoURL,
		})
	}

	// Konversi feedbacks
	var feedbackResponses []Feedback
	for _, feedback := range complaint.Feedbacks {
		feedbackResponses = append(feedbackResponses, Feedback{
			ID: feedback.ID,
			Admin: Admin{
				ID:    feedback.Admin.ID,
				Email: feedback.Admin.Email,
				Role:  feedback.Admin.Role,
			},
			Content:   feedback.Content,
			Response:  feedback.Response,
			CreatedAt: feedback.CreatedAt,
		})
	}

	return CreateComplaintResponseWithAdmin{
		ID: complaint.ID,
		Admin: Admin{
			ID:    complaint.Admin.ID,
			Email: complaint.Admin.Email,
			Role:  complaint.Admin.Role,
		},
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
		Reason:          complaint.Reason,
		Feedbacks:       feedbackResponses,
		CreatedAt:       complaint.CreatedAt,
		UpdatedAt:       complaint.UpdatedAt,
	}
}

func ComplaintFromEntitiesWithFeedback(complaint entities.Complaint) CreateComplaintResponseWithAdmin {
	var photoResponses []ComplaintPhoto
	for _, photo := range complaint.Photos {
		photoResponses = append(photoResponses, ComplaintPhoto{
			ID:       photo.ID,
			PhotoURL: photo.PhotoURL,
		})
	}

	// Konversi feedbacks
	var feedbackResponses []Feedback
	for _, feedback := range complaint.Feedbacks {
		feedbackResponses = append(feedbackResponses, Feedback{
			ID: feedback.ID,
			Admin: Admin{
				ID:    feedback.Admin.ID,
				Email: feedback.Admin.Email,
				Role:  feedback.Admin.Role,
			},
			Content:   feedback.Content,
			Response:  feedback.Response,
			CreatedAt: feedback.CreatedAt,
		})
	}

	return CreateComplaintResponseWithAdmin{
		ID: complaint.ID,
		Admin: Admin{
			ID:    complaint.Admin.ID,
			Email: complaint.Admin.Email,
			Role:  complaint.Admin.Role,
		},
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
		Reason:          complaint.Reason,
		Feedbacks:       feedbackResponses,
		CreatedAt:       complaint.CreatedAt,
		UpdatedAt:       complaint.UpdatedAt,
	}
}
