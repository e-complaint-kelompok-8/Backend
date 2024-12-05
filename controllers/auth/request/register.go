package request

import "capstone/entities"

type RegisterRequest struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone_number"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (rr RegisterRequest) ToEntities() entities.User {
	return entities.User{
		ID:       rr.ID,
		Name:     rr.Name,
		Phone:    rr.Phone,
		Email:    rr.Email,
		Password: rr.Password,
	}
}
