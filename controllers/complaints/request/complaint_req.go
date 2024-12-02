package request

import "capstone/entities"

type CreateComplaintRequest struct {
	UserID          int      `json:"user_id" validate:"required"`
	ComplaintNumber string   `json:"complaint_number"`
	Title           string   `json:"title"`
	Location        string   `json:"location"`
	CategoryID      int      `json:"category_id" validate:"required"`
	Description     string   `json:"description" validate:"required"`
	PhotoURLs       []string `json:"photo_urls"`
}

func (req CreateComplaintRequest) ToEntity() entities.Complaint {
	return entities.Complaint{
		UserID:          req.UserID,
		CategoryID:      req.CategoryID,
		ComplaintNumber: req.ComplaintNumber,
		Title:           req.Title,
		Location:        req.Location,
		Description:     req.Description,
	}
}
