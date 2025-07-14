package group

import (
	"database/sql"
	"fmt"

	"social-network/backend/models"
)

func (gRepo *GroupRepository) CreateComment(comment *models.CommentGroup) (*models.CommentGroup, *models.ErrorJson) {
	comment_created := &models.CommentGroup{}
	query := `INSERT INTO comments(postID, userID, content)  VALUES(?, ?, ?) 
	RETURNING commentID, content, createdAt;`
	stmt, err := gRepo.db.Prepare(query)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()
	if err := stmt.QueryRow(comment.PostId, comment.UserId, comment.Content).Scan(
		&comment_created.Id, &comment_created.Content,
		&comment_created.CreatedAt); err != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	return comment_created, nil
}

// But hna comments dyal wa7d l post specific
func (gRepo *GroupRepository) GetComments(user_id, postId, offset int) ([]models.CommentGroup, *models.ErrorJson) {
	var where string
	if offset == 0 {
		where = `comments.postID = ?`
	} else {
		where = `comments.postID = ? AND comments.commentID < ?`
	}
	var comments []models.CommentGroup
	query := fmt.Sprintf(`
	with
    cte_likes as (
        SELECT
            entityID,
            count(*) as total_likes
        FROM
            reactions
        WHERE
            reactions.entityTypeID = 2
			AND reactions.reaction = 1
        GROUP BY
            entityID
    )
	SELECT
		users.nickname,
		comments.commentID,
		content,
		comments.createdAt,
		coalesce(cte_likes.total_likes, 0) as total_likes,
        coalesce(reactions.userID,0) as liked
	FROM
		comments
		INNER JOIN users ON comments.userID = users.userID
		LEFT JOIN cte_likes ON cte_likes.entityID = comments.commentID
        LEFT JOIN reactions ON comments.commentID = reactions.entityID 
        AND reactions.userID = ?  AND reactions.reaction =1 AND reactions.entityTypeID = 2
	WHERE 
		%v
	ORDER BY
		comments.createdAt DESC
	LIMIT
		10;
	`, where)

	rows, err := gRepo.db.Query(query, user_id, postId, offset)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500 , Error: fmt.Sprintf("%v", err)}
	}
	if rows.Err() == sql.ErrNoRows {
		return comments, nil
	}

	for rows.Next() {
		var comment models.CommentGroup
		if err = rows.Scan(&comment.Username, &comment.Id, &comment.Content,
			&comment.CreatedAt, &comment.TotalLikes, &comment.Liked); err != nil {
			return comments,  &models.ErrorJson{Status: 500 , Error: fmt.Sprintf("%v", err)}
		}
		comments = append(comments, comment)
	}
	defer rows.Close()
	return comments, nil
}
