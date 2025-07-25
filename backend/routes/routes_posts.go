package routes

import (
	"database/sql"
	"net/http"

	hp "social-network/backend/handlers/post"
	"social-network/backend/middleware"
	rp "social-network/backend/repositories/post"
	sp "social-network/backend/services/post"
	sa "social-network/backend/services/auth"
)

func SetPostRoutes(mux *http.ServeMux, db *sql.DB, authService *sa.AuthService) {
	postRepo := rp.NewPostRepository(db)
	postService := sp.NewPostService(postRepo)
	postHandler := hp.NewPostHandler(postService)

	mux.Handle("/api/posts", middleware.NewMiddleWare(postHandler, authService))
	mux.Handle("/api/posts/like/{id}", middleware.NewMiddleWare(postHandler, authService))
	mux.Handle("/api/posts/comment", middleware.NewMiddleWare(postHandler, authService))
	mux.Handle("/api/posts/comments/{id}", middleware.NewMiddleWare(postHandler, authService))
}
