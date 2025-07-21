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

func SetProfileRoutes(mux *http.ServeMux, db *sql.DB, authService *auth.AuthService) (*http.ServeMux, *s.ProfileService) {
	repo := r.NewProfileRepository(db)
	service := s.NewProfileService(repo)
	profile := h.NewProfileHandler(service)
	editProfile := h.NewEditProfileHandler(service)
	connections := h.NewUserConnectionHandler(service)
	response := h.NewResponseHandler(service)
	posts := h.NewProfilePostsHandler(service)
	actions := h.NewFollowActionHandler(service)


	mux.Handle("/api/profile/{id}/info/", middleware.NewMiddleWare(profile, authService))
	mux.Handle("/api/profile/{id}/edit/", middleware.NewMiddleWare(editProfile, authService))
	mux.Handle("/api/profile/{id}/connections/", middleware.NewMiddleWare(connections, authService))
	mux.Handle("/api/profile/{id}/response/", middleware.NewMiddleWare(response, authService))
	mux.Handle("/api/profile/{id}/data/", middleware.NewMiddleWare(posts, authService))
	mux.Handle("/api/profile/{id}/actions/", middleware.NewMiddleWare(actions, authService))
	
	return mux, service
}
