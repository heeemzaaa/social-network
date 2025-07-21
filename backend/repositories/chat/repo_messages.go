package chat

import (
	"fmt"
	"log"

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
			VALUES (?,?,?,?,?,?) RETURNING sender_id ,target_id ,content, created_at;
			`

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		log.Println("Error preparing the query to add the message: ", err)
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	if err = stmt.QueryRow(message.ID, message.SenderID, message.TargetID, message.Type, message.Content, message.CreatedAt).Scan(
		&message_created.SenderID, &message_created.TargetID,
		&message_created.Content, &message_created.CreatedAt); err != nil {
		log.Println("Error adding the message to the database: ", err)
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}

	var firstName, lastName string
	query = `SELECT firstName, lastName FROM users WHERE userID = ?`

	stmt, err = repo.db.Prepare(query)
	if err != nil {
		log.Println("Error preparing the query to fetch the user: ", err)
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	err = stmt.QueryRow(message.SenderID).Scan(&firstName, &lastName)
	if err != nil {
		log.Println("Error getting the fullname of the user: ", err)
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
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
			m.id,                                 -- [CHANGED] Added message ID to SELECT
			m.sender_id,                          -- [CHANGED] Added sender_id to SELECT
			m.target_id,                          -- [CHANGED] Added target_id to SELECT
			s.firstName || ' ' || s.lastName AS sender_name,
			r.firstName || ' ' || r.lastName AS receiver_name,
			m.content,
			m.created_at
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
			m.id,                                 -- [CHANGED] Added message ID to SELECT
			m.sender_id,                          -- [CHANGED] Added sender_id to SELECT
			m.target_id,                          -- [CHANGED] Added target_id to SELECT
			s.firstName || ' ' || s.lastName AS sender_name,
			m.content,
			m.created_at
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
		var message models.Message

		if type_ == "private" {
			err := rows.Scan(
				&message.ID,       // [CHANGED] New: scan real message ID
				&message.SenderID, // [CHANGED] New: scan real sender_id
				&message.TargetID, // [CHANGED] New: scan real target_id
				&message.SenderName,
				&message.ReceiverName,
				&message.Content,
				&message.CreatedAt,
			)
			if err != nil {
				return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
			}
		} else { // group
			err := rows.Scan(
				&message.ID,       // [CHANGED] New: scan real message ID
				&message.SenderID, // [CHANGED] New: scan real sender_id
				&message.TargetID, // [CHANGED] New: scan real target_id
				&message.SenderName,
				&message.Content,
				&message.CreatedAt,
			)
			if err != nil {
				return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
			}
		}


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

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		log.Println("Error preparing the query to edit the read status: ", err)
		return &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	_, err = repo.db.Exec(query, receiver_id, sender_id)
	if err != nil {
		log.Println("Error executting the query to edit the seen status: ", err)
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	return nil
}
