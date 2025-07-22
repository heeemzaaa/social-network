package profile

import (
	"fmt"
	"log"

	"social-network/backend/models"
)

// here I will check if the user is following the profile or not
func (repo *ProfileRepository) IsFollower(userID string, authUserID string) (bool, *models.ErrorJson) {
	var exist int
	query := `SELECT EXISTS
		(SELECT 1 FROM followers WHERE userID = ? AND followerID = ? LIMIT 1)
	`

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		log.Println("Error preparing the query the check if the user is a follower :", err)
		return false, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	err = stmt.QueryRow(userID, authUserID).Scan(&exist)
	if err != nil {
		log.Printf("Error checking if the user following the chosing profile: %v", err)
		return false, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}

	return exist == 1, nil
}

// here I will check if the user have a private profile or a public one
func (repo *ProfileRepository) Visibility(userID string) (string, *models.ErrorJson) {
	var visibility string
	query := `SELECT visibility FROM users WHERE userID = ?`

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		log.Println("Error preparing the query to see the visibility: ", err)
		return "", &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	err = stmt.QueryRow(userID).Scan(&visibility)
	if err != nil {
		log.Println("Error checking visibility: ", err)
		return "", &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	return visibility, nil
}

// here I will check if the user has the authorization to access my profile
// is the user getting the profile is me ?
// is my account private, if yes , is he a follower?
// is my account public ?
func (repo *ProfileRepository) CheckProfileAccess(userID string, authUserID string) (bool, *models.ErrorJson) {
	visibility, err := repo.Visibility(userID)
	if err != nil {
		return false, &models.ErrorJson{Status: err.Status, Error: err.Error}
	}

	// If profile is public, good to go
	if visibility == "public" {
		return true, nil
	}

	// If profile is private, ncheckiw wach follower
	if visibility == "private" {
		var isFollower bool
		isFollower, err := repo.IsFollower(userID, authUserID)
		if err != nil {
			return false, &models.ErrorJson{Status: err.Status, Error: err.Error}
		}
		return isFollower, nil
	}

	return false, &models.ErrorJson{Status: 400, Error: "Invalid visibility type !"}
}

func (repo *ProfileRepository) IsRequested(profileID string, authUserID string) (bool, *models.ErrorJson) {
	isRequested := false
	query := `SELECT EXISTS (SELECT 1 FROM follow_requests WHERE userID = ? AND requestorID = ? LIMIT 1)`

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		log.Println("Error preparing the query to check if there's a request: ", err)
		return false, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	err = stmt.QueryRow(profileID, authUserID).Scan(&isRequested)
	if err != nil {
		log.Println("Error checking if there's a request: ", err)
		return false, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}

	return isRequested, nil
}
