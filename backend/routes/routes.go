package routes

import (
	"database/sql"
	"net/http"
)

func SetRoutes(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()
	mux, authService, authRepo := SetAuthRoutes(mux, db)
	mux, chatService := SetChatRoutes(mux, db, authService)
	mux, notifService := SetNotificationsRoutes(mux, db, authService, authRepo, chatService)
	mux, profileService := SetProfileRoutes(mux, db, authService, notifService)
	SetPostRoutes(mux, db, authService)
	SetGroupRoutes(mux, db, authService, profileService, notifService)

	return mux
}
