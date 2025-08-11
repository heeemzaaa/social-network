package routes

import (
	"database/sql"
	"net/http"

	"social-network/backend/middleware"

	hc "social-network/backend/handlers/chat"
	sa "social-network/backend/services/auth"

	hn "social-network/backend/handlers/notification"
	rn "social-network/backend/repositories/notification"
	sn "social-network/backend/services/notification"

	rg "social-network/backend/repositories/group"

	rp "social-network/backend/repositories/profile"
	sp "social-network/backend/services/profile"
)

// SetNotificationsRoutes sets up the routes for notifications and returns the updated mux and notification service.
func SetNotificationsRoutes(mux *http.ServeMux, db *sql.DB, authService *sa.AuthService, chatServer *hc.ChatServer) (*http.ServeMux, *sn.NotificationService) {
	repo := rn.NewNotifRepository(db)
	profileRepo := rp.NewProfileRepository(db)
	profileService := sp.NewProfileService(profileRepo)

	groupRepo := rg.NewGroupRepository(db)

	service := sn.NewNotifService(repo, authService, profileService, groupRepo, chatServer)

	new := hn.NewNotificationHandler(service)
	update := hn.NewUpdateNotificationHandler(service)
	mux.Handle("/api/notifications/", middleware.NewMiddleWare(new, authService))
	mux.Handle("/api/notifications/update/", middleware.NewMiddleWare(update, authService))

	return mux, service
}
