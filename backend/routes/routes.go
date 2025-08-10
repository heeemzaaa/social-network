package routes

import (
	"database/sql"
	"net/http"
)

func SetRoutes(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()
	mux, authService := SetAuthRoutes(mux, db)
	mux, chatService := SetChatRoutes(mux, db, authService)
	mux, notifService := SetNotificationsRoutes(mux, db, authService, chatService)
	mux, profileService := SetProfileRoutes(mux, db, authService, notifService)
	SetGroupRoutes(mux, db, authService, profileService, notifService)
	SetPostRoutes(mux, db, authService)

	SetImageRoutes(mux, db, authService)

	return mux
}
