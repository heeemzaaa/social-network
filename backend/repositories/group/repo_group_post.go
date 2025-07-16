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
	query := `INSERT INTO group_posts (postID, groupID, userID, content, imagePath) 
	VALUES (?, ?, ?, ?, ?) 
	RETURNING postID, groupID, userID, content, imagePath, createdAt`
	stmt, err := grepo.db.Prepare(query)
	if err != nil {
		fmt.Println("WHY? 1", err.Error())
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()
	errScan := stmt.QueryRow(postId, post.GroupId, post.UserId, post.Content, post.ImagePath).Scan(
		&post_created.Id,
		&post_created.GroupId,
		&post_created.UserId,
		&post_created.Content,
		&post_created.ImagePath,
		&post_created.CreatedAt)
	if errScan != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", errScan)}
	}

	return post_created, nil
}

// all the posts
// add the offset and the limit after
func (grepo *GroupRepository) GetPosts(user_id string, offset int) ([]models.PostGroup, *models.ErrorJson) {
	fmt.Println("hunaaaaaaa")
	var posts []models.PostGroup
	// query needs an update because the reactions table does not exist
	// also the tables names are not correct
	query := `
	WITH
    cte_comments as (
        SELECT
            postID,
            count(*) as total_comments
        from
            group_posts_comments
        GROUP BY
            postID
    )
	SELECT DISTINCT
		concat (users.firstName, " ", users.lastName),
		group_posts.postID,
		group_posts.createdAt,
		group_posts.content,
		coalesce(cte_comments.total_comments, 0) as total_comments
	FROM
		group_posts
		INNER JOIN users ON group_posts.userID = users.userID
		LEFT JOIN cte_comments ON cte_comments.postID = group_posts.postID
	ORDER BY
		group_posts.createdAt DESC
	LIMIT
		20
	OFFSET
		?;

   `

	stmt, err := grepo.db.Prepare(query)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	rows, err := stmt.Query(offset)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}

	if rows.Err() == sql.ErrNoRows {
		return posts, nil
	}

	for rows.Next() {
		var post models.PostGroup
		if err := rows.Scan(&post.FullName, &post.Id, &post.CreatedAt,
			&post.Content, &post.TotalComments); err != nil {
			return posts, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
		}

		posts = append(posts, post)

	}
	defer rows.Close()
	return posts, nil
}

// got everything done here
