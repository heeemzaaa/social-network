package routes

import (
	"database/sql"
	"net/http"

	"social-network/backend/middleware"
	"social-network/backend/repositories/group"
	s "social-network/backend/services/images"

	"social-network/backend/repositories/post"
	"social-network/backend/repositories/profile"
	"social-network/backend/services/auth"
	groupService "social-network/backend/services/group"
	postService "social-network/backend/services/post"
	profileService "social-network/backend/services/profile"
)

func SetImageRoutes(
	mux *http.ServeMux,
	db *sql.DB,
	authService *auth.AuthService,
) {
	// ✅ Repositories
	repoGroups := group.NewGroupRepository(db)
	repoProfiles := profile.NewProfileRepository(db)
	repoPosts := post.NewPostRepository(db)

	// ✅ Services
	groupSvc := groupService.NewGroupService(repoGroups, nil) // pass real dependencies if needed
	profileSvc := profileService.NewProfileService(repoProfiles)
	postSvc := postService.NewPostService(repoPosts)

	// ✅ ServiceImages
	serviceImages := s.NewServiceImages(groupSvc, profileSvc, postSvc)

	// ✅ Middleware
	imgMiddleware := middleware.NewImagesMiddleware(nil, serviceImages)

	// ✅ Static file server
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle(
		"/static/",
		middleware.NewMiddleWare(imgMiddleware.AuthImageMiddleware(http.StripPrefix("/static/", fs)), authService,),
	)
}
