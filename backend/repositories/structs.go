package repositories

import (
	"database/sql"
)

type AppRepository struct {
	db *sql.DB
}

// NewPostRepository creates a new repository
func NewAppRepository(db *sql.DB) *AppRepository {
	return &AppRepository{db: db}
}
