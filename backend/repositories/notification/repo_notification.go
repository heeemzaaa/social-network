package notification

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"social-network/backend/models"
)

type NotifRepository struct {
	db *sql.DB
}

func NewNotifRepository(db *sql.DB) *NotifRepository {
	return &NotifRepository{db: db}
}

// function check if user has notification containe false seen
func (repo *NotifRepository) IsHasSeenFalse(user_id string) (bool, *models.ErrorJson) {
	var exists bool

	query := `SELECT EXISTS (SELECT 1 FROM notifications WHERE recieverId = ? AND seen = 0 LIMIT 1)`
	err := repo.db.QueryRow(query, user_id).Scan(&exists)
	if err != nil {
		return false, &models.ErrorJson{Status: 500, Error: "", Message: fmt.Sprintf("%v", err)}
	}
	return exists, nil
}

// update notification status value by notif id
func (repo *NotifRepository) UpdateStatusById(notif_id, value string) *models.ErrorJson {
	query := `UPDATE notifications SET notifStatus = ? WHERE notifId = ?`
	_, err := repo.db.Exec(query, value, notif_id)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	return nil
}

// update notification status value by notif id and type
func (repo *NotifRepository) UpdateStatusByType(userID, recieverID, value, notifType string) *models.ErrorJson {
	query := `UPDATE notifications SET notifStatus = ? WHERE senderId = ? AND recieverId = ? AND notifType = ?`
	_, err := repo.db.Exec(query, value, userID, recieverID, notifType)
	if err != nil {
		fmt.Println("ERROR UPDATE STATUS BY TYPE === > ", err)
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	fmt.Println("UPDATE STATE SECCES")
	return nil
}

// update notification seen by notif id
func (repo *NotifRepository) UpdateSeen(notif_id string) *models.ErrorJson {
	query := `UPDATE notifications SET seen = 1 WHERE notifId = ?`
	_, err := repo.db.Exec(query, notif_id)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	// fmt.Println("UPDATE SEEN SECCES")
	return nil
}

// select notification by notif_id
func (repo *NotifRepository) SelectNotification(notif_id string) (models.Notification, *models.ErrorJson) {
	query := `SELECT * FROM notifications WHERE notifId = ?`

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return models.Notification{}, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	var notification models.Notification
	err = stmt.QueryRow(notif_id).Scan(&notification.Id, &notification.RecieverId, &notification.SenderId, &notification.SenderFullName, &notification.Seen, &notification.Type, &notification.Status, &notification.GroupId, &notification.GroupName, &notification.EventId, &notification.CreatedAt)
	if err != nil {
		return notification, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v 1", err)}
	}
	return notification, nil
}

// insert new notification
func (repo *NotifRepository) InsertNewNotification(data models.Notification) *models.ErrorJson {
	if strings.HasPrefix(data.Type, "follow") {
		if errJson := repo.DeleteNotification(data.SenderId, data.RecieverId, data.Type); errJson != nil {
			return errJson
		}
	} else {
		if errJson := repo.DeleteGroupNotification(data.SenderId, data.RecieverId, data.Type, data.GroupId); errJson != nil {
			return errJson
		}
	}
	fmt.Println("dataaaa", data)
	query := `
	INSERT INTO notifications (notifId, recieverId, senderId, senderFullName, seen, notifType, notifStatus, groupId, groupName, eventId, createdAt)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v 3", err)}
	}
	defer stmt.Close()

	_, err = stmt.Exec(data.Id, data.RecieverId, data.SenderId, data.SenderFullName, data.Seen, data.Type, data.Status, data.GroupId, data.GroupName, data.EventId, data.CreatedAt)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v 4", err)}
	}
	return nil
}

// selct all notifications by userid
func (repo *NotifRepository) SelectAllNotification(userid string) ([]models.Notification, *models.ErrorJson) {
	all := []models.Notification{}
	query := `SELECT * FROM notifications WHERE recieverId = ?`

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		fmt.Println("****************************************************************")
		log.Println("Error preparing the query to fetch the notifications: ", err)
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
			fmt.Println("error: ------ rows : next")
			return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
		}
		all = append(all, notification)
	}

	return all, nil
}

// selct all notifications by userid
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
		err = rows.Scan(&notification.Id, &notification.RecieverId, &notification.SenderId, &notification.SenderFullName, &notification.Seen, &notification.Type, &notification.Status, &notification.GroupId, &notification.GroupName, &notification.EventId, &notification.CreatedAt)
		if err != nil {
			fmt.Println("error: ------ rows : next")
			return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
		}
		all = append(all, notification)
	}
	return all, nil
}

// delete the request works in both cases , accept and decline
func (repo *NotifRepository) DeleteGroupNotification(userID, authUserID, notifType, groupId string) *models.ErrorJson {
	query := `DELETE FROM notifications WHERE senderId = ? AND recieverId = ? AND notifType = ? AND groupId = ?`
	_, err := repo.db.Exec(query, userID, authUserID, notifType, groupId)
	if err != nil {
		return &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err), Message: "faild to delete group notification"}
	}
	return nil
}

// delete the request works in both cases , accept and decline
func (repo *NotifRepository) DeleteFollowNotification(userID, authUserID, notifType, status string) *models.ErrorJson {
	query := `DELETE FROM notifications WHERE senderId = ? AND recieverId = ? AND notifType = ? AND notifStatus = ?`
	_, err := repo.db.Exec(query, userID, authUserID, notifType, status)
	if err != nil {
		return &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err), Message: "faild to delete follow notification"}
	}
	return nil
}

func (repo *NotifRepository) DeleteNotification(userID, authUserID, notifType string) *models.ErrorJson {
	query := `DELETE FROM notifications WHERE senderId = ? AND recieverId = ? AND notifType = ?`
	_, err := repo.db.Exec(query, userID, authUserID, notifType)
	if err != nil {
		return &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err), Message: "faild to delete notification"}
	}
	return nil
}

// delete the request works in both cases , accept and decline
func (repo *NotifRepository) DeleteNotifById(notifID string) error {
	query := `DELETE FROM notifications WHERE notifId = ?`
	_, err := repo.db.Exec(query, notifID)
	if err != nil {
		return fmt.Errorf("error deleting the notification from the notifications table: %v", err)
	}
	return nil
}
