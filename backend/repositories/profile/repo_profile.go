package profile

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"social-network/backend/models"
)

type ProfileRepository struct {
	db *sql.DB
}

func NewProfileRepository(db *sql.DB) *ProfileRepository {
	return &ProfileRepository{db: db}
}

// here I will get the userID based based on the sessionToken to pass it to other functions
func (repo *ProfileRepository) GetID(sessionToken string) (string, error) {
	var userID string
	query := `SELECT userID from sessions WHERE sessionToken = ?`
	err := repo.db.QueryRow(query, sessionToken).Scan(&userID)
	if err != nil {
		log.Println("Error getting the userID from the database:", err)
		return "", fmt.Errorf("error getting the userID from the database: %v", err)
	}
	return userID, nil
}

// here I will check if the user is following the profile or not
func (repo *ProfileRepository) IsFollower(userID string, authUserID string) (bool, error) {
	var exist int
	query := `SELECT EXISTS
		(SELECT 1 FROM followers WHERE userID = ? AND followerID = ? LIMIT 1)
	`

	err := repo.db.QueryRow(query, userID, authUserID).Scan(&exist)
	if err != nil {
		log.Printf("Error checking if the user following the chosing profile: %v", err)
		return false, fmt.Errorf("%v", err)
	}

	return exist == 1, nil
}

// here I will check if the user have a private profile or a public one
func (repo *ProfileRepository) Visibility(userID string) (string, error) {
	fmt.Println(userID)
	var visibility string
	query := `SELECT visibility FROM users WHERE userID = ?`
	err := repo.db.QueryRow(query, userID).Scan(&visibility)
	if err != nil {
		return "", fmt.Errorf("error fetching user visibility: %v", err)
	}
	return visibility, nil
}

// here I will insert the follow request if the user has a public account
func (repo *ProfileRepository) FollowDone(userID string, authUserID string) error {
	query := `INSERT INTO followers (userID, followerID) VALUES(?,?) `
	_, err := repo.db.Exec(query, userID, authUserID)
	if err != nil {
		return fmt.Errorf("error inserting the data into the followers table:%v", err)
	}
	return nil
}

// here I will just insert the request into the table of the followrequests
func (repo *ProfileRepository) FollowPrivate(userID string, authUserID string) error {
	sentAt := time.Now().UTC().Format(time.RFC3339)
	query := `INSERT INTO follow_requests (userID, requestorID, sent_at) VALUES(?, ?, ?) `
	_, err := repo.db.Exec(query, userID, authUserID, sentAt)
	if err != nil {
		return fmt.Errorf("error inserting the data into the follow_requests table:%v", err)
	}
	return nil
}

// delete the request works in both cases , accept and decline
func (repo *ProfileRepository) DeleteRequest(userID string, authUserID string) error {
	query := `DELETE FROM follow_requests WHERE userID = ? AND requestorID = ?`
	_, err := repo.db.Exec(query, userID, authUserID)
	if err != nil {
		return fmt.Errorf("error deleting the data into the follow_requests table: %v", err)
	}
	return nil
}

// if the user accepted the follow request
// I will delete it from the table and add him to the following list
func (repo *ProfileRepository) AcceptedRequest(userID string, authUserID string) error {
	err := repo.DeleteRequest(userID, authUserID)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	err = repo.FollowDone(userID, authUserID)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	return nil
}

// here I will just delete it from the table of the follow requests
func (repo *ProfileRepository) RejectedRequest(userID string, authUserID string) error {
	err := repo.DeleteRequest(userID, authUserID)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	return nil
}

// here we unfollow the user chosen , we delete it from the table of followers
func (repo *ProfileRepository) Unfollow(userID string, authUserID string) error {
	query := `DELETE FROM followers WHERE userID = ? AND followerID = ?`
	_, err := repo.db.Exec(query, userID, authUserID)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	return nil
}

// here we accept all the requests where the userID
// change its visibility to public only
func (repo *ProfileRepository) AcceptAllrequest(userID string) error {
	var users []models.User
	query := `SELECT requestorID FROM follow_requests WHERE userID = ?`
	rows, err := repo.db.Query(query, userID)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id)
		if err != nil {
			return fmt.Errorf("%v", err)
		}
		users = append(users, user)
	}

	for _, user := range users {
		query := `INSERT INTO followers (userID, followerID) VALUES(?,?)`
		_, err := repo.db.Exec(query, userID, user.Id)
		if err != nil {
			return fmt.Errorf("%v", err)
		}
	}

	query = `DELETE FROM follow_requests WHERE userID = ?`
	_, err = repo.db.Exec(query, userID)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	return nil
}

