package services

import (
	"social-network/backend/models"
	pr "social-network/backend/repositories/post"

	"github.com/google/uuid"
)

type PostService struct {
	repo *pr.PostsRepository
}

func NewPostService(repo *pr.PostsRepository) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) CreatePost(post *models.Post) error {
	return s.repo.CreatePost(post)
}

func (s *PostService) GetAllPosts(userID uuid.UUID) ([]models.Post, error) {
	return s.repo.GetAllPosts(userID)
}

func (s *PostService) GetPostByID(postID string) (models.Post, error) {
	return s.repo.GetPostByID(postID)
}

func (s *PostService) HandleLike(postID uuid.UUID, userID uuid.UUID) error {
	return s.repo.HandleLike(postID, userID)
}
