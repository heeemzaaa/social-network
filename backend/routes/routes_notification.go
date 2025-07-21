package routes

import (
	"database/sql"
	"net/http"

	"social-network/backend/middleware"
	"social-network/backend/services/auth"

	h "social-network/backend/handlers/notification"
	ar "social-network/backend/repositories/auth"
	nr "social-network/backend/repositories/notification"
	ns "social-network/backend/services/notification"
)

func SetNotificationsRoutes(mux *http.ServeMux, db *sql.DB, authService *auth.AuthService) *http.ServeMux {
	repo := nr.NewNotifRepository(db)
	auth_repo := ar.NewAuthRepository(db)
	service := ns.NewNotifService(repo, auth_repo)
	service_update := ns.NewNotifServiceUpdate(repo)

	multi := h.NewNotificationHandler(service)
	solo := h.NewUpdateHandler(service_update)
	mux.Handle("/api/notifications/", middleware.NewMiddleWare(multi, authService))
	mux.Handle("/api/notifications/update/", middleware.NewMiddleWare(solo, authService))

	return mux
}
