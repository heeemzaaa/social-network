package routes

import (
	"database/sql"
	"net/http"

	"social-network/backend/middleware"
)

func SetRoutes(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()
	mux, authService := SetAuthRoutes(mux, db)
	mux, notifService := SetNotificationsRoutes(mux, db, authService)
	mux, profileService := SetProfileRoutes(mux, db, authService, notifService)
	SetPostRoutes(mux, db, authService)
	SetGroupRoutes(mux, db, authService, profileService, notifService)
	SetChatRoutes(mux, db, authService)
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", middleware.NewMiddleWare(http.StripPrefix("/static/", fs), authService))
	return mux
}
