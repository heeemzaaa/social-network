package profile

import "database/sql"

type ProfileRepository struct {
	db *sql.DB
}

func NewProfileRepository(db *sql.DB) *ProfileRepository {
	return &ProfileRepository{db:db}
}

func (repo *ProfileRepository) GetProfileData() {
	
}

