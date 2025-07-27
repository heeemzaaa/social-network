package profile

import (
	"social-network/backend/models"
)

// custom posts to each users lil2assaf
func (s *ProfileService) GetPosts(profileID, authUserID, lastPostTime string) ([]models.Post, bool, *models.ErrorJson) {
	var posts []models.Post
	var isMine bool

	if profileID == "" || authUserID == ""  {
		return posts, false, &models.ErrorJson{Status: 400, Error: "Invalid data !"}
	}

	if profileID == authUserID {
		isMine = true
	}

	posts, err := s.repo.GetPosts(profileID, authUserID, lastPostTime, isMine)
	if err != nil {
		return posts, false, &models.ErrorJson{Status: err.Status, Error: err.Error}
	}

	access, errAccess := s.CheckProfileAccess(profileID, authUserID)
	if errAccess != nil {
		return []models.Post{}, false, &models.ErrorJson{Status: errAccess.Status, Error: errAccess.Error}
	}

	return posts, access, nil
}
