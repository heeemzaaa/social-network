package routes

import (
	"database/sql"
	"net/http"

	hp "social-network/backend/handlers/post"
	"social-network/backend/middleware"
	rp "social-network/backend/repositories/post"
	sa "social-network/backend/services/auth"
	sp "social-network/backend/services/post"
)

// SetPostRoutes registers post-related routes
func SetPostRoutes(mux *http.ServeMux, db *sql.DB, authService *sa.AuthService) {
	postRepo := rp.NewPostRepository(db)          // repo layer
	postService := sp.NewPostService(postRepo)    // service layer
	postHandler := hp.NewPostHandler(postService) // handler layer
	// endpoints -> middleware to check authentification -> args the authService to check the session + the post handler
	mux.Handle("/api/posts", middleware.NewMiddleWare(postHandler, authService))               // posts
	mux.Handle("/api/posts/like/{id}", middleware.NewMiddleWare(postHandler, authService))     // like post
	mux.Handle("/api/posts/comment", middleware.NewMiddleWare(postHandler, authService))       // comment
	mux.Handle("/api/posts/comments/{id}", middleware.NewMiddleWare(postHandler, authService)) // comment
}
