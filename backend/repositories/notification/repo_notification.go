package notification

import (
	"database/sql"
	"fmt"

	"social-network/backend/models"
)

type NotifRepository struct {
	db *sql.DB
}

func NewNotifRepository(db *sql.DB) *NotifRepository {
	return &NotifRepository{db: db}
}

func (repo *NotifRepository) SelectNotificationById(notifId string) (models.Notification, *models.ErrorJson) {
	query := `SELECT * FROM notifications WHERE notifId = ?`

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return models.Notification{}, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	var notification models.Notification
	if err = stmt.QueryRow(notifId).Scan(
		&notification.Id,
		&notification.RecieverId,
		&notification.SenderId,
		&notification.SenderFullName,
		&notification.Seen,
		&notification.Type,
		&notification.Status,
		&notification.GroupId,
		&notification.GroupName,
		&notification.EventId,
		&notification.CreatedAt,
	); err != nil {
		return notification, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	return notification, nil
}

// select all notifications for reciever by sort first status [later-first], second sort by time
func (repo *NotifRepository) SelectAllNotification(userId string) ([]models.Notification, *models.ErrorJson) {
	all := []models.Notification{}
	query := `SELECT * FROM notifications WHERE recieverId = ?`

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	rows, err := stmt.Query(userId)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("cannot get all norification, err: %v", err)}
	}
	defer rows.Close()

	for rows.Next() {
		var notification models.Notification
		err = rows.Scan(&notification.Id, &notification.RecieverId, &notification.SenderId, &notification.SenderFullName, &notification.Seen, &notification.Type, &notification.Status, &notification.GroupId, &notification.GroupName, &notification.EventId, &notification.CreatedAt)
		if err != nil {
			return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
		}
		all = append(all, notification)
	}

	return all, nil
}

// select all notifications by recieverId and notifType
func (repo *NotifRepository) SelectAllNotificationByType(userId, notifType string) ([]models.Notification, *models.ErrorJson) {
	all := []models.Notification{}
	query := `SELECT * FROM notifications WHERE recieverId = ? AND notifType = ?`

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	rows, err := stmt.Query(userId, notifType)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("cannot get all norification by type, err: %v", err)}
	}
	defer rows.Close()

	for rows.Next() {
		var notification models.Notification
		if err = rows.Scan(
			&notification.Id,
			&notification.RecieverId,
			&notification.SenderId,
			&notification.SenderFullName,
			&notification.Seen,
			&notification.Type,
			&notification.Status,
			&notification.GroupId,
			&notification.GroupName,
			&notification.EventId,
			&notification.CreatedAt,
		); err != nil {
			return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
		}
		all = append(all, notification)
	}
	return all, nil
}

// delete duplicate notification before insert notification with the same state
func (repo *NotifRepository) InsertNewNotification(data models.Notification) *models.ErrorJson {
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

// update notification status
func (repo *NotifRepository) UpdateStatus(notifId, status string) *models.ErrorJson {
	query := `UPDATE notifications SET notifStatus = ? WHERE notifId = ?`

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	_, err = stmt.Exec(status, notifId)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	return nil
}

// update notification seen
func (repo *NotifRepository) UpdateSeen(notifId string) *models.ErrorJson {
	query := `UPDATE notifications SET seen = 1 WHERE notifId = ?`

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	_, err = stmt.Exec(notifId)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	return nil
}

// function check if user has notification containe false seen
func (repo *NotifRepository) IsHasSeenFalse(userId string) (bool, *models.ErrorJson) {
	var exists bool
	query := `SELECT EXISTS (SELECT 1 FROM notifications WHERE recieverId = ? AND seen = 0 LIMIT 1)`

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return false, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	err = stmt.QueryRow(userId).Scan(&exists)
	if err != nil {
		return false, &models.ErrorJson{Status: 500, Error: "", Message: fmt.Sprintf("%v", err)}
	}
	return exists, nil
}

// delete duplicated follow notification
func (repo *NotifRepository) DeleteFollowNotification(userId, authUserId, notifType string) *models.ErrorJson {
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
func (repo *NotifRepository) DeleteGroupNotification(userId, authUserId, notifType, groupId string) *models.ErrorJson {
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
