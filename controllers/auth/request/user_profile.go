package request

type UpdateNameRequest struct {
	Name string `json:"name" validate:"required"`
}

type UpdatePhotoRequest struct {
	PhotoURL string `json:"photo" validate:"required,url"`
}

type UpdatePasswordRequest struct {
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required,min=6"`
}
