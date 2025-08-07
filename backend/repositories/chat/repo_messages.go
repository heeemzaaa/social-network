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

	query := `
		INSERT INTO messages (id, sender_id, target_id, type, content)
		VALUES (?, ?, ?, ?, ?)
		RETURNING sender_id, target_id, content, created_at, type , 
		(SELECT users.firstName || ' ' || users.lastName AS sender_name  
		FROM users WHERE users.userID =  ?)
	`

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		log.Println("Error preparing the insert message query: ", err)
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	if err = stmt.QueryRow(message.ID, message.SenderID, message.TargetID, message.Type, message.Content, message.SenderID).Scan(
		&message_created.SenderID,
		&message_created.TargetID,
		&message_created.Content,
		&message_created.CreatedAt,
		&message_created.Type,
		&message_created.SenderName,
	); err != nil {
		log.Println("Error executing insert and select sender name: ", err)
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}

	return message_created, nil
}

// the one logged in trying to see the messages will not be got from the query
// sender and receiver and the offset and limit als
func (repo *ChatRepository) GetMessages(sender_id, target_id, type_ string) ([]models.Message, *models.ErrorJson) {
	messages := []models.Message{}
	var query string
	var args []any

	switch type_ {
	case "private":
		query = `
		SELECT
			m.id,                                
			m.sender_id,                         
			m.target_id,                 
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
		ORDER BY m.created_at DESC
		`
		args = append(args, sender_id, target_id, sender_id, target_id)

	case "group":
		query = `
		SELECT
			m.id,                     
			m.sender_id,                   
			m.target_id,                 
			s.firstName || ' ' || s.lastName AS sender_name,
			m.content,
			m.created_at
		FROM messages m
		INNER JOIN users s ON m.sender_id = s.userID
		WHERE m.type = 'group' AND m.target_id = ?
		ORDER BY m.created_at DESC LIMIT 10
		`
		args = append(args, target_id)
	}

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return messages, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer rows.Close()

	for rows.Next() {
		var message models.Message

		if type_ == "private" {
			err := rows.Scan(
				&message.ID,
				&message.SenderID,
				&message.TargetID,
				&message.SenderName,
				&message.ReceiverName,
				&message.Content,
				&message.CreatedAt,
			)
			if err != nil {
				return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
			}
		} else {
			err := rows.Scan(
				&message.ID,
				&message.SenderID,
				&message.TargetID,
				&message.SenderName,
				&message.Content,
				&message.CreatedAt,
			)
			if err != nil {
				return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
			}
		}
		message.Type = type_
		messages = append(messages, message)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error in the whole process of scan => in get messages: ", err)
		return []models.Message{}, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
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

	_, err = stmt.Exec(receiver_id, sender_id)
	if err != nil {
		log.Println("Error executting the query to edit the seen status: ", err)
		return &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	return nil
}
