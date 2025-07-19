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

// 3 -- function check if user has notification containe false seen
func (repo *NotifRepository) IsHasSeenFalse(user_id string) (int, *models.ErrorJson) {

	// get data with query param request
	// private := models.Notification{}
	// private.Id = "YYYYYYYYhed111c6487eb"
	// private.Reciever_Id = "555555dddddddd863-ed111c6487eb"
	// private.Sender_Id = "uuid.New().String()"
	// private.Seen = false
	// private.Type = "follow-private"
	// private.Status = "later"
	// private.Content = "SENDER_NAME sent follow request"
	return 0, nil
}
// 4 -- function update status
func (repo *NotifRepository) UpdateStatus(notif_id, user_id string) (bool, *models.ErrorJson) {

	// get data with query param request
	// private := models.Notification{}
	// private.Id = "YYYYYYYYhed111c6487eb"
	// private.Reciever_Id = "555555dddddddd863-ed111c6487eb"
	// private.Sender_Id = "uuid.New().String()"
	// private.Seen = false
	// private.Type = "follow-private"
	// private.Status = "later"
	// private.Content = "SENDER_NAME sent follow request"
	return true, nil
}
// 5 -- function update seen
func (repo *NotifRepository) UpdateSeen(notif_id, user_id string) (bool, *models.ErrorJson) {

	// get data with query param request
	// private := models.Notification{}
	// private.Id = "YYYYYYYYhed111c6487eb"
	// private.Reciever_Id = "555555dddddddd863-ed111c6487eb"
	// private.Sender_Id = "uuid.New().String()"
	// private.Seen = false
	// private.Type = "follow-private"
	// private.Status = "later"
	// private.Content = "SENDER_NAME sent follow request"
	return true, nil
}



// select notification by notif_id
func (repo *NotifRepository) SelectNotification(notif_id string) (models.Notification, *models.ErrorJson) {

	// get data with query param request
	private := models.Notification{}
	private.Id = "YYYYYYYYhed111c6487eb"
	private.Reciever_Id = "555555dddddddd863-ed111c6487eb"
	private.Sender_Id = "uuid.New().String()"
	private.Seen = false
	private.Type = "follow-private"
	private.Status = "later"
	private.Content = "SENDER_NAME sent follow request"
	return private, nil
}
// insert new notification
func (repo *NotifRepository) InsertNewNotification(data models.Notification) *models.ErrorJson {
	// fmt.Println("INSERT NEW NOTIFICATION: -------- data = ", data)
	query := `
	INSERT INTO notifications (notif_id, reciever_Id, sender_Id, seen, notif_type, notif_state, content)
	VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		fmt.Println("INSERT: ERR 111 ", err)
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v 3", err)}
	}
	defer stmt.Close()

	_, err = stmt.Exec(data.Id, data.Reciever_Id, data.Sender_Id, data.Seen, data.Type, data.Status, data.Content)
	if err != nil {
		fmt.Println("INSERT: ERR 222 ", err)
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
		// fmt.Println("inside rows.next -----")
		var notification models.Notification
		err = rows.Scan(&notification.Id, &notification.Reciever_Id, &notification.Sender_Id, &notification.Seen, &notification.Type, &notification.Status, &notification.Content)
		if err != nil {
			fmt.Println("error: ------ rows : next")
			return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
		}
		all = append(all, notification)
	}

	// fmt.Println("allNotifications: --------", all)

	return all, nil
}


// func (repo *NotifRepository) GetUserNameByUserId(user_id string) (bool, string, *models.ErrorJson) {
// 	var nickname string
// 	query := `SELECT * FROM users WHERE userID = ?`
// 	row := repo.db.QueryRow(query, user_id)
// 	// fmt.Println("row : ", row)
// 	err := row.Scan(&nickname)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return false, "res 11111", nil
// 		}
// 		fmt.Printf("Scan error: %v\n", err)
// 		return false, "res 22222", &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
// 	}
// 	return true, nickname, nil
// }
