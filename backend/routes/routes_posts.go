package routes

import (
	"database/sql"
	"net/http"

	hp "social-network/backend/handlers"
	"social-network/backend/middleware"
	rp "social-network/backend/repositories"
	sp "social-network/backend/services"
	sa "social-network/backend/services/auth"
)

func SetPostRoutes(mux *http.ServeMux, db *sql.DB, authService *sa.AuthService) {

	postRepo := rp.NewPostRepository(db)
	postService := sp.NewPostService(postRepo)
	postHandler := hp.NewPostHandler(postService)

	mux.Handle("/api/posts", middleware.NewMiddleWare(postHandler, authService)) 
}
