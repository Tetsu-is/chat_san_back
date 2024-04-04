package handler

import (
	"chat_san/model"
	"chat_san/service"
	"encoding/json"
	"net/http"
	"strconv"
)

type MessageHandler struct {
	svc *service.MessageService
}

func NewMessageHandler(svc *service.MessageService) *MessageHandler {
	return &MessageHandler{svc: svc}
}

func (h *MessageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		req := &model.ReadMessageRequest{}

		q := r.URL.Query()
		offset := q.Get("offset")
		limit := q.Get("limit")

		parsedOffset, err := strconv.ParseInt(offset, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		parsedLimit, err := strconv.ParseInt(limit, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		req.Offset, req.Limit = parsedOffset, parsedLimit

		res, err := h.Read(r.Context(), req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
