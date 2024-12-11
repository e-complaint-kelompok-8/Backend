package response

import "capstone/entities"

type UserProfileResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone_number"`
	Email    string `json:"email"`
	Password string `json:"password"`
	PhotoURL string `json:"photo"`
}

func UserProfileFromEntities(user entities.User) UserProfileResponse {
	return UserProfileResponse{
		ID:       user.ID,
		Name:     user.Name,
		Phone:    user.Phone,
		Email:    user.Email,
		Password: user.Password,
		PhotoURL: user.PhotoURL,
	}
}
