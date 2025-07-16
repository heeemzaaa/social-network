package routes

import (
	"database/sql"
	"net/http"

	group "social-network/backend/handlers/group"
	"social-network/backend/middleware"
	gRepo "social-network/backend/repositories/group"
	authService "social-network/backend/services/auth"
	gService "social-network/backend/services/group"
)

// ##### routes i have to implement to all the users #####
// POST /groups/  create a group done
// GET /groups?filter=owned   done
// GET  /groups?filter=availabe done
// GET /groups?filter=joined done
// POST /groups/{id}   join a specific group done
// GET /groups/{id}  get the general info of a specific group (title description and so on )
//  ##### routes i have to implement to all the user who is a member of a specific group  #####
// GET /groups/{id}/events  (get the events of a specific group)
// POST /groups/{id}/events  (add a event to a specific group)
// almost done !!
/**********************************************************/
// GET /groups/{id}/posts  (get the posts of a specific group)
// POST /groups/{id}/posts  (add a post to a specific group)
/**********************************************************/
// GET /groups/{id}/posts/{post_id}/comments (get the comments of a specific post of specific group)
// POST /groups/{id}/posts/{post_id}/comments  (add a comment to a specific post of a specific group)
/***********************************************************/
// GET /api/groups/{group_id}/events/{event-id}/  (get the details of a specific event of a specific group)
// POST /api/groups/{group_id}/events/{event-id}/ (show interest to  an event to a specific group)
/**********************************************************/
// GET /groups/{group_id}/members  GET the members of a specific group
// GET /groups/{group_id}/members/{id} to the profile of a user of a specific group
/**********************************************************/
// POST  /groups/{group_id}/react

func SetGroupRoutes(mux *http.ServeMux, db *sql.DB, authService *authService.AuthService) {
	//  auth service
	// authRepo := ra.NewAuthRepository(db)
	// authService := sa.NewAuthServer(authRepo)
	// other setups
	groupRepo := gRepo.NewGroupRepository(db)
	groupService := gService.NewGroupService(groupRepo)
	GroupHandler := group.NewGroupHandler(groupService)
	GroupIDHandler := group.NewGroupIDHandler(groupService)
	GroupEventHandler := group.NewGroupEventHandler(groupService)
	GroupPostHandler := group.NewGroupPostHandler(groupService)
	GroupCommentHandler := group.NewGroupCommentHandler(groupService)
	GroupEventIDHandler := group.NewGroupEventIDHandler(groupService)
	GroupReactionHandler := group.NewReactionHandler(groupService)
	mux.Handle("/api/groups/{group_id}/events/{event_id}/", middleware.NewMiddleWare(GroupEventIDHandler, authService))
	mux.Handle("/api/groups/{group_id}/posts/{post_id}/comments/", middleware.NewMiddleWare(GroupCommentHandler, authService))
	mux.Handle("/api/groups/{group_id}/posts/", middleware.NewMiddleWare(GroupPostHandler, authService))
	mux.Handle("/api/groups/{group_id}/events/", middleware.NewMiddleWare(GroupEventHandler, authService))
	mux.Handle("/api/groups/{group_id}/react/like", middleware.NewMiddleWare(GroupReactionHandler, authService))
	mux.Handle("/api/groups/{group_id}/", middleware.NewMiddleWare(GroupIDHandler, authService))
	mux.Handle("/api/groups/", middleware.NewMiddleWare(GroupHandler, authService))
}
