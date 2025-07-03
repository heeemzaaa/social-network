package routes

import (
	"database/sql"
	"net/http"

	ha "social-network/backend/handlers/auth"
	ra "social-network/backend/repositories/auth"
	sa "social-network/backend/services/auth"
)

func SetAuthRoutes(mux *http.ServeMux, db *sql.DB) *http.ServeMux {
	authRepo := ra.NewAuthRepository(db)
	authService := sa.NewAuthServer(authRepo)
	AuthHandler := ha.NewAuthHandler(authService)

	mux.Handle("/api/auth/", AuthHandler)
	return mux
}
