package response

import "capstone/entities"

type AdminProfileResponse struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Photo    string `json:"photo"`
}

func AdminProfileFromEntities(admin entities.Admin) AdminProfileResponse {
	return AdminProfileResponse{
		ID:       admin.ID,
		Email:    admin.Email,
		Password: admin.Password,
		Role:     admin.Role,
		Photo:    admin.Photo,
	}
}
