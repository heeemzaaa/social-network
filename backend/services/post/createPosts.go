package services

import (
	"log"

	"social-network/backend/models"
)

func (s *PostService) CreatePost(post *models.Post) *models.ErrorJson {
	if post.User.Id == "" || post.Id == "" {
		log.Println("Error creating the post , because of invalid data !")
		return &models.ErrorJson{Status: 400, Error: "Invalid data !"}
	}
	return s.repo.CreatePost(post)
}
