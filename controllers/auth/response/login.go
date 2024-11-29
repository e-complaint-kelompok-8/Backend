package response

import "capstone/entities"

type LoginResponse struct {
	Message string       `json:"message"`
	Admin   entities.Admin `json:"admin"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}