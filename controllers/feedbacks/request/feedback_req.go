package request

type FeedbackRequesResponse struct {
	Response string `json:"response" validate:"required"`
}

type FeedbackRequesContent struct {
	Content string `json:"content" validate:"required"`
}

type FeedbackRequest struct {
	ComplaintID int    `json:"complaint_id" validate:"required"`
	Content     string `json:"content" validate:"required"`
}
