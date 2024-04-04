package router

import (
	"chat_san/handler"
	"chat_san/service"
	"database/sql"
	"net/http"
)

func NewRouter(db *sql.DB) *http.ServeMux {
	msgSrv := service.NewMessageService(db)

	mux := http.NewServeMux()
	mux.Handle("/messages", handler.NewMessageHandler(msgSrv))

	return mux
}
