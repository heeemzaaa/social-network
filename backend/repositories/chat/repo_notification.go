package chat

import (
	"fmt"

	"social-network/backend/models"
)

// delete duplicate notification before insert notification with the same state
func (repo *ChatRepository) InsertNewNotification(data models.Notification) *models.ErrorJson {
	query := `
		INSERT INTO notifications (
			notifId, recieverId, senderId, senderFullName, seen, notifType, notifStatus, groupId, groupName, eventId, createdAt
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	_, err = stmt.Exec(data.Id, data.RecieverId, data.SenderId, data.SenderFullName, data.Seen, data.Type, data.Status, data.GroupId, data.GroupName, data.EventId, data.CreatedAt)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	return nil
}

// delete duplicated follow notification
func  (repo *ChatRepository) DeleteFollowNotification(userId, authUserId, notifType string) *models.ErrorJson {
	query := `DELETE FROM notifications WHERE senderId = ? AND recieverId = ? AND (notifType = "follow-private" OR notifType = "follow-public")`

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	_, err = stmt.Exec(userId, authUserId)
	if err != nil {
		return &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err), Message: "faild to delete notification"}
	}
	return nil
}

// delete duplicated group notification
func  (repo *ChatRepository) DeleteGroupNotification(userId, authUserId, notifType, groupId string) *models.ErrorJson {
	query := `DELETE FROM notifications WHERE senderId = ? AND recieverId = ? AND notifType = ? AND groupId = ?`

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	_, err = stmt.Exec(userId, authUserId, notifType, groupId)
	if err != nil {
		return &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err), Message: "faild to delete group notification"}
	}
	return nil
}