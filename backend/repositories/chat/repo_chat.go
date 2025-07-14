package chat

import (
	"database/sql"
	"fmt"

	"social-network/backend/models"
	"social-network/backend/utils"
)

type ChatRepository struct {
	db *sql.DB
}

func NewChatRepository(db *sql.DB) *ChatRepository {
	return &ChatRepository{db: db}
}

func (repo *ChatRepository) GetUsers(offset, user_id int) ([]models.User, *models.ErrorJson) {
	var users []models.User
	query := `
		WITH 
		cte_latest_interaction AS (
    		SELECT
        		CASE 
            		WHEN sender_id = ? THEN target_id 
            		ELSE sender_id 
        		END AS userID,
        	MAX(created_at) AS lastInteraction,
        	message
    	FROM messages
    		WHERE (sender_id = ? OR target_id = ?)
      		AND type = 'private'
    		GROUP BY userID
		),
		cte_ordered_users AS (
    		SELECT 
       			 i.message, 
       			 COALESCE(i.lastInteraction, 0) AS lastInteraction,  
        		u.userID, 
        		u.firstName
    		FROM users u 
    		LEFT JOIN cte_latest_interaction i 
       		ON i.userID = u.userID
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
    	COALESCE(u.message, '') AS message, 
   			u.lastInteraction, 
    	COALESCE(n.notifications, 0) AS notifications
		FROM cte_ordered_users u
			LEFT JOIN cte_notifications n ON u.userID = n.sender_id
			ORDER BY u.lastInteraction DESC, u.firstName, u.lastName;
`

	rows, err := repo.db.Query(query, user_id, user_id, user_id, user_id, user_id)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Nickname, &user.LastMessage, &user.LastInteraction); err != nil {
			return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v ", err)}
		}
		users = append(users, user)
	}
	defer rows.Close()

	return users, nil
}

// here I will get the session Id by the token given by the browser to check the auth
func (repo *ChatRepository) GetSessionbyTokenEnsureAuth(token string) (*models.Session, *models.ErrorJson) {
	session := models.Session{}
	query := `SELECT sessions.userID, sessions.sessionToken , users.firstName, users.lastName 
	FROM sessions INNER JOIN users ON users.userID = sessions.userID
	WHERE sessionToken = ?`
	row := repo.db.QueryRow(query, token).Scan(&session.UserId, &session.Token, &session.FirstName, &session.LastName)
	if row == sql.ErrNoRows {
		return nil, &models.ErrorJson{Status: 401, Message: " Unauthorized Access"}
	}
	return &session, nil
}

// get the userID from the session , we will need it also
func (repo *ChatRepository) GetUserIdFromSession(sessionID string) (string, *models.ErrorJson) {
	var userID string
	query := `SELECT userID FROM sessions WHERE sessionToken = ?`
	errQuery := repo.db.QueryRow(query, sessionID).Scan(&userID)
	if errQuery != nil {
		return "", &models.ErrorJson{Status: 500, Error: "", Message: fmt.Sprintf("%v", errQuery)}
	}

	return userID, nil
}

// add the message , I think in the case of inserting a message ,
// I think the group and user here have the same logic
// still need to check this
func (repo *ChatRepository) AddMessage(message *models.Message) (*models.Message, *models.ErrorJson) {
	message_created := &models.Message{}
	message.ID = utils.NewUUID()
	query := `INSERT INTO messages (id, sender_id,target_id, type, content, created_at) 
	VALUES (?,?,?,?,?,?) RETURNING sender_id ,target_id ,content, created_at;`
	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return nil, models.NewErrorJson(500, "", fmt.Sprintf("%v", err))
	}
	defer stmt.Close()
	if err = stmt.QueryRow(message.ID, message.SenderID, message.TargetID, message.Type, message.Content, message.CreatedAt).Scan(
		&message_created.SenderID, &message_created.TargetID,
		&message_created.Content, &message_created.CreatedAt); err != nil {
		return nil, models.NewErrorJson(500, "", fmt.Sprintf("%v 1", err))
	}
	var firstName, lastName string
	query = `SELECT firstName, lastName FROM users WHERE userID = ?`
	err = repo.db.QueryRow(query, message.SenderID).Scan(&firstName, &lastName)
	if err != nil {
		return nil, models.NewErrorJson(500, "", fmt.Sprintf("%v", err))
	}
	message_created.SenderName = firstName + " " + lastName
	return message_created, nil
}

