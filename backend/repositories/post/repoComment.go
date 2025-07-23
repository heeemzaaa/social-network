package repositories

import (
	"fmt"
	"log"

	"social-network/backend/models"

	"github.com/google/uuid"
)

func (r *PostsRepository) CreateComment(userID string, postID string, content string, image_url string) (string, string, *models.ErrorJson) {
	commentID := uuid.New().String()
	query := `INSERT INTO comments (commentID, postID, userID, content, image_url) VALUES (?, ?, ?, ?, ?)`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		log.Println("Error preparing the query to get posts by id: ", err)
		return "", "", &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	_, err = stmt.Exec(commentID, postID, userID, content, image_url)
	if err != nil {
		return "", "", &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}

	userName, _ := r.GetUserName(userID)
	return commentID, userName, nil
}
