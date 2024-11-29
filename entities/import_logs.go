package entities

import "time"

type ImportLog struct {
	ID           int       `json:"id"`
	FileName     string    `json:"file_name"`
	ImportedBy   int       `json:"imported_by"`
	SuccessCount int       `json:"success_count"`
	FailureCount int       `json:"failure_count"`
	CreatedAt    time.Time `json:"created_at"`
}