// the one logged in trying to see the messages will not be got from the query
// sender and receiver and the offset and limit als
func (repo *ChatRepository) GetMessages(sender_id, target_id, lastMessageTime, type_ string) ([]models.Message, *models.ErrorJson) {
	var messages []models.Message
	var query string
	var args []any

	switch type_ {
	case "private":
		query = `
		SELECT
			s.firstName || ' ' || s.lastName AS sender_name,
			r.firstName || ' ' || r.lastName AS receiver_name,
			m.content,
			m.created_at,
			m.id
		FROM messages m
		INNER JOIN users s ON m.sender_id = s.userID
		INNER JOIN users r ON m.target_id = r.userID
		WHERE m.type = 'private'
		  AND m.sender_id IN (?, ?)
		  AND m.target_id IN (?, ?)
		`
		args = append(args, sender_id, target_id, sender_id, target_id)

		if lastMessageTime != "" {
			query += " AND m.created_at < ?"
			args = append(args, lastMessageTime)
		}

		query += " ORDER BY m.created_at DESC LIMIT 10"

	case "group":
		query = `
		SELECT
			s.firstName || ' ' || s.lastName AS sender_name,
			m.content,
			m.created_at,
			m.id
		FROM messages m
		INNER JOIN users s ON m.sender_id = s.userID
		WHERE m.type = 'group' AND m.target_id = ?
		`
		args = append(args, target_id)

		if lastMessageTime != "" {
			query += " AND m.created_at < ?"
			args = append(args, lastMessageTime)
		}

		query += " ORDER BY m.created_at DESC LIMIT 10"
	}

	rows, err := repo.db.Query(query, args...)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	defer rows.Close()

	for rows.Next() {
		// comment
		var message models.Message
		if type_ == "private" {
			err := rows.Scan(&message.SenderName, &message.ReceiverName, &message.Content, &message.CreatedAt, &message.ID)
			if err != nil {
				return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
			}
		} else {
			err := rows.Scan(&message.SenderName, &message.Content, &message.CreatedAt, &message.ID)
			if err != nil {
				return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
			}
		}
		messages = append(messages, message)
	}

	return messages, nil
}

// check if the id givven is a member inside the group givven
func (repo *ChatRepository) IsMember(userID, groupID string) (bool, *models.ErrorJson) {
	var exists bool
	query := `SELECT EXISTS (SELECT 1 FROM group_membership WHERE userID = ? AND groupID = ? LIMIT 1)`
	err := repo.db.QueryRow(query, userID, groupID).Scan(&exists)
	if err != nil {
		return false, &models.ErrorJson{Status: 500, Error: "", Message: fmt.Sprintf("%v", err)}
	}
	return exists, nil
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

// check if the user exist, I will need this a lot hhh
func (repo *ChatRepository) UserExists(userID string) (bool, *models.ErrorJson) {
	var exists bool
	fmt.Println(userID)
	query := `SELECT EXISTS (SELECT 1 FROM users WHERE userID = ? LIMIT 1)`
	err := repo.db.QueryRow(query, userID).Scan(&exists)
	if err != nil {
		return false, &models.ErrorJson{Status: 500, Error: "", Message: fmt.Sprintf("%v", err)}
	}
	return exists, nil
}

// check if the group Id givven exists
func (repo *ChatRepository) GroupExists(groupID string) (bool, *models.ErrorJson) {
	var exists bool

	query := `SELECT EXISTS (SELECT 1 FROM groups WHERE groupID = ? LIMIT 1)`
	err := repo.db.QueryRow(query, groupID).Scan(&exists)
	if err != nil {
		return false, &models.ErrorJson{Status: 500, Error: "", Message: fmt.Sprintf("%v", err)}
	}
	return exists, nil
}
