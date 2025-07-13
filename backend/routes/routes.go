package routes

import (
	"database/sql"
	"net/http"
	
)

func SetRoutes(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()
	mux, authService := SetAuthRoutes(mux, db)
	SetGroupRoutes(mux, db, authService)

	return mux
}
