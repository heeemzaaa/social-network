package profile

import (
	"database/sql"
	"fmt"
	"log"
	"social-network/backend/models"
)

type ProfileRepository struct {
	db *sql.DB
}

func NewProfileRepository(db *sql.DB) *ProfileRepository {
	return &ProfileRepository{db: db}
}

//here I will get the userID based based on the sessionID to pass it to other functions
func (repo *ProfileRepository) GetID(sessionID string) (string, error) {
	var userID string
	query := `SELECT userID from sessions WHERE sessionID = ?`
	err := repo.db.QueryRow(query, sessionID).Scan(&userID)
	if err != nil {
		log.Println("Error getting the userID from the database:" , err)
		return "" , fmt.Errorf("error getting the userID from the database: %v" , err)
	}
	return userID, nil
}

// here I will check if the user has the authorization to access my profile
// is the user getting the profile is me ?
// is my account private, if yes , is he a follower?
// is my account public ?
func (repo *ProfileRepository) CheckProfileAccess(userID string, authUserID string) (bool, error) {
	if userID == authUserID {
		return true, nil
	}

	var visibility string
	query := `SELECT visibility FROM users WHERE userID = ?`
	err := repo.db.QueryRow(query, userID).Scan(&visibility)
	if err != nil {
		return false, fmt.Errorf("error fetching user visibility: %v", err)
	}

	// If profile is public, good to go
	if visibility == "public" {
		return true, nil
	}

	// If profile is private, ncheckiw wach follower
	if visibility == "private" {
		var isFollower int
		query = `
		SELECT EXISTS 
		(SELECT 1 FROM followers WHERE userID = ? AND followerID = ? LIMIT 1)
		`
		err := repo.db.QueryRow(query, userID, authUserID).Scan(&isFollower)
		if err != nil {
			return false, fmt.Errorf("error checking follower status: %v", err)
		}
		return isFollower == 1, nil
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
			&profile.User.AvatarPath,
		)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, fmt.Errorf("UserNotFoundError: no user with ID %v", profileID)
			}
			log.Printf("RowScanError in backend/repositories/profile/repo_profile.go/GetProfileData: %v", err)
			return nil, fmt.Errorf("RowScanError: failed to scan user profile: %w", err)
		}
		return &profile, nil
	}

	// user with access , all the data
	query = `
        SELECT
            email, firstName, lastName, birthDate, nickname, avatarPath, aboutMe,
            (SELECT COUNT(*) FROM posts WHERE userID = ?) AS post_count,
            (SELECT COUNT(*) FROM followers WHERE userID = ?) AS follower_count,
            (SELECT COUNT(*) FROM followers WHERE followerID = ?) AS following_count,
            (SELECT COUNT(*) FROM groups WHERE userID = ?) AS group_count
        FROM users
        WHERE userID = ?
    `
	err := repo.db.QueryRow(query, profileID, profileID, profileID, profileID, profileID).Scan(
		&profile.User.Email,
		&profile.User.FirstName,
		&profile.User.LastName,
		&profile.User.BirthDate,
		&profile.User.Nickname,
		&profile.User.AvatarPath,
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
	return &profile, nil
}

// here I will get all the user's posts
func (repo *ProfileRepository) GetPosts(profileID string) (*[]models.Post, error) {
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
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Nickname, &user.AvatarPath)
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
		log.Printf("RowScanError in backend/repositories/profile/repo_profile.go/GetFollowers: %v", err)
		return nil, fmt.Errorf("RowScanError: failed to get the followersID: %w", err)

	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Nickname, &user.AvatarPath)
		if err != nil {
			log.Printf("RowScanError in backend/repositories/profile/repo_profile.go/GetFollowing: %v", err)
			return nil, fmt.Errorf("RowScanError: failed to scan following: %w", err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		log.Printf("RowsError in backend/repositories/profile/repo_profile.go/GetFollowing: %v", err)
		return nil, fmt.Errorf("RowsError: failed to iterate following: %w", err)
	}

	return &users, nil
}

// here I will update anything in the user data
func (repo *ProfileRepository) UpdateProfileData(user *models.User) (*models.User, error) {
	return nil, nil
}
