package model

import "time"

type (
	Message struct {
		ID uint `json:"id"`
		// UserID   uint   `json:"user_id"`
		// UserName string `json:"user_name"`
		Text      string    `json:"text"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	CreateMessageRequest struct {
		Text string `json:"text"`
	}

	CreateMessageResponse struct {
		Message *Message `json:"message"`
	}

	ReadMessageRequest struct {
		Offset int64 `json:"offset"`
		Limit  int64 `json:"limit"`
	}

	ReadMessageResponse struct {
		Messages []*Message `json:"messages"`
	}
)
