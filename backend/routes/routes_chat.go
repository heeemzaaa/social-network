package routes

import (
	"database/sql"
	"net/http"

	h "social-network/backend/handlers/chat"
	"social-network/backend/middleware"
	r "social-network/backend/repositories/chat"
	auth "social-network/backend/services/auth"
	s "social-network/backend/services/chat"
)

func SetChatRoutes(mux *http.ServeMux, db *sql.DB, authService *auth.AuthService) *http.ServeMux {
	repo := r.NewChatRepository(db)
	service := s.NewChatService(repo)
	handlerChat := h.NewChatServer(service)
	handlerMessage := h.NewMessagesHandler(service)
	handlerChatNav := h.NewChatNavigation(service)

	mux.Handle("/ws/chat/", middleware.NewMiddleWare(handlerChat, authService))
	mux.Handle("/api/messages", middleware.NewMiddleWare(handlerMessage, authService))
	mux.Handle("/api/get-users", middleware.NewMiddleWare(handlerChatNav, authService))

	return mux
}
