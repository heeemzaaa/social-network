package routes

import (
	"database/sql"
	"net/http"

	ha "social-network/backend/handlers/auth"
	middleware "social-network/backend/middleware"
	ra "social-network/backend/repositories/auth"
	sa "social-network/backend/services/auth"
)

func SetAuthRoutes(mux *http.ServeMux, db *sql.DB) {
	authRepo := ra.NewAuthRepository(db)
	authService := sa.NewAuthService(authRepo)
	AuthHandler := ha.NewAuthHandler(authService)
	logoutHandler := ha.NewLogoutHandler(authService)
	loggedInHandler := ha.NewLoggedInHanlder(authService)
	// mux.Handle("/api/auth/", AuthHandler)
	mux.Handle("/api/loggedin", loggedInHandler)
	mux.Handle("/api/auth/logout", middleware.NewMiddleWare(logoutHandler, authService))
	mux.Handle("/api/auth/", middleware.NewLoginMiddleware(AuthHandler, authService))
}
