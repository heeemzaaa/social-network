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

// select notification by notification_id
func (repo *NotifRepository) SelectNotificationById(notif_id string) (models.Notification, *models.ErrorJson) {
	query := `SELECT * FROM notifications WHERE notifId = ?`

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return models.Notification{}, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	var notification models.Notification
	if err = stmt.QueryRow(notif_id).Scan(
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

// select all notifications by reciever_id
func (repo *NotifRepository) SelectAllNotification(userid string) ([]models.Notification, *models.ErrorJson) {
	all := []models.Notification{}
	query := `SELECT * FROM notifications WHERE recieverId = ?`

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	rows, err := stmt.Query(userid)
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

// select all notifications by userid and type
func (repo *NotifRepository) SelectAllNotificationByType(userid, notifType string) ([]models.Notification, *models.ErrorJson) {
	all := []models.Notification{}
	query := `SELECT * FROM notifications WHERE recieverId = ? AND notifType = ?`

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	rows, err := stmt.Query(userid, notifType)
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

// insert new notification
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
func (repo *NotifRepository) UpdateStatusById(notif_id, status string) *models.ErrorJson {
	query := `UPDATE notifications SET notifStatus = ? WHERE notifId = ?`

	// prepare !!!

	_, err := repo.db.Exec(query, status, notif_id)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	return nil
}

// update notification seen
func (repo *NotifRepository) UpdateSeen(notif_id string) *models.ErrorJson {
	query := `UPDATE notifications SET seen = 1 WHERE notifId = ?`

	// prepare !!!

	_, err := repo.db.Exec(query, notif_id)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	return nil
}

// function check if user has notification containe false seen
func (repo *NotifRepository) IsHasSeenFalse(user_id string) (bool, *models.ErrorJson) {
	var exists bool

	query := `SELECT EXISTS (SELECT 1 FROM notifications WHERE recieverId = ? AND seen = 0 LIMIT 1)`

	// prepare !!!

	err := repo.db.QueryRow(query, user_id).Scan(&exists)
	if err != nil {
		return false, &models.ErrorJson{Status: 500, Error: "", Message: fmt.Sprintf("%v", err)}
	}
	return exists, nil
}

// delete duplicated follow notification
func (repo *NotifRepository) DeleteFollowNotification(userID, authUserID, notifType string) *models.ErrorJson {
	query := `DELETE FROM notifications WHERE senderId = ? AND recieverId = ? AND (notifType = "follow-private" OR notifType = "follow-public")`

	// prepare !!!

	_, err := repo.db.Exec(query, userID, authUserID)
	if err != nil {
		return &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err), Message: "faild to delete notification"}
	}
	return nil
}

// delete duplicated group notification
func (repo *NotifRepository) DeleteGroupNotification(userID, authUserID, notifType, groupId string) *models.ErrorJson {

	// prepare !!!

	query := `DELETE FROM notifications WHERE senderId = ? AND recieverId = ? AND notifType = ? AND groupId = ?`
	_, err := repo.db.Exec(query, userID, authUserID, notifType, groupId)
	if err != nil {
		return &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err), Message: "faild to delete group notification"}
	}
	return nil
}