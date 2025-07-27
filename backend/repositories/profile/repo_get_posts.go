package profile

import (
	"database/sql"
	"fmt"
	"log"

	"social-network/backend/models"
)

// here I will get the posts of the user with conditions
func (repo *ProfileRepository) GetPosts(profileID string, userID string, lastPostTime string, myProfile bool) ([]models.Post, *models.ErrorJson) {
	var query string
	posts := []models.Post{}
	var args []any

	switch myProfile {
	case true:
		query = `
		SELECT
    		p.postID,
    		u.userID, CONCAT(u.firstName, ' ', u.lastName) AS fullName, u.nickname, u.avatarPath,
    		p.content AS postContent, p.image_url AS postMedia, p.createdAt AS postCreatedAt, p.privacy,
    		(SELECT COUNT(*) FROM reactions r1 WHERE r1.entityType = 'post' AND r1.entityID = p.postID) AS post_total_likes,
    		(SELECT COUNT(*) FROM comments c1 WHERE c1.postID = p.postID) AS post_total_comments
		FROM posts p
		JOIN users u ON u.userID = p.userID
		WHERE p.userID = ?
		`
		args = append(args, userID)

		if lastPostTime != "" {
			query += " AND p.createdAt < ?"
			args = append(args, lastPostTime)
		}

		query += ` ORDER BY p.createdAt DESC LIMIT 10`
	default:
		query = `
		WITH 
			follower_check AS (
    			SELECT EXISTS (
        			SELECT 1 FROM followers WHERE userID = ? AND followerID = ?
    			) AS is_follower
			),
			user_visibility AS (
    			SELECT visibility FROM users WHERE userID = ?
			)
		SELECT 
    		p.postID,
    		u.userID, CONCAT(u.firstName, ' ', u.lastName) AS fullName, u.nickname, u.avatarPath,
    		p.content AS postContent, p.image_url AS postMedia, p.createdAt AS postCreatedAt, p.privacy,
    		(SELECT COUNT(*) FROM reactions r1 WHERE r1.entityType = 'post' AND r1.entityID = p.postID) AS post_total_likes,
   	 		(SELECT COUNT(*) FROM comments c1 WHERE c1.postID = p.postID) AS post_total_comments
		FROM posts p
		JOIN users u ON u.userID = p.userID
		JOIN user_visibility vis
		JOIN follower_check fc
		WHERE 
		    p.userID = ?
		    AND (
	        (vis.visibility = 'public' AND (
    	        p.privacy = 'public'
        	    OR (p.privacy = 'almost private' AND fc.is_follower)
            	OR (p.privacy = 'private' AND EXISTS (
                	SELECT 1 FROM post_access pa WHERE pa.postID = p.postID AND pa.userID = ?
            	))
        	))
        OR (vis.visibility = 'private' AND fc.is_follower = 1 AND (
            p.privacy = 'public'
            OR p.privacy = 'almost private'
            OR (p.privacy = 'private' AND EXISTS (
                SELECT 1 FROM post_access pa WHERE pa.postID = p.postID AND pa.userID = ?
            	))
        	))
    	)
		`
		args = append(args, profileID, userID, profileID, profileID, userID, userID)

		if lastPostTime != "" {
			query += " AND p.createdAt < ?"
			args = append(args, lastPostTime)
		}

		query += ` ORDER BY p.createdAt DESC LIMIT 10`
	}

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		log.Println("Error preparing the query to get posts of the profile: ", err)
		return posts, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return posts, nil
		}
		log.Println("Error getting the posts of the profile: ", err)
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}

	for rows.Next() {
		var post models.Post
		var user models.User

		err := rows.Scan(
			&post.Id,
			&user.Id, &user.FullName, &user.Nickname, &user.ImagePath,
			&post.Content, &post.Img, &post.CreatedAt, &post.Privacy,
			&post.TotalLikes, &post.TotalComments,
		)
		if err != nil {
			log.Println("Error scanning the post: ", err)
			return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
		}
		post.User = user
		posts = append(posts, post)

	}

	if err := rows.Err(); err != nil {
		log.Println("Error in the whole process of scan => in get posts: ", err)
		return []models.Post{}, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}

	return posts, nil
}
