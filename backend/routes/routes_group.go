package routes

import (
	"net/http"

	group "social-network/backend/handlers/group"
)

// ##### routes i have to implement to all the users #####
// POST /groups/  create a group
// GET /groups?filter=owned
// GET  /groups?filter=availabe
// POST /groups/{id}   join a specific group
//  ##### routes i have to implement to all the user who is a member of a specific group  #####
// GET /groups/{id}/posts  (get the posts of a specific group)
// POST /groups/{id}/posts  (add a post to a specific group)
/**********************************************************/
// GET /groups/{id}/events  (get the events of a specific group)
// POST /groups/{id}/events  (add a event to a specific group)
/**********************************************************/
// GET /groups/{id}/posts/{id} (get the comments of a specific post of specific group)
// POST /groups/{id}/posts/{id}  (add a comment to a specific post of a specific group)
/***********************************************************/

func SetGroupRoutes(
	GroupHandler *group.GroupHanlder,
) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/api/groups", GroupHandler)
	mux.Handle("/api/groups/{id}", GroupHandler)

	return mux
}
