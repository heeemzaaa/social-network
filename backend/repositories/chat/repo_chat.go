package chat

import (
	"database/sql"
	"fmt"
	"social-network/backend/models"
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
	query := `INSERT INTO messages (senderID,receiverID,message, createdAt) 
	VALUES (?,?,?,?) RETURNING senderID ,receiverID ,message, createdAt;`
	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return nil, models.NewErrorJson(500, fmt.Sprintf("%v", err))
	}
	defer stmt.Close()
	if err = stmt.QueryRow(message.SenderID, message.TargetID, message.Content, message.CreatedAt).Scan(
		&message_created.SenderID, &message_created.TargetID,
		&message_created.Content, &message_created.CreatedAt); err != nil {
		return nil, models.NewErrorJson(500, fmt.Sprintf("%v 1", err))
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
				s.fisrtName, s.lastName AS sender,
				r.firstName, r.lastName AS receiver,
				messages.content,
				messages.created_at,
				messages.id
			FROM
				messages INNER JOIN users s
				ON messages.sender_id = s.userID 
				JOIN users r ON 
				messages.receiverID = r.userID
			WHERE
				sender_id IN (?, ?)
				AND receiverID IN (?, ?)
			ORDER BY  messages.createdAt DESC
			LIMIT
				10;`
	default:
		query = `
			SELECT
				s.nickname AS sender,
				r.nickname AS receiver,
				messages.message,
				messages.createdAt,
				messages.messageID
			FROM
				messages INNER JOIN users s
				ON messages.senderID = s.userID 
				JOIN users r ON 
				messages.receiverID = r.userID
			WHERE
				senderID IN (?, ?)
				AND receiverID IN (?, ?)
				AND messages.messageID < ?
			ORDER BY  messages.createdAt DESC
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

func (repo *ChatRepository) EditReadStatus(sender_id, target_id string) *models.ErrorJson {
	query := `
	UPDATE messages
	SET
		readStatus = 1
	WHERE
		senderID = ?
		AND receiverID = ?
	`
	_, err := repo.db.Exec(query, target_id, sender_id)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	return nil
}
