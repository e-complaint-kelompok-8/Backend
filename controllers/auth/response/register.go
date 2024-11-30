package response

import "capstone/entities"

type RegisterResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"no_telp"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func RegisterFromEntities(user entities.User) RegisterResponse {
	return RegisterResponse{
		ID:    user.ID,
		Name:  user.Name,
		Phone: user.Phone,
		Email: user.Email,
		Role:  user.Role,
	}
}
