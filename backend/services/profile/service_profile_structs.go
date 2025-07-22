package profile

import (
	pr "social-network/backend/repositories/profile"
)

type ProfileService struct {
	repo *pr.ProfileRepository
}

func NewProfileService(repo *pr.ProfileRepository) *ProfileService {
	return &ProfileService{repo: repo}
}
