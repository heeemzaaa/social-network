package services

type PostsRepository struct {
	repo *PostsRepository
}

func NewPostsService(repo *PostsRepository) *PostsRepository {
	return &PostsRepository{repo: repo}
}
