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

// function check if user has notification containe false seen
func (repo *NotifRepository) IsHasSeenFalse(user_id string) (bool, *models.ErrorJson) {
	var exists bool

	query := `SELECT EXISTS (SELECT 1 FROM notifications WHERE reciever_Id = ? AND seen = 0 LIMIT 1)`
	err := repo.db.QueryRow(query, user_id).Scan(&exists)
	if err != nil {
		return false, &models.ErrorJson{Status: 500, Error: "", Message: fmt.Sprintf("%v", err)}
	}
	return exists, nil
}

// update notification status value by notif id
func (repo *NotifRepository) UpdateStatusById(notif_id, value string) *models.ErrorJson {
	query := `UPDATE notifications SET notif_state = ? WHERE notif_id = ?`
	_, err := repo.db.Exec(query, value, notif_id)
	if err != nil {
		fmt.Println("ERROR UPDATE === > ", err)
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	fmt.Println("UPDATE STATE SECCES")
	return nil
}
// update notification status value by notif id and type
func (repo *NotifRepository) UpdateStatusByType(userID, recieverID, value, notifType string) *models.ErrorJson {
	query := `UPDATE notifications SET notif_state = ? WHERE sender_Id = ? AND reciever_Id = ? AND notif_type = ?`
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
	query := `UPDATE notifications SET seen = 1 WHERE notif_id = ?`
	_, err := repo.db.Exec(query, notif_id)
	if err != nil {
		fmt.Println("ERROR UPDATE === > ", err)
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	// fmt.Println("UPDATE SEEN SECCES")
	return nil
}

// select notification by notif_id
func (repo *NotifRepository) SelectNotification(notif_id string) (models.Notification, *models.ErrorJson) {
	query := `SELECT * FROM notifications WHERE notif_id = ?`

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return models.Notification{}, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	var notification models.Notification
	err = stmt.QueryRow(notif_id).Scan(&notification.Id, &notification.Reciever_Id, &notification.Sender_Id, &notification.Seen, &notification.Type, &notification.Status, &notification.Content, &notification.CreatedAt)
	if err != nil {
		return notification, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v 1", err)}
	}
	return notification, nil
}

// insert new notification
func (repo *NotifRepository) InsertNewNotification(data models.Notification) *models.ErrorJson {
	fmt.Println("dataaaa", data)
	query := `
	INSERT INTO notifications (notif_id, reciever_Id, sender_Id, seen, notif_type, notif_state, content, createdAt)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`
	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v 3", err)}
	}
	defer stmt.Close()

	_, err = stmt.Exec(data.Id, data.Reciever_Id, data.Sender_Id, data.Seen, data.Type, data.Status, data.Content, data.CreatedAt)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v 4", err)}
	}
	return nil
}

// selct all notifications by userid
func (repo *NotifRepository) SelectAllNotification(userid string) ([]models.Notification, *models.ErrorJson) {
	all := []models.Notification{}
	query := `SELECT * FROM notifications WHERE reciever_Id = ?`

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	rows, err := stmt.Query(userid)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("sql.no.rows: no exist data -----")
			return nil, nil
		}
		fmt.Println("sql.no.error: no exist data -----")
	}
	defer rows.Close()

	for rows.Next() {
		var notification models.Notification
		err = rows.Scan(&notification.Id, &notification.Reciever_Id, &notification.Sender_Id, &notification.Seen, &notification.Type, &notification.Status, &notification.Content, &notification.CreatedAt)
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
	query := `SELECT * FROM notifications WHERE reciever_Id = ? AND notif_type = ?`

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	rows, err := stmt.Query(userid, notifType)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("sql.no.rows: no exist data -----")
			return nil, nil
		}
		fmt.Println("sql.no.error: no exist data -----")
	}
	defer rows.Close()

	for rows.Next() {
		var notification models.Notification
		err = rows.Scan(&notification.Id, &notification.Reciever_Id, &notification.Sender_Id, &notification.Seen, &notification.Type, &notification.Status, &notification.Content, &notification.CreatedAt)
		if err != nil {
			fmt.Println("error: ------ rows : next")
			return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
		}
		all = append(all, notification)
	}
	return all, nil
}

// delete the request works in both cases , accept and decline
func (repo *NotifRepository) DeleteNotification(userID, authUserID, notifType string) *models.ErrorJson {
	query := `DELETE FROM notifications WHERE sender_Id = ? AND reciever_Id = ? AND notif_type = ?`
	_, err := repo.db.Exec(query, userID, authUserID, notifType)
	if err != nil {
		return &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err), Message: "error deleting the notification from the notifications"}
	}
	return nil
}

// delete the request works in both cases , accept and decline
func (repo *NotifRepository) DeleteNotifById(notifID string) error {
	query := `DELETE FROM notifications WHERE noti_id = ?`
	_, err := repo.db.Exec(query, notifID)
	if err != nil {
		return fmt.Errorf("error deleting the notification from the notifications table: %v", err)
	}
	return nil
}
