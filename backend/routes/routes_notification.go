package routes

import (
	"database/sql"
	"net/http"

	"social-network/backend/middleware"

	ra "social-network/backend/repositories/auth"
	"social-network/backend/services/auth"

	hn "social-network/backend/handlers/notification"
	sn "social-network/backend/services/notification"
	rn "social-network/backend/repositories/notification"

	rp "social-network/backend/repositories/profile"
	rg "social-network/backend/repositories/group"
)

func SetNotificationsRoutes(mux *http.ServeMux, db *sql.DB, authService *auth.AuthService, authRepo *ra.AuthRepository) (*http.ServeMux, *sn.NotificationService) {
	repo := rn.NewNotifRepository(db)
	profileRepo := rp.NewProfileRepository(db)
	groupRepo := rg.NewGroupRepository(db)

	service := sn.NewNotifService(repo, authRepo)
	serviceUpdate := sn.NewNotifServiceUpdate(repo, profileRepo, groupRepo)

	new := hn.NewNotificationHandler(service)
	update := hn.NewUpdateHandler(serviceUpdate)
	mux.Handle("/api/notifications/", middleware.NewMiddleWare(new, authService))
	mux.Handle("/api/notifications/update/", middleware.NewMiddleWare(update, authService))

	return mux, service
}
