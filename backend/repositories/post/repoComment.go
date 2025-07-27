package repositories

import (
	"fmt"
	"log"

	"social-network/backend/models"
	"social-network/backend/utils"
)

func (r *PostsRepository) CreateComment(userID string, postID string, content string, image_url string) (*models.Comment, *models.ErrorJson) {
	commentID := utils.NewUUID()
	var comment models.Comment

	query := `INSERT INTO comments c (commentID, postID, userID, content, image_url) VALUES (?, ?, ?, ?, ?)
				RETURNING c.commentID, c.postID, c.userID, c.content, c.image_url, CONCAT(u.firstName, ' ', u.lastName), u.nickname, u.avatarPath
				INNER JOIN users u ON u.userID = c.userID
	`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		log.Println("Error preparing the query to get posts by id: ", err)
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	err = stmt.QueryRow(commentID, postID, userID, content, image_url).Scan(
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
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}

	return &comment, nil
}
