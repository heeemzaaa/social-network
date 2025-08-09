package post

import (
	"fmt"
	"log"

	"social-network/backend/models"
)

func (r *PostsRepository) PostExist(postID string) (bool, *models.ErrorJson) {
	var exists bool

	query := `SELECT EXISTS (SELECT 1 FROM posts WHERE postID = ? LIMIT 1)`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		log.Println("Error preparing the query to check if the post exist: ", err)
		return false, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	err = stmt.QueryRow(postID).Scan(&exists)
	if err != nil {
		log.Println("Error checking if the post exist: ", err)
		return false, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}

	return exists, nil
}

func (r *PostsRepository) PostPrivacy(postID string) (string, *models.ErrorJson) {
	var privacy string
	query := `SELECT privacy FROM posts WHERE postID = ?`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		log.Println("Error preparing the query to check the post privacy: ", err)
		return "", &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	err = stmt.QueryRow(postID).Scan(&privacy)
	if err != nil {
		log.Println("Error checking the post privacy: ", err)
		return "", &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	return privacy, nil
}

func (r *PostsRepository) PrivatePostAccess(userID string, postID string) (bool, *models.ErrorJson) {
	access := false
	query := `SELECT EXISTS (SELECT 1 FROM post_access WHERE postID = ? AND userID = ?)`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		log.Println("Error preparing the query to check the private post access: ", err)
		return false, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	err = stmt.QueryRow(postID, userID).Scan(&access)
	if err != nil {
		log.Println("Error checking if the user has the access to the post: ", err)
		return false, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	} 
	
	return access, nil
}
