package routes

import (
	"database/sql"
	"net/http"
)

func SetRoutes(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()
	
	mux  = SetAuthRoutes(mux, db)

	return mux
}
