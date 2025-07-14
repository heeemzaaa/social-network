package routes

import (
	"database/sql"
	"net/http"

	h "social-network/backend/handlers/chat"
	r "social-network/backend/repositories/chat"
	auth "social-network/backend/services/auth"
	s "social-network/backend/services/chat"
)

func SetChatRoutes(mux *http.ServeMux, db *sql.DB, authService *auth.AuthService) *http.ServeMux {
	repo := r.NewChatRepository(db)
	service := s.NewChatService(repo)
	handlerChat := h.NewChatServer(service)
	handlerMessage := h.NewMessagesHandler(service)

	mux.Handle("/ws/chat", handlerChat)
	mux.Handle("/api/", handlerMessage)

	return mux
}
