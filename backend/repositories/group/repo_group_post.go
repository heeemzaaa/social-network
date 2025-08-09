package group

import (
	"fmt"

	"social-network/backend/models"

	"github.com/google/uuid"
)

func (grepo *GroupRepository) CreatePost(post *models.PostGroup) (*models.PostGroup, *models.ErrorJson) {
	post_created := &models.PostGroup{}
	postId := uuid.New()
	query := `
	INSERT INTO
    group_posts (postID, groupID, userID, content, imagePath)
    VALUES
    (?, ?, ?, ?, ?) RETURNING postID,
    groupID,
    userID,
    content,
    imagePath,
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
    ),
	(
        SELECT
            avatarPath
        FROM
            users
        WHERE
            users.userID = ?
    );
	`
	stmt, err := grepo.db.Prepare(query)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	errScan := stmt.QueryRow(postId, post.GroupId, post.User.Id, post.Content, post.ImagePath, post.User.Id, post.User.Id, post.User.Id).Scan(
		&post_created.Id,
		&post_created.GroupId,
		&post_created.User.Id,
		&post_created.Content,
		&post_created.ImagePath,
		&post_created.CreatedAt,
		&post_created.User.FullName,
		&post_created.User.Nickname,
		&post_created.User.ImagePath,
	)
	if errScan != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", errScan)}
	}

	return post_created, nil
}

// all the posts
// add the offset and the limit after
func (grepo *GroupRepository) GetPosts(userId, groupId string, offset string) ([]models.PostGroup, *models.ErrorJson) {
	posts := []models.PostGroup{}
	
	var where string
	if offset == "0" {
		where = "WHERE group_posts.groupID = ?"
	} else {
		where = `WHERE  group_posts.groupID = ? AND group_posts.createdAt < (
			SELECT createdAt from group_posts WHERE postID = ? 
		)`
	}

	query := fmt.Sprintf(`
	WITH
    cte_likes as (
        select
            entityID,
            count(*) as total_likes
        from
            group_reactions
        WHERE
            group_reactions.entityType = "post"
            AND group_reactions.reaction = 1
        group by
            entityID
    ),
    cte_comments as (
        SELECT
            postID,
            count(*) as total_comments
        from
            group_posts_comments
        GROUP BY
            postID
    )
	SELECT
		concat (users.firstName, " ", users.lastName),
		users.nickname,
		users.userID,
		users.avatarPath,
		group_posts.postID,
		group_posts.createdAt,
		group_posts.content,
		group_posts.imagePath,
		coalesce(cte_likes.total_likes, 0) as total_likes,
		coalesce(cte_comments.total_comments, 0) as total_comments,
		coalesce(group_reactions.userID, 0) as liked
	FROM
		group_posts
		INNER JOIN users ON group_posts.userID = users.userID
		LEFT JOIN cte_likes ON group_posts.postID = cte_likes.entityID
		LEFT JOIN cte_comments ON cte_comments.postID = group_posts.postID
		LEFT JOIN group_reactions ON group_reactions.entityID = group_posts.postID
		AND group_reactions.userID = ?
		AND group_reactions.reaction = 1
		AND group_reactions.entityType = "post"
	%v
	ORDER BY
		group_posts.createdAt DESC
	LIMIT
		20	
   `, where)

	args := []any{userId, groupId}
	if offset != "0" {
		args = append(args, offset)
	}
	stmt, err := grepo.db.Prepare(query)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v1", err)}
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v2", err)}
	}
	defer rows.Close()

	for rows.Next() {
		var post models.PostGroup
		if err := rows.Scan(&post.User.FullName,
			&post.User.Nickname,
			&post.User.Id,
			&post.User.ImagePath,
			&post.Id,
			&post.CreatedAt,
			&post.Content,
			&post.ImagePath,
			&post.TotalLikes,
			&post.TotalComments,
			&post.Liked); err != nil {
			return posts, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v3", err)}
		}
		posts = append(posts, post)
	}
	return posts, nil
}

// got everything done here
