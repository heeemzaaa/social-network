package routes

import (
	"database/sql"
	"net/http"
	r "social-network/backend/repositories/profile"
	s "social-network/backend/services/profile"
	h "social-network/backend/handlers/profile"
)

func SetProfileRoutes(mux *http.ServeMux, db *sql.DB) *http.ServeMux {
	repo := r.NewProfileRepository(db)
	service := s.NewProfileService(repo)
	handler := h.NewProfileHandler(service)
	mux.Handle("/api/profile/{id}/", handler)
	return mux
}
