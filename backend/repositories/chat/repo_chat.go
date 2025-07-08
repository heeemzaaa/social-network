package chat

import (
	"database/sql"
	"fmt"
	"social-network/backend/models"

	"github.com/google/uuid"
)

type ChatRepository struct {
	db *sql.DB
}

func NewChatRepository(db *sql.DB) *ChatRepository {
	return &ChatRepository{db: db}
}

// here I will get the session Id by the token given by the browser to check the auth 
func (repo *ChatRepository) GetSessionbyTokenEnsureAuth(token string) (*models.Session, *models.ErrorJson) {
	session := models.Session{}
	query := `SELECT sessions.userID, sessions.sessionToken , sessions.expiresAt, users.firstName, users.lastName 
	FROM sessions INNER JOIN users ON users.userID = sessions.userID
	WHERE sessionToken = ?`
	row := repo.db.QueryRow(query, token).Scan(&session.UserId, &session.Token, &session.ExpDate, &session.FirstName, &session.LastName)
	if row == sql.ErrNoRows {
		return nil, &models.ErrorJson{Status: 401, Message: " Unauthorized Access"}
	}
	return &session, nil
}

// add the message , I think in the case of inserting a message ,
// I think the group and user here have the same logic 
// still need to check this
func (repo *ChatRepository) AddMessage(message *models.Message) (*models.Message, *models.ErrorJson) {
	message_created := &models.Message{}
	message.ID = uuid.New().String()
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
	return message_created, nil
}

// the one logged in trying to see the messages will not be got from the query
// sender and receiver and the offset and limit als
func (repo *ChatRepository) GetMessages(sender_id, target_id string, offset int, type_ string) ([]models.Message, *models.ErrorJson) {
	var messages []models.Message
	var query string
	var args []any

	switch type_ {
	case "private":
		switch offset {
		case 0:
			query = `
		SELECT
			s.firstName || ' ' || s.lastName AS sender_name,
			r.firstName || ' ' || r.lastName AS receiver_name,
			messages.content,
			messages.created_at,
			messages.id
		FROM messages
		INNER JOIN users s ON messages.sender_id = s.userID
		INNER JOIN users r ON messages.target_id = r.userID
		WHERE messages.type = 'private' AND
		      messages.sender_id IN (?, ?) AND
		      messages.target_id IN (?, ?)
		ORDER BY messages.created_at DESC
		LIMIT 10;`
			args = append(args, sender_id, target_id, sender_id, target_id)

		default:
			query = `
		SELECT
			s.firstName || ' ' || s.lastName AS sender_name,
			r.firstName || ' ' || r.lastName AS receiver_name,
			messages.content,
			messages.created_at,
			messages.id
		FROM messages
		INNER JOIN users s ON messages.sender_id = s.userID
		INNER JOIN users r ON messages.target_id = r.userID
		WHERE messages.type = 'private' AND
		      messages.sender_id IN (?, ?) AND
		      messages.target_id IN (?, ?) AND
		      messages.id < ?
		ORDER BY messages.created_at DESC
		LIMIT 10;`
			args = append(args, sender_id, target_id, sender_id, target_id, offset)
		}

	case "group":
		switch offset {
		case 0:
			query = `
		SELECT
			s.firstName || ' ' || s.lastName AS sender_name,
			messages.content,
			messages.created_at,
			messages.id
		FROM messages
		INNER JOIN users s ON messages.sender_id = s.userID
		WHERE messages.type = 'group' AND messages.target_id = ?
		ORDER BY messages.created_at DESC
		LIMIT 10;`
			args = append(args, target_id)

		default:
			query = `
		SELECT
			s.firstName || ' ' || s.lastName AS sender_name,
			messages.content,
			messages.created_at,
			messages.id
		FROM messages
		INNER JOIN users s ON messages.sender_id = s.userID
		WHERE messages.type = 'group' AND messages.target_id = ? AND messages.id < ?
		ORDER BY messages.created_at DESC
		LIMIT 10;`
			args = append(args, target_id, offset)
		}
	}

	rows, err := repo.db.Query(query, args...)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	defer rows.Close()

	for rows.Next() {
		var message models.Message
		switch type_ {
		case "private":
			err := rows.Scan(&message.SenderName, &message.ReceiverName, &message.Content, &message.CreatedAt, &message.ID)
			if err != nil {
				return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
			}
		case "group":
			err := rows.Scan(&message.SenderName, &message.Content, &message.CreatedAt, &message.ID)
			if err != nil {
				return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
			}
		}
		messages = append(messages, message)
	}

	return messages, nil

}

// here I will get the full name of the receiver, check inside the case of private message
func (repo *ChatRepository) GetFullNameById(targetID string) (string, string, *models.ErrorJson) {
	var fisrtName, lastName string
	query := `SELECT firstName, lastName FROM users WHERE userID = ?`
	err := repo.db.QueryRow(query, targetID).Scan(&fisrtName, &lastName)
	if err != nil {
		return "", "", &models.ErrorJson{Status: 500, Error: "", Message: fmt.Sprintf("%v", err)}
	}
	return fisrtName, lastName, nil
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
