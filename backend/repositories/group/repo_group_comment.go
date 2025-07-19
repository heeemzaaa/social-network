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
	if err := stmt.QueryRow(comment.PostId, comment.User.Id, comment.Content).Scan(
		&comment_created.Id, &comment_created.Content,
		&comment_created.CreatedAt); err != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	return comment_created, nil
}

// But hna comments dyal wa7d l post specific
func (gRepo *GroupRepository) GetComments(userId, postId, groupId  string, offset int) ([]models.CommentGroup, *models.ErrorJson) {
	
	var comments []models.CommentGroup
	query := `
	WITH
    cte_likes as (
        select
            entityID,
            count(*) as total_likes
        from
            group_reactions
        WHERE
            group_reactions.entityType = "comment"
            AND group_reactions.reaction = 1
        group by
            entityID
    ),
	SELECT
		concat (users.firstName, " ", users.lastName),
		users.nickname,
		users.userID,
		group_posts_comments.postID,
		group_posts_comments.createdAt,
		group_posts_comments.content,
		coalesce(cte_likes.total_likes, 0) as total_likes,
		coalesce(group_reactions.userID, 0) as liked
	FROM
		group_posts_comments
		INNER JOIN users ON group_posts_comments.userID = users.userID
		LEFT JOIN cte_likes ON group_posts_comments.commentID = cte_likes.entityID
		LEFT JOIN group_reactions ON group_reactions.entityID = group_posts_comments.commentID
		AND group_reactions.userID = ?
		AND group_reactions.reaction = 1
		AND group_reactions.entityType = "comment"
	WHERE group_posts_comments.groupID = ?
	ORDER BY
		group_posts_comments.createdAt DESC
	LIMIT
		20
	OFFSET
		?;
	`

	rows, err := gRepo.db.Query(query, userId, postId, offset)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	if rows.Err() == sql.ErrNoRows {
		return comments, nil
	}

	for rows.Next() {
		var comment models.CommentGroup
		if err = rows.Scan(&comment.User.FullName, &comment.User.Id, &comment.User.Nickname, &comment.Id, &comment.Content,
			&comment.CreatedAt, &comment.TotalLikes, &comment.Liked); err != nil {
			return comments, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
		}
		comments = append(comments, comment)
	}
	defer rows.Close()
	return comments, nil
}
