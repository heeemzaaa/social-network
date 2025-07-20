package chat

import (
	"fmt"
	"time"

	"social-network/backend/models"
)

func (repo *ChatRepository) GetUsers(authUserID string) (*[]models.User, *models.ErrorJson) {
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
    	u.lastInteraction, 
		u.avatarPath,
    COALESCE(n.notifications, 0) AS notifications
	FROM cte_ordered_users u
	LEFT JOIN cte_notifications n ON u.userID = n.sender_id
	ORDER BY u.lastInteraction DESC, u.firstName, u.lastName;
`

	rows, err := repo.db.Query(query, authUserID, authUserID, authUserID, authUserID, authUserID, authUserID, authUserID)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
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
			return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("scan error: %v", err)}
		}
		if lastInteractionStr != "" {
			user.LastInteraction, err = time.Parse(time.RFC3339, lastInteractionStr)
			if err != nil {
				return nil, &models.ErrorJson{Status: 400, Error: "Invalid time format !"}
			}
		}
		users = append(users, user)
	}

	return &users, nil
}

func (repo *ChatRepository) GetGroups(authUserID string) (*[]models.Group, *models.ErrorJson) {
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
    	COALESCE(lgm.lastInteraction, CURRENT_TIMESTAMP) AS lastInteraction
	FROM cte_my_groups g
		LEFT JOIN cte_latest_group_messages lgm ON lgm.groupID = g.groupID
		ORDER BY lgm.lastInteraction DESC NULLS LAST, g.title ASC;
	`

	rows, err := repo.db.Query(query, authUserID)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}

	for rows.Next() {
		var group models.Group
		err := rows.Scan(&group.Title, &group.ImagePath, &lastInteractionStr)
		if err != nil {
			return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
		}

		if lastInteractionStr != "" {
			group.LastInteraction, err = time.Parse(time.RFC3339, lastInteractionStr)
			if err != nil {
				return nil, &models.ErrorJson{Status: 400, Error: "Invalid time format !"}
			}
		}
		groups = append(groups, group)
	}

	return &groups, nil
}

// here I will get the userIDs of all the members in a group , to broadcast the messages to them
func (repo *ChatRepository) GetMembersOfGroup(groupID string) ([]string, *models.ErrorJson) {
	var users []string

	query := `SELECT userID FROM group_membership WHERE groupID = ?`

	rows, err := repo.db.Query(query, groupID)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Error: "", Message: fmt.Sprintf("%v", err)}
	}

	for rows.Next() {
		var userID string
		err := rows.Scan(&userID)
		if err != nil {
			return nil, &models.ErrorJson{Status: 500, Error: "", Message: fmt.Sprintf("%v", err)}
		}

		users = append(users, userID)
	}

	return users, nil
}
