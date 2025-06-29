package routes

import (
	"net/http"

	"social-network/backend/handlers/auth"
)

func SetAuthRoutes(
	AuthHandler *auth.AuthHandler,
) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/auth/login", AuthHandler)
	mux.Handle("/auth/register", AuthHandler)
	mux.Handle("/auth/is_logged_in", AuthHandler)
	mux.Handle("/auth/logout", AuthHandler)
	return mux
}
