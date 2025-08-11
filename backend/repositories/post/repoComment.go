package post

import (
	"fmt"
	"log"

	"social-network/backend/models"
	"social-network/backend/utils"
)

func (r *PostsRepository) CreateComment(userID string, postID string, content string, image_url string) (*models.Comment, *models.ErrorJson) {
	commentID := utils.NewUUID()
	var comment models.Comment

	query := `
    			INSERT INTO comments (commentID, postID, userID, content, image_url)
    			VALUES (?, ?, ?, ?, ?)
    			RETURNING commentID, postID, userID, content, image_url,
				(
        			SELECT
            			concat (firstName, ' ', lastName)
        			FROM
            			users
        			WHERE
            			users.userID = ?
    			) AS fullName,
    			(
        			SELECT
            			nickname
        			FROM
            			users
        			WHERE
            			users.userID = ?
    			),(
        			SELECT
            			avatarPath
        			FROM
            			users
        			WHERE
            			users.userID = ?
    			)
	`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		log.Println("Error preparing the query to get posts by id: ", err)
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	err = stmt.QueryRow(commentID, postID, userID, content, image_url, userID, userID, userID).Scan(
		&comment.Id,
		&comment.PostId,
		&comment.User.Id,
		&comment.Content,
		&comment.Img,
		&comment.User.FullName,
		&comment.User.Nickname,
		&comment.User.ImagePath,
	)
	if err != nil {
		log.Println("Error getting the data of the comment: ", err)
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	return &comment, nil
}
