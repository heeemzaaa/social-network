package routes

import (
	"database/sql"
	"net/http"
)

func SetPostsRoutes(mux *http.ServeMux, db *sql.DB) *http.ServeMux {
	// repo := r.NewPostsRepository(db)
	// service := s.NewPostsService(repo)
	// handler := h.NewPostsHandler(service)
	return mux
}
