package chat

import (
	"fmt"
	"time"

	"social-network/backend/models"
	"social-network/backend/utils"
)

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
func (repo *ChatRepository) GetMessages(sender_id, target_id string, lastMessageTime time.Time, type_ string) ([]models.Message, *models.ErrorJson) {
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
		 AND (
  			(m.sender_id = ? AND m.target_id = ?)
  		OR
  			(m.sender_id = ? AND m.target_id = ?)
		)
`
		args = append(args, sender_id, target_id, target_id, sender_id)

		if !lastMessageTime.IsZero() {
			query += " AND m.created_at < ?"
			args = append(args, lastMessageTime.Format(time.RFC3339))
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

		if !lastMessageTime.IsZero() {
			query += " AND m.created_at < ?"
			args = append(args, lastMessageTime.Format(time.RFC3339))
		}

		query += " ORDER BY m.created_at DESC LIMIT 10"
	}

	rows, err := repo.db.Query(query, args...)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	defer rows.Close()

	for rows.Next() {
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
		message.Type = type_
		message.TargetID = target_id
		message.SenderID = sender_id
		messages = append(messages, message)
	}

	return messages, nil
}

func (repo *ChatRepository) EditReadStatus(sender_id, receiver_id string) *models.ErrorJson {
	query := `
	UPDATE messages
	SET
		readStatus = 1
	WHERE
		sender_iD = ?
		AND target_id = ?
		AND type = 'private'
		AND readStatus = 0
	`
	_, err := repo.db.Exec(query, receiver_id, sender_id)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	return nil
}
