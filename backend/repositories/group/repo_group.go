package group

import (
	"database/sql"
)

type GroupRepository struct {
	db *sql.DB
}

// NewPostRepository creates a new repository
func NewAppRepository(db *sql.DB) *GroupRepository {
	return &GroupRepository{db: db}
}

func (repo *GroupRepository) CreateGroup() {
}

func (repo *GroupRepository) GetJoinedGroups() {
}

func (repo *GroupRepository) GetAvailableGroups() {
}

func (repo *GroupRepository) GetGroupById() {
}
