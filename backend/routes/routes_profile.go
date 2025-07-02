package routes

import (
	"net/http"
	"social-network/backend/handlers/profile"
)

func SetProfileRoutes(
	ProfileHandler *profile.ProfileHandler,
) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/api/profile/{id}/", ProfileHandler)
	return mux
}
