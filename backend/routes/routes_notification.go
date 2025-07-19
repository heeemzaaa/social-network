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
	service_update := s.NewNotifServiceUpdate(repo)

	multi := h.NewNotificationHandler(service)
	solo := h.NewUpdateHandler(service_update)
	mux.Handle("/api/notifications/", middleware.NewMiddleWare(multi, authService))
	mux.Handle("/api/notifications/update/", middleware.NewMiddleWare(solo, authService))

	return mux
}