// here I will check if the user has the authorization to access my profile
// is the user getting the profile is me ?
// is my account private, if yes , is he a follower?
// is my account public ?
func (repo *ProfileRepository) CheckProfileAccess(userID string, authUserID string) (bool, error) {
	visibility, err := repo.Visibility(userID)
	if err != nil {
		return false, fmt.Errorf("%v", err)
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
			return false, fmt.Errorf("%v", err)
		}
		return isFollower, nil
	}

	return false, fmt.Errorf("invalid visibility status: %s", visibility)
}

// here I will get the data of the user
func (repo *ProfileRepository) GetProfileData(profileID string, access bool) (*models.Profile, error) {
	var profile models.Profile
	var query string

	// user with no access , we decided ana o ayoub nbiyno lihom gher
	// fullName, nickname, avatar, still need to check if that's okay
	if !access {
		query = "SELECT firstName, lastName, nickname, avatarPath FROM users WHERE userID = ?"
		err := repo.db.QueryRow(query, profileID).Scan(
			&profile.User.FirstName,
			&profile.User.LastName,
			&profile.User.Nickname,
			&profile.User.ImagePath,
		)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, fmt.Errorf("UserNotFoundError: no user with ID %v", profileID)
			}
			log.Printf("RowScanError in backend/repositories/profile/repo_profile.go/GetProfileData: %v", err)
			return nil, fmt.Errorf("RowScanError: failed to scan user profile: %w", err)
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
	err := repo.db.QueryRow(query, profileID).Scan(
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
			return nil, fmt.Errorf("UserNotFoundError: no user with ID %v", profileID)
		}
		log.Printf("RowScanError in backend/repositories/profile/repo_profile.go/GetProfileData: %v", err)
		return nil, fmt.Errorf("RowScanError: failed to scan user profile and counts: %w", err)
	}
	profile.User.Id = profileID
	return &profile, nil
}

// here I will get all the user's posts
func (repo *ProfileRepository) GetPosts(profileID string, userID string) (*[]models.Post, error) {
	return nil, nil
}

// here I will get my followers as a user
func (repo *ProfileRepository) GetFollowers(profileID string) (*[]models.User, error) {
	var query string
	var users []models.User

	query = `SELECT u.userID, u.firstName, u.lastName, u.nickname, u.avatarPath  FROM followers f
			JOIN users u ON f.followerID = u.userID
			WHERE f.userID = ?
	
	`
	rows, err := repo.db.Query(query, profileID)
	if err != nil {
		log.Printf("RowScanError in backend/repositories/profile/repo_profile.go/GetFollowers: %v", err)
		return nil, fmt.Errorf("RowScanError: failed to get the followersID: %w", err)

	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Nickname, &user.ImagePath)
		if err != nil {
			log.Printf("RowScanError in backend/repositories/profile/repo_profile.go/GetFollowers: %v", err)
			return nil, fmt.Errorf("RowScanError: failed to scan follower: %w", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Printf("RowsError in backend/repositories/profile/repo_profile.go/GetFollowers: %v", err)
		return nil, fmt.Errorf("RowsError: failed to iterate followers: %w", err)
	}

	return &users, nil
}

// here I will get the followed users
func (repo *ProfileRepository) GetFollowing(profileID string) (*[]models.User, error) {
	var query string
	var users []models.User

	query = `SELECT u.userID, u.firstName, u.lastName, u.nickname, u.avatarPath 
			FROM followers f
			JOIN users u ON f.userID = u.userID
			WHERE followerID = ?
	
	`
	rows, err := repo.db.Query(query, profileID)
	if err != nil {
		return nil, fmt.Errorf("%v", err)

	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Nickname, &user.ImagePath)
		if err != nil {
			return nil, fmt.Errorf("%v", err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return &users, nil
}

// here I will update anything in the user data
func (repo *ProfileRepository) UpdateProfileData(user *models.User) (*models.User, error) {
	return nil, nil
}

// here I will update the visibility of a private status
func (repo *ProfileRepository) ToPublicAccount(userID string) error {
	query := `UPDATE users SET visibility = ? WHERE userID = ?`
	_, err := repo.db.Exec(query, "public", userID)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	return nil
}

// change the visibility to private
func (repo *ProfileRepository) ToPrivateAccount(userID string) error {
	query := `UPDATE users SET visibility = ? WHERE userID = ?`
	_, err := repo.db.Exec(query, "private", userID)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	return nil
}
