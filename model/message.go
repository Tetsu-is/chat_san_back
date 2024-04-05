package model

import "time"

type (
	Message struct {
		ID        int64     `json:"id"`
		Body      string    `json:"body"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	CreateMessageRequest struct {
		Body string `json:"body"`
	}

	CreateMessageResponse struct {
		Message Message `json:"message"`
	}

	ReadMessageRequest struct {
		Offset int64 `json:"offset"`
		Limit  int64 `json:"limit"`
	}

	ReadMessageResponse struct {
		Messages []*Message `json:"messages"`
	}
)
