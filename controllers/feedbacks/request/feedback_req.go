package request

type Feedbackrequest struct {
	Response string `json:"response" validate:"required"`
}
