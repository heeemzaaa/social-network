package services

import (
	pr "social-network/backend/repositories/post"
)
// PostService provides business logic for managing posts.
type PostService struct {
	repo *pr.PostsRepository
}

// NewPostService creates a new PostService with the given repository to acces to it ... give it the repo to use create post addpost ect...
func NewPostService(repo *pr.PostsRepository) *PostService {
	return &PostService{repo: repo}
}
