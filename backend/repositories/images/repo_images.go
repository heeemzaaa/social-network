package images

import "database/sql"

type RepoImages struct {
	db *sql.DB
}

func NewRepoImages(db *sql.DB) *RepoImages {
	return &RepoImages{db: db}
}