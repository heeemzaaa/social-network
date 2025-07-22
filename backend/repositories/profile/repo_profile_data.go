package profile

import (
	"database/sql"
	"fmt"
	"log"

	"social-network/backend/models"
)

// here I will get the data of the user
func (repo *ProfileRepository) GetProfileData(profileID string, access bool) (*models.Profile, *models.ErrorJson) {
	var profile models.Profile
	var query string

	// user with no access , we decided ana o ayoub nbiyno lihom gher
	// fullName, nickname, avatar, still need to check if that's okay
	if !access {
		query = "SELECT firstName, lastName, nickname, avatarPath FROM users WHERE userID = ?"

		stmt, err := repo.db.Prepare(query)
		if err != nil {
			log.Println("Error preparing the query to get the profile data: ", err)
			return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
		}
		defer stmt.Close()

		err = stmt.QueryRow(profileID).Scan(
			&profile.User.FirstName,
			&profile.User.LastName,
			&profile.User.Nickname,
			&profile.User.ImagePath,
		)
		if err != nil {
			if err == sql.ErrNoRows {
				log.Println("Error user not found to get the profile data!")
				return nil, &models.ErrorJson{Status: 404, Error: "User not found !"}
			}
			log.Println("Error getting the data of a user: ", err)
			return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
		}
		profile.User.Id = profileID
		return &profile, nil
	}

	// user with access , all the data
	query = `
		SELECT 
    		u.email, u.firstName, u.lastName, u.birthDate, u.nickname, u.avatarPath, u.aboutMe,
    		COUNT(DISTINCT p.postID) AS post_count,
    		COUNT(DISTINCT f.followerID) AS follower_count,
    		COUNT(DISTINCT fl.userID) AS following_count,
    		COUNT(DISTINCT g.groupID) AS group_count
		FROM users u
			LEFT JOIN posts p ON u.userID = p.userID
			LEFT JOIN followers f ON u.userID = f.userID
			LEFT JOIN followers fl ON u.userID = fl.followerID
			LEFT JOIN group_membership g ON u.userID = g.userID
		WHERE u.userID = ?
		GROUP BY u.userID
	`

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		log.Println("Error preparing the query to get the profile data: ", err)
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	err = stmt.QueryRow(profileID).Scan(
		&profile.User.Email,
		&profile.User.FirstName,
		&profile.User.LastName,
		&profile.User.BirthDate,
		&profile.User.Nickname,
		&profile.User.ImagePath,
		&profile.User.AboutMe,
		&profile.NumberOfPosts,
		&profile.NumberOfFollowers,
		&profile.NumberOfFollowing,
		&profile.NumberOfGroups,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Error user not found to get the profile data!")
			return nil, &models.ErrorJson{Status: 404, Error: "User not found !"}
		}
		log.Println("Error getting the data of the user: ", err)
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}

	profile.User.Id = profileID
	return &profile, nil
}
