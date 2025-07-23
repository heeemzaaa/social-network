package routes

import (
	"database/sql"
	"net/http"
	
)

func SetRoutes(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()
	mux, authService := SetAuthRoutes(mux, db)
	mux, notifService := SetNotificationsRoutes(mux, db, authService)
	SetGroupRoutes(mux, db, authService)
	SetProfileRoutes(mux, db, authService, notifService)
	SetChatRoutes(mux, db, authService)
	
	return mux
}
