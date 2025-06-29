package profile

import "database/sql"

type ProfileRepository struct {
	db *sql.DB
}

func NewProfileRepository(db *sql.DB) *ProfileRepository {
	return &ProfileRepository{db:db}
}


// hna ghangeter user informations
func (repo *ProfileRepository) GetProfileData() {
	
}

// hna ghangeter l posts kamlin dyal dak luser
func (repo *ProfileRepository) GetPosts() {

}

// hna ghangeter l followers
func (repo *ProfileRepository) GetFollowers() {

}

func (repo *ProfileRepository) GetFollowed() {

}

func (repo *ProfileRepository) GetAuthorization() {

}

func (repo *ProfileRepository) GetNumberOfGroups() {

}

func (repo *ProfileRepository) UpdateUserData() {

}



