package routes

import (
	"database/sql"
	"net/http"

	"social-network/backend/middleware"

	ar "social-network/backend/repositories/auth"
	"social-network/backend/services/auth"

	hn "social-network/backend/handlers/notification"
	ns "social-network/backend/services/notification"
	nr "social-network/backend/repositories/notification"

	pr "social-network/backend/repositories/profile"
)

func SetNotificationsRoutes(mux *http.ServeMux, db *sql.DB, authService *auth.AuthService) (*http.ServeMux, *ns.NotificationService) {
	repo := nr.NewNotifRepository(db)
	auth_repo := ar.NewAuthRepository(db)
	profile_repo := pr.NewProfileRepository(db)

	service := ns.NewNotifService(repo, auth_repo)
	service_update := ns.NewNotifServiceUpdate(repo, profile_repo)

	new := hn.NewNotificationHandler(service)
	update := hn.NewUpdateHandler(service_update)
	mux.Handle("/api/notifications/", middleware.NewMiddleWare(new, authService))
	mux.Handle("/api/notifications/update/", middleware.NewMiddleWare(update, authService))

	return mux, service
}
