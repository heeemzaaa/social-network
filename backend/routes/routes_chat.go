package routes

import (
	"database/sql"
	"net/http"
	r "social-network/backend/repositories/chat"
	s "social-network/backend/services/chat"
	h "social-network/backend/handlers/chat"
)

func SetChatRoutes(mux *http.ServeMux, db *sql.DB) *http.ServeMux {
	repo := r.NewChatRepository(db)
	service := s.NewChatService(repo)
	handlerChat := h.NewChatServer(service)
	handlerMessage := h.NewMessagesHandler(service)

	mux.Handle("/ws/chat", handlerChat)
	mux.Handle("/api/messages/", handlerMessage)

	return mux
}
