package post

import (
	pr "social-network/backend/repositories/post"
)

type PostService struct {
	repo *pr.PostsRepository
}

func NewPostService(repo *pr.PostsRepository) *PostService {
	return &PostService{repo: repo}
}
