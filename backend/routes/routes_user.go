package routes

import (
	"net/http"

	group "social-network/backend/handlers/group"
)

// auth 

func SetGroupRoutes(
	GroupHandler *group.GroupHanlder,
) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/api/groups", GroupHandler)
	mux.Handle("/api/groups/{id}", GroupHandler)

	return mux
}
