package post

import (
	"fmt"
	"log"

	"social-network/backend/models"
)

func (r *PostsRepository) GetAllPosts(userID string) ([]models.Post, *models.ErrorJson) {
	query := `
SELECT DISTINCT  p.postID, p.userID,   p.content,   p.createdAt,   p.privacy,   p.image_url, CONCAT(u.firstName, ' ', u.lastName) AS fullName, u.nickname, u.avatarPath,  
    COUNT(DISTINCT r1.reactionID) AS total_likes,
    CASE WHEN r2.reaction = 1 THEN 1 ELSE 0 END AS liked, 
    COUNT(DISTINCT c.commentID) AS total_comments  -- <-- count of comments

FROM posts p
JOIN users u ON p.userID = u.userID
LEFT JOIN post_access pa ON p.postID = pa.postID
LEFT JOIN followers f ON p.userID = f.userID
LEFT JOIN reactions r1 ON r1.entityID = p.postID AND r1.entityType = 'post' AND r1.reaction = 1
LEFT JOIN reactions r2 ON r2.entityID = p.postID AND r2.entityType = 'post' AND r2.userID = ?  -- current user
LEFT JOIN comments c ON c.postID = p.postID  -- <-- join with comments
WHERE
    p.privacy = 'public'
    OR p.userID = ?                    -- author's own posts
    OR (p.privacy = 'private' AND pa.userID = ?)
    OR (p.privacy = 'almost private' AND f.followerID = ?)
GROUP BY
    p.postID, p.userID, p.content, p.createdAt, p.privacy, p.image_url, fullName, r2.reaction
ORDER BY
    p.createdAt DESC;

`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		log.Println("Error preparing the query to get posts: ", err)
		return []models.Post{}, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	rows, err := stmt.Query(userID, userID, userID, userID)
	if err != nil {
		log.Println("error getting the post from database: ", err)
		return []models.Post{}, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var p models.Post
		if err := rows.Scan(
			&p.Id,
			&p.User.Id,
			&p.Content,
			&p.CreatedAt,
			&p.Privacy,
			&p.Img,
			&p.User.FullName,
			&p.User.Nickname,
			&p.User.ImagePath,
			&p.TotalLikes,
			&p.Liked,
			&p.TotalComments); err != nil {
			log.Println("Error scanning the posts: ", err)
			return []models.Post{}, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
		}
		posts = append(posts, p)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error scanning all feed's posts: ", err)
		return []models.Post{}, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}

	return posts, nil
}
