package repositories

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
		return  false, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}

	return exists, nil
}
