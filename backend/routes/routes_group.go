package routes

import (
	"database/sql"
	"net/http"

	group "social-network/backend/handlers/group"
	"social-network/backend/middleware"
	gRepo "social-network/backend/repositories/group"
	authService "social-network/backend/services/auth"
	gService "social-network/backend/services/group"
	notificationService "social-network/backend/services/notification"
	profileService "social-network/backend/services/profile"
)

// cleaning the code and fixing the bugs !!

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
/**********************************************************/
// POST  /groups/{group_id}/react

/*******************************************************************************************/
// so the table created will be only one and not two
// separation of concerns ???
// for the requests

// POST   /groups/{group_id}/join-request  (the userID here is gotten from the context the one who is sending the request and the)
// one who will be processing it (the receiver_id) is the admin of the group
// DELETE  /groups/{group_id}/join-request  (the same here)
// GET /groups/{group_id}/join-request where the receiver is always the group admin
// for the requests acceptation  (to the admin of the group)
// POST  /groups/{group_id}/join-request/{user_id}
// DELETE  /groups/{group_id}/join-request/{user_id}
// But the admin has to have always access to be able to accept them (so the get ??????)
/********************************************************/
// invitations
// POST /groups/{group_id}/accept
// DELETE  /groups/{group_id}/decline



//  for approving or declining a request (we need to have the id of the request or the inviatation where )
//  it makes some sense because (having the group id and the user id is tooo much )
// for the delete especially 
func SetGroupRoutes(mux *http.ServeMux, db *sql.DB,
	authService *authService.AuthService,
	profileService *profileService.ProfileService,
	notifService *notificationService.NotificationService,
) {
	//  auth service
	// authRepo := ra.NewAuthRepository(db)
	// authService := sa.NewAuthServer(authRepo)
	// other setups
	groupRepo := gRepo.NewGroupRepository(db)
	groupService := gService.NewGroupService(groupRepo, profileService, notifService)
	GroupHandler := group.NewGroupHandler(groupService)
	GroupIDHandler := group.NewGroupIDHandler(groupService)
	GroupEventHandler := group.NewGroupEventHandler(groupService)
	GroupPostHandler := group.NewGroupPostHandler(groupService)
	GroupCommentHandler := group.NewGroupCommentHandler(groupService)
	GroupEventIDHandler := group.NewGroupEventIDHandler(groupService)
	GroupReactionHandler := group.NewReactionHandler(groupService)
	MembersHandler := group.NewMembersHanlder(groupService)
	RequestHandler := group.NewGroupRequestsHandler(groupService, notifService)
	DeclineApproveHandler := group.NewApproveDeclineReqHandler(groupService)
	InvitationsHandler := group.NewGroupInvitationHandler(groupService, notifService)
	AcceptRejectHanlder := group.NewAcceptRejectInvHandler(groupService)
	mux.Handle("/api/groups/{group_id}/events/{event_id}/", middleware.NewMiddleWare(GroupEventIDHandler, authService))
	mux.Handle("/api/groups/{group_id}/posts/{post_id}/comments/", middleware.NewMiddleWare(GroupCommentHandler, authService))
	mux.Handle("/api/groups/{group_id}/posts/", middleware.NewMiddleWare(GroupPostHandler, authService))
	mux.Handle("/api/groups/{group_id}/events/", middleware.NewMiddleWare(GroupEventHandler, authService))
	mux.Handle("/api/groups/{group_id}/react/like", middleware.NewMiddleWare(GroupReactionHandler, authService))
	mux.Handle("/api/groups/{group_id}/join-request/admin", middleware.NewMiddleWare(DeclineApproveHandler, authService))
	mux.Handle("/api/groups/{group_id}/invitations/", middleware.NewMiddleWare(InvitationsHandler, authService))
	mux.Handle("/api/groups/{group_id}/invitations/response", middleware.NewMiddleWare(AcceptRejectHanlder, authService))
	mux.Handle("/api/groups/{group_id}/join-request", middleware.NewMiddleWare(RequestHandler, authService))
	mux.Handle("/api/groups/{group_id}/members", middleware.NewMiddleWare(MembersHandler, authService))
	mux.Handle("/api/groups/{group_id}/", middleware.NewMiddleWare(GroupIDHandler, authService))
	mux.Handle("/api/groups/", middleware.NewMiddleWare(GroupHandler, authService))
}
