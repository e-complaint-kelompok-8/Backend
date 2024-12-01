package response

import "capstone/entities"

type LoginResponse struct {
	Message string         `json:"message"`
	Admin   entities.Admin `json:"admin"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type LoginUserResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"no_telp"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func LoginFromEntities(user entities.User) LoginUserResponse {
	return LoginUserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Phone: user.Phone,
		Email: user.Email,
		Token: user.Token,
	}
}
