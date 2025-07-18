package routes

import (
	"database/sql"
	"net/http"

	"social-network/backend/middleware"
	"social-network/backend/services/auth"

	h "social-network/backend/handlers/notification"
	r "social-network/backend/repositories/notification"
	s "social-network/backend/services/notification"
)

func SetNotificationsRoutes(mux *http.ServeMux, db *sql.DB, authService *auth.AuthService) *http.ServeMux {
	repo := r.NewNotifRepository(db)
	service := s.NewNotifService(repo)

	new := h.NewNotificationHandler(service)

	mux.Handle("/api/notifications/", middleware.NewMiddleWare(new, authService))
	// mux.Handle("/api/notification/", new)

	// mux.HandleFunc("/api/notification", handlers.GetAllNotification)
	// mux.HandleFunc("/api/notif", handlers.HandlerNotif)





	return mux
}
