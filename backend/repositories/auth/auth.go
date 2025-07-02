package auth

import (
	"database/sql"
)

type AuthRepository struct {
	db *sql.DB
}

// NewPostRepository creates a new repository
func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

// func (repo *AuthRepository) AddUser() *models.ErrorJson {
// }

// func (repo *AuthRepository) UserExists() *models.ErrorJson {
// }
