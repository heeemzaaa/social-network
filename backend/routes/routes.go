package routes

import (
	"database/sql"
	"net/http"
)

func SetRoutes(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()
	mux, authService := SetAuthRoutes(mux, db)
	mux, profileService := SetProfileRoutes(mux, db, authService)
	SetGroupRoutes(mux, db, authService, profileService)
	SetChatRoutes(mux, db, authService)

	return mux
}
