package response

import "capstone/entities"

type User struct {
	ID         int         `json:"id"`
	Name       string      `json:"name"`
	Phone      string      `json:"phone"`
	Email      string      `json:"email"`
	PhotoURL   string      `json:"photo_url"`
	Verified   bool        `json:"verified"`
	Complaints []Complaint `json:"complaints"`
}

func FromEntitiesUsers(users []entities.User) []User {
	var userResponses []User
	for _, user := range users {
		userResponses = append(userResponses, User{
			ID:         user.ID,
			Name:       user.Name,
			Phone:      user.Phone,
			Email:      user.Email,
			PhotoURL:   user.PhotoURL,
			Verified:   user.Verified,
			Complaints: FromEntitiesComplaints(user.Complaints),
		})
	}
	return userResponses
}

type Complaint struct {
	ID          int              `json:"id"`
	Title       string           `json:"title"`
	Location    string           `json:"location"`
	Status      string           `json:"status"`
	Description string           `json:"description"`
	Category    Category         `json:"category"`
	Photos      []ComplaintPhoto `json:"photos"`
	Feedbacks   []Feedback       `json:"feedbacks"`
	Reason      string           `json:"reason"`
	CreatedAt   string           `json:"created_at"`
}

func FromEntitiesComplaints(complaints []entities.Complaint) []Complaint {
	var complaintResponses []Complaint
	for _, complaint := range complaints {
		complaintResponses = append(complaintResponses, Complaint{
			ID:          complaint.ID,
			Title:       complaint.Title,
			Location:    complaint.Location,
			Status:      complaint.Status,
			Description: complaint.Description,
			Category:    FromEntitiesCategory(complaint.Category),
			Photos:      FromEntitiesPhotos(complaint.Photos),
			Feedbacks:   FromEntitiesFeedbacks(complaint.Feedbacks),
			Reason:      complaint.Reason,
			CreatedAt:   complaint.CreatedAt.Format("02 Jan 2006"),
		})
	}
	return complaintResponses
}

func FromEntityUser(user entities.User) User {
	return User{
		ID:         user.ID,
		Name:       user.Name,
		Phone:      user.Phone,
		Email:      user.Email,
		PhotoURL:   user.PhotoURL,
		Verified:   user.Verified,
		Complaints: FromEntitiesComplaints(user.Complaints),
	}
}

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func FromEntitiesCategory(category entities.Category) Category {
	return Category{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}
}

type ComplaintPhoto struct {
	ID       int    `json:"id"`
	PhotoURL string `json:"photo_url"`
}

func FromEntitiesPhotos(photos []entities.ComplaintPhoto) []ComplaintPhoto {
	var photoResponses []ComplaintPhoto
	for _, photo := range photos {
		photoResponses = append(photoResponses, ComplaintPhoto{
			ID:       photo.ID,
			PhotoURL: photo.PhotoURL,
		})
	}
	return photoResponses
}

type Feedback struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

func FromEntitiesFeedbacks(feedbacks []entities.Feedback) []Feedback {
	var feedbackResponses []Feedback
	for _, feedback := range feedbacks {
		feedbackResponses = append(feedbackResponses, Feedback{
			ID:          feedback.ID,
			Description: feedback.Content,
		})
	}
	return feedbackResponses
}
