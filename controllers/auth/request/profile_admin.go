package request

type UpdateAdminRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Photo    string `json:"photo"`
}
