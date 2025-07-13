package repositories

import "database/sql"

type PostsRepository struct {
	db *sql.DB
}

func NewPostsRepository(db *sql.DB) *PostsRepository {
	return &PostsRepository{db: db}
}
