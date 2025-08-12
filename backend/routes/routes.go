package routes

import (
	"database/sql"
	"net/http"
)

func SetRoutes(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()
	mux, authService := SetAuthRoutes(mux, db)
	SetChatRoutes(mux, db, authService)
	// mux, notifService := SetNotificationsRoutes(mux, db, authService, chatService)
	mux, profileService := SetProfileRoutes(mux, db, authService)
	SetGroupRoutes(mux, db, authService, profileService)
	SetPostRoutes(mux, db, authService)

	SetImageRoutes(mux, db, authService)

	return mux
}
