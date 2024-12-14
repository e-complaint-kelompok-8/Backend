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
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Location  string `json:"location"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}

func FromEntitiesComplaints(complaints []entities.Complaint) []Complaint {
	var complaintResponses []Complaint
	for _, complaint := range complaints {
		complaintResponses = append(complaintResponses, Complaint{
			ID:        complaint.ID,
			Title:     complaint.Title,
			Location:  complaint.Location,
			Status:    complaint.Status,
			CreatedAt: complaint.CreatedAt.Format("02 Jan 2006"),
		})
	}
	return complaintResponses
}
