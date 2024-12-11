package request

type RequestCS struct {
	Query string `json:"request" validate:"required"`
}
