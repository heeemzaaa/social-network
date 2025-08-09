package images

import (
	"social-network/backend/services/group"
	"social-network/backend/services/post"
	"social-network/backend/services/profile"
)

type ServiceImages struct {
	groupService   *group.GroupService
	profileService *profile.ProfileService
	postService    *post.PostService
}

func NewServiceImages(
	groupService *group.GroupService,
	profileService *profile.ProfileService,
	postService *post.PostService,
) *ServiceImages {
	return &ServiceImages{
		groupService:   groupService,
		profileService: profileService,
		postService:    postService,
	}
}
