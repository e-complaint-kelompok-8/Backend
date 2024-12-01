package request

import "capstone/entities"

type LoginAdminRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (lar LoginAdminRequest) ToEntities() entities.Admin {
	return entities.Admin{
		Email:    lar.Email,
		Password: lar.Password,
	}
}