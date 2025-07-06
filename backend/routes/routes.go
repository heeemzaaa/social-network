package routes

import (
	"database/sql"
	"net/http"
	
)

func SetRoutes(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()
	SetAuthRoutes(mux, db)
	SetGroupRoutes(mux, db)

	return mux
}
