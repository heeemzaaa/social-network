package chat

import (
	"fmt"
	"log"
	"time"

	"social-network/backend/models"
)

// get the users that I have connection with
func (repo *ChatRepository) GetUsers(authUserID string) ([]models.User, *models.ErrorJson) {
	var users []models.User
	var lastInteractionStr string

	query := `WITH 
cte_latest_interaction AS (
    SELECT
        CASE 
            WHEN sender_id = ? THEN target_id 
            ELSE sender_id 
        END AS userID,
        MAX(created_at) AS lastInteraction
    FROM messages
    WHERE (sender_id = ? OR target_id = ?)
      AND type = 'private'
    GROUP BY userID
),
cte_connections AS (
    SELECT userID FROM followers WHERE followerID = ?     
    UNION
    SELECT followerID AS userID FROM followers WHERE userID = ?  
),
cte_ordered_users AS (
    SELECT 
        i.lastInteraction,
        u.userID, 
        u.firstName,
        u.lastName,
        u.avatarPath
    FROM users u 
    JOIN cte_connections f ON u.userID = f.userID
    LEFT JOIN cte_latest_interaction i ON i.userID = u.userID
    WHERE u.userID != ?
),
cte_notifications AS (
    SELECT 
        sender_id,
        COUNT(*) AS notifications 
    FROM messages 
    WHERE readStatus = 0
      AND target_id = ?
      AND type = 'private'
    GROUP BY sender_id
)
SELECT 
    u.userID, 
    u.firstName,
    u.lastName,
    COALESCE(u.lastInteraction, '') AS lastInteraction,
    u.avatarPath,
    COALESCE(n.notifications, 0) AS notifications
FROM cte_ordered_users u
LEFT JOIN cte_notifications n ON u.userID = n.sender_id
ORDER BY 
    CASE WHEN u.lastInteraction IS NULL THEN 1 ELSE 0 END,
    u.lastInteraction DESC,
    u.firstName,
    u.lastName;

`

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		log.Println("Error preparing the query to get users: ", err)
		return users, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	rows, err := stmt.Query(authUserID, authUserID, authUserID, authUserID, authUserID, authUserID, authUserID)
	if err != nil {
		log.Println("Error getting the users: ", err)
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err := rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&lastInteractionStr,
			&user.ImagePath,
			&user.Notifications,
		); err != nil {
			log.Println("Error scanning the users: ", err)
			return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
		}
		if lastInteractionStr != "" {
			user.LastInteraction, err = time.Parse(time.RFC3339, lastInteractionStr)
			if err != nil {
				log.Println("Error parsing time in get users: ", err)
				return nil, &models.ErrorJson{Status: 400, Error: "Invalid time format !"}
			}
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error in the whole process of scan => in get users: ", err)
		return []models.User{}, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}

	return users, nil
}

func (repo *ChatRepository) GetGroups(authUserID string) ([]models.Group, *models.ErrorJson) {
	var groups []models.Group
	lastInteractionStr := ""
	query := `
	WITH cte_latest_group_messages AS (
    SELECT
        m.target_id AS groupID,
        MAX(m.created_at) AS lastInteraction
    FROM messages m
    WHERE m.type = 'group'
    GROUP BY m.target_id
),
cte_my_groups AS (
    SELECT g.groupID, g.title, g.imagePath
    FROM groups g
    INNER JOIN group_membership gm ON gm.groupID = g.groupID
    WHERE gm.userID = ?
)
SELECT 
    g.title,
    g.imagePath,
    COALESCE(lgm.lastInteraction, '') AS lastInteraction
FROM cte_my_groups g
LEFT JOIN cte_latest_group_messages lgm ON lgm.groupID = g.groupID
ORDER BY 
    CASE WHEN lgm.lastInteraction IS NULL THEN 1 ELSE 0 END,
    lgm.lastInteraction DESC,
    g.title ASC;

	`

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		log.Println("Error preparing the query to get groups: ", err)
		return groups, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	rows, err := stmt.Query(authUserID)
	if err != nil {
		log.Println("Error getting groups: ", err)
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}

	for rows.Next() {
		var group models.Group
		err := rows.Scan(&group.Title, &group.ImagePath, &lastInteractionStr)
		if err != nil {
			log.Println("Error scanning groups: ", err)
			return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
		}

		if lastInteractionStr != "" {
			group.LastInteraction, err = time.Parse(time.RFC3339, lastInteractionStr)
			if err != nil {
				log.Println("Error parsing time in get groups: ", err)
				return nil, &models.ErrorJson{Status: 400, Error: "Invalid time format !"}
			}
		}
		groups = append(groups, group)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error in the whole process of scan => in get groups: ", err)
		return []models.Group{}, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}

	return groups, nil
}

// here I will get the userIDs of all the members in a group , to broadcast the messages to them
func (repo *ChatRepository) GetMembersOfGroup(groupID string) ([]string, *models.ErrorJson) {
	var users []string

	query := `SELECT userID FROM group_membership WHERE groupID = ?`

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		log.Println("Error preparing the query to get members of a group: ", err)
		return users, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	rows, err := stmt.Query(groupID)
	if err != nil {
		log.Println("Error getting the members of a group: ", err)
		return nil, &models.ErrorJson{Status: 500, Error: "", Message: fmt.Sprintf("%v", err)}
	}

	for rows.Next() {
		var userID string
		err := rows.Scan(&userID)
		if err != nil {
			log.Println("Error scanning the members of a group: ", err)
			return nil, &models.ErrorJson{Status: 500, Error: "", Message: fmt.Sprintf("%v", err)}
		}

		users = append(users, userID)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error in the whole process of scan => in get members of groups: ", err)
		return []string{}, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}

	return users, nil
}
