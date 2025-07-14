package group

import (
	"database/sql"
	"fmt"

	"social-network/backend/models"

	"github.com/google/uuid"
)

func (grepo *GroupRepository) CreatePost(post *models.PostGroup) (*models.PostGroup, *models.ErrorJson) {
	post_created := &models.PostGroup{}
	postId := uuid.New()
	query := `INSERT INTO posts(postID, groupID, userID, content, imagePath) 
	VALUES (?, ?, ?, ?, ?) 
	RETURNING postID , content ,createdAt`
	stmt, err := grepo.db.Prepare(query)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()
	err = stmt.QueryRow(postId, post.GroupId, post.UserId, post.Content, post.ImagePath).Scan(&post_created.Id,
		&post_created.Content, &post_created.CreatedAt)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}

	return post_created, nil
}

// all the posts
// add the offset and the limit after
func (grepo *GroupRepository) GetPosts(user_id string, offset int) ([]models.PostGroup, *models.ErrorJson) {

	var posts []models.PostGroup
    // query needs an update because the reactions table does not exist 
	// also the tables names are not correct 
	query := `
	with
    cte_likes as (
        select
            entityID,
            count(*) as total_likes
        from
            reactions
        WHERE
            reactions.entityTypeID = 1
            AND reactions.reaction = 1
        group by
            entityID
    ),
    cte_comments as (
        SELECT
            postID,
            count(*) as total_comments
        from
            comments
        GROUP BY
            postID
    )
	SELECT
		DISTINCT
		users.nickname,
		posts.postID,
		posts.createdAt,
		posts.content,
		coalesce(cte_likes.total_likes, 0) as total_likes,
		coalesce(cte_comments.total_comments, 0) as total_comments,
		coalesce(reactions.userID,0) as liked
	FROM
		posts
		INNER JOIN users ON posts.userID = users.userID
		LEFT JOIN cte_likes ON posts.postID = cte_likes.entityID
		LEFT JOIN cte_comments ON cte_comments.postID = posts.postID
		LEFT JOIN reactions ON reactions.entityID = posts.postID 
		AND reactions.userID = ? AND reactions.reaction = 1 AND reactions.entityTypeID = 1
	ORDER BY
		posts.createdAt DESC
		LIMIT 20  OFFSET  ? ;`

		
	rows, err := grepo.db.Query(query, user_id, offset)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}

	if rows.Err() == sql.ErrNoRows {
		return posts, nil
	}

	for rows.Next() {
		var post models.PostGroup
		if err := rows.Scan(&post.Username, &post.Id, &post.CreatedAt,
			&post.Content, &post.TotalLikes, &post.TotalComments, &post.Liked); err != nil {
			return posts, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
		}

		posts = append(posts, post)

	}
	defer rows.Close()
	return posts, nil
}

// got everything done here
