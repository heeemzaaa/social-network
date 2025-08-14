package chat

import (
	"fmt"

	"social-network/backend/models"
	"social-network/backend/utils"
)

// delete duplicate notification before insert notification with the same state
func (repo *ChatRepository) InsertNewNotification(data models.Notification) *models.ErrorJson {
	notificationId := utils.NewUUID()
	query := `
		INSERT INTO notifications (
			notifId, senderId, targetId, notifType, notifStatus, content
		) VALUES (?, ?, ?, ?, ?, ?)
	`

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	_, err = stmt.Exec(notificationId, data.SenderID,
		data.TargetID, data.Type, data.Status, data.Content)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	return nil
}

// delete duplicated follow notification
func (repo *ChatRepository) DeleteFollowNotification(userId, authUserId, notifType string) *models.ErrorJson {
	query := `DELETE FROM notifications WHERE senderId = ? AND targetId = ? AND (notifType = "follow-private" OR notifType = "follow-public")`

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	_, err = stmt.Exec(userId, authUserId)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err), Error: "500 - faild to delete follow notification"}
	}
	return nil
}

// delete duplicated group notification
func (repo *ChatRepository) DeleteGroupNotification(userId, authUserId, notifType, groupId string) *models.ErrorJson {
	query := `DELETE FROM notifications WHERE senderId = ? AND targetId = ? AND notifType = ? AND groupId = ?`

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	_, err = stmt.Exec(userId, authUserId, notifType, groupId)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err), Error: "500 - faild to delete group notification"}
	}
	return nil
}
