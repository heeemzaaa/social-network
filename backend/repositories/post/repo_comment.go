package repositories

import (
	"fmt"
	"log"

	"social-network/backend/models"
)

func (repo *PostsRepository) GetComments(postID string) ([]models.Comment, *models.ErrorJson) {
	comments := []models.Comment{}

	query := `
	SELECT 
	u.userID, CONCAT(u.firstName, ' ', u.lastName) AS fullName, u.nickname, u.avatarPath,
	c.commentID, c.content, c.image_url, c.createdAt
	FROM comments c
	INNER JOIN users u ON c.userID = u.userID
	WHERE c.postID = ?
	ORDER BY c.createdAt ASC
	`

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return comments, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	rows, err := stmt.Query(postID)
	if err != nil {
		return comments, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer rows.Close()

	for rows.Next() {
		var comment models.Comment

		err := rows.Scan(
			&comment.User.Id,
			&comment.User.FullName,
			&comment.User.Nickname,
			&comment.User.ImagePath,
			&comment.Id,
			&comment.Content,
			&comment.Img,
			&comment.CreatedAt,
		)
		if err != nil {
			return []models.Comment{}, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
		}
		comments = append(comments, comment)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error scanning all comments of the post: ", err)
		return []models.Comment{}, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}

	return comments, nil
}
