package routes

import (
	"database/sql"
	"net/http"
)

func SetPostsRoutes(mux *http.ServeMux, db *sql.DB) *http.ServeMux {
	// postRepo := r.NewPostRepository(db)
	// postService := s.NewPostService(postRepo)
	// postHandler := h.SetPostsRoutes(postService)
	// fmt.Print(postHandler)
	// Example route to create a post (POST /api/posts)
	// mux.Handle("/api/posts", middleware.NewPostMiddleware(postHandler))

	// You can add more routes here, e.g.,
	// mux.Handle("/api/posts", middleware.NewAuthMiddleware(postHandler))

	return mux
}
