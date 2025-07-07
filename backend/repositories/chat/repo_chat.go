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
	if err = stmt.QueryRow(message.ID, message.SenderID, message.TargetID,message.Type, message.Content, message.CreatedAt).Scan(
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
	switch offset {
	case 0:
		query = `
			SELECT
				s.firstName, s.lastName AS sender,
				r.firstName, r.lastName AS receiver,
				messages.content,
				messages.created_at,
				messages.id
			FROM
				messages INNER JOIN users s
				ON messages.sender_id = s.userID 
				JOIN users r ON 
				messages.target_id = r.userID
			WHERE
				sender_id IN (?, ?)
				AND target_id IN (?, ?)
			ORDER BY  messages.createdAt DESC
			LIMIT
				10;`
	default:
		query = `
			SELECT
				s.firstName, s.lastName AS sender,
				r.firstName, r.lastName AS receiver,
				messages.id
				messages.content,
				messages.created_at,
			FROM
				messages INNER JOIN users s
				ON messages.sender_id = s.userID 
				JOIN users r ON 
				messages.target_id = r.userID
			WHERE
				sender_id IN (?, ?)
				AND target_id IN (?, ?)
				AND messages.id < ?
			ORDER BY  messages.created_at DESC
			LIMIT
				10;`

	}

	rows, err := repo.db.Query(query, sender_id, target_id, sender_id, target_id, offset)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}

	for rows.Next() {
		var message models.Message
		if err := rows.Scan(&message.Content,
			&message.CreatedAt, &message.ID); err != nil {
			return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
		}
		messages = append(messages, message)
	}

	return messages, nil
}

// here I will get the full name of the receiver
func (repo *ChatRepository) GetFullNameById(targetID string) (string, string, *models.ErrorJson) {
	var fisrtName, lastName string
	
	query := `SELECT firstName, lastName FROM users WHERE userID = ?`
	err := repo.db.QueryRow(query, targetID).Scan(&fisrtName, &lastName)
	if err != nil {
		return "", "", &models.ErrorJson{Status: 500, Error: "" , Message: fmt.Sprintf("%v", err)}
	}
	return fisrtName,lastName, nil
}

func (repo *ChatRepository) GetMembersOfGroup(groupID string) ([]string, *models.ErrorJson) {
	return nil,nil
}
