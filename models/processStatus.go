package models

// ProcessStatus :
type ProcessStatus struct {
	Status  string `json:"status" example:"full-processed-message"`
	Message MessageForRead
	Error   error
}
