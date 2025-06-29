package routes

import (
	"net/http"

	group "social-network/backend/handlers/group"
)

func SetUserRoutes(
	GroupHandler *group.GroupHanlder,
) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/auth/logout", GroupHandler)
	mux.Handle("/api/groups/{id}", GroupHandler)
	mux.Handle("/api/groups/{id}", GroupHandler)

	return mux
}
