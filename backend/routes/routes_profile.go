package routes

import (
	"database/sql"
	"net/http"

	h "social-network/backend/handlers/profile"
	r "social-network/backend/repositories/profile"
	auth "social-network/backend/services/auth"
	s "social-network/backend/services/profile"
	middleware  "social-network/backend/middleware"
)

func SetProfileRoutes(mux *http.ServeMux, db *sql.DB, authService *auth.AuthService) *http.ServeMux {
	repo := r.NewProfileRepository(db)
	service := s.NewProfileService(repo)
	handler := h.NewProfileHandler(service)
	mux.Handle("/api/profile/{id}/", middleware.NewMiddleWare(handler, authService))
	return mux
}
