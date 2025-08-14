package repositories

import (
	"database/sql"
)
// PostsRepository handles database operations for posts
// in purposr of attaching methods to it that all share the same database connection
type PostsRepository struct {
	db *sql.DB
}

// NewPostRepository creates a new PostsRepository with a database connection (like an instance)
func NewPostRepository(db *sql.DB) *PostsRepository {
	return &PostsRepository{db: db}
}
