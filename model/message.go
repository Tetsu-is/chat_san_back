package model

type (
	Message struct {
		ID uint `json:"id"`
		// UserID   uint   `json:"user_id"`
		// UserName string `json:"user_name"`
		Text string `json:"text"`
	}
)
