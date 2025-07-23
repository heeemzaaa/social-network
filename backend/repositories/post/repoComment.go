package repositories

import (
	"fmt"

	"github.com/google/uuid"
)

func (r *PostsRepository) CreateComment(userID string, postID string, content string, image_url string) (string, string, error) {
	fmt.Println("entred rtepository comment")
	commentID := uuid.New().String()
	query := `INSERT INTO comments (commentID, postID, userID, content, image_url) VALUES (?, ?, ?, ?, ?)`

	_, err := r.db.Exec(query, commentID, postID, userID, content, image_url)
	if err != nil {
		return "", "", err
	}
	userNmae, _ := r.GetUserName(userID)
	return commentID, userNmae, nil
}
