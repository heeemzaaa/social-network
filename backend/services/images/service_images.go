package images

import (
	"social-network/backend/repositories/group"
	"social-network/backend/repositories/images"
	"social-network/backend/repositories/profile"
)

type ServiceImages struct {
	repo *images.RepoImages
	groupRepo *group.GroupRepository
	profileRepo *profile.ProfileRepository
}

func NewServiceImages(repo *images.RepoImages) *ServiceImages {
	return &ServiceImages{repo: repo}
} 
