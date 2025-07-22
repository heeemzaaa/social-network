package repositories

import (
	"fmt"

	"social-network/backend/models"
)

func (repo *PostsRepository) GetComments(postID string) ([]models.Comment, *models.ErrorJson) {
	comments := []models.Comment{}

	query := `
	SELECT 
	u.userID, u.firstName, u.lastName, u.avatarPath,
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
			&comment.User.FirstName,
			&comment.User.LastName,
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
	return comments, nil
}
