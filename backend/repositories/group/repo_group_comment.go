package group

import (
	"database/sql"
	"fmt"

	"social-network/backend/models"
	"social-network/backend/utils"
)

func (gRepo *GroupRepository) CreateComment(comment *models.CommentGroup) (*models.CommentGroup, *models.ErrorJson) {
	commentId := utils.NewUUID()
	comment_created := &models.CommentGroup{}
	query := `
	INSERT INTO
    group_posts_comments (
        commentID,
        postID,
        groupID,
        userID,
        content,
        imageContent
    )
VALUES
    (?, ?, ?, ?, ?, ?) RETURNING commentID,
    postID,
    groupID,
    userID,
    content,
    imageContent,
    createdAt,
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
    );`

	stmt, err := gRepo.db.Prepare(query)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v 1", err)}
	}
	defer stmt.Close()
	if err := stmt.QueryRow(commentId,
		comment.PostId,
		comment.GroupId,
		comment.User.Id,
		comment.Content,
		comment.ImagePath,
		comment.User.Id,
		comment.User.Id,
		comment.User.Id,
	).Scan(
		&comment_created.Id,
		&comment_created.PostId,
		&comment_created.GroupId,
		&comment_created.User.Id,
		&comment_created.Content,
		&comment_created.ImagePath,
		&comment_created.CreatedAt,
		&comment_created.User.FullName,
		&comment_created.User.Nickname,
		&comment_created.User.ImagePath); err != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v 2", err)}
	}
	return comment_created, nil
}

// But hna comments dyal wa7d l post specific
func (gRepo *GroupRepository) GetComments(userId, postId, groupId string, offset int) ([]models.CommentGroup, *models.ErrorJson) {
	comments := []models.CommentGroup{}
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
    )
	SELECT
		concat (users.firstName, " ", users.lastName),
		users.nickname,
		users.userID,
		group_posts_comments.commentID,
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
	WHERE  group_posts_comments.groupID = ? AND group_posts_comments.groupID
	ORDER BY
		group_posts_comments.createdAt DESC
	LIMIT
		20
	OFFSET
		?;
	`

	// need to prepare the query to see the problem fiiin
	stmt, err := gRepo.db.Prepare(query)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	rows, err := stmt.Query(userId, postId, offset)
	if err != nil {
		if err == sql.ErrNoRows {
			return comments, nil
		}
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
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
