package profile

import (
	"database/sql"
	"fmt"
	"log"

	"social-network/backend/models"
)

// here I will get my followers as a user
func (repo *ProfileRepository) GetFollowers(profileID string) ([]models.User, *models.ErrorJson) {
	var query string
	users := []models.User{}

	query = `SELECT u.userID, u.firstName, u.lastName, u.nickname, u.avatarPath  FROM followers f
			JOIN users u ON f.followerID = u.userID
			WHERE f.userID = ?
	`

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		log.Println("Error preparing a query to get the followers: ", err)
		return users, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
		defer stmt.Close()


	rows, err := stmt.Query(profileID)
	if err != nil {
		if err == sql.ErrNoRows {
			return users, nil
		}

		log.Println("Error getting the followers: ", err)
		return users, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}

	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Nickname, &user.ImagePath)
		if err != nil {
			log.Println("Error scanning the followers: ", err)
			return []models.User{}, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error in the whole process of scan => in get followers: ", err)
		return []models.User{}, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}

	return users, nil
}

// here I will get the followed users
func (repo *ProfileRepository) GetFollowing(profileID string) ([]models.User, *models.ErrorJson) {
	var query string
	users := []models.User{}

	query = `SELECT u.userID, u.firstName, u.lastName, u.nickname, u.avatarPath 
			FROM followers f
			JOIN users u ON f.userID = u.userID
			WHERE followerID = ?
	
	`

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		log.Println("Error preparing the query to get the following users: ", err)
		return users, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	rows, err := stmt.Query(profileID)
	if err != nil {
		if err == sql.ErrNoRows {
			return users, nil
		}

		log.Println("Error getting the following users: ", err)
		return []models.User{}, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Nickname, &user.ImagePath)
		if err != nil {
			log.Println("Error scaning the following user: ", err)
			return []models.User{}, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error scanning all following users: ", err)
		return []models.User{}, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}

	return users, nil
}
