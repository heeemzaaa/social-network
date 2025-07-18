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

func (repo *NotifRepository) InsertNewNotification(data models.Notification) *models.ErrorJson {
	fmt.Println("INSERT: -------- data = ", data)
	query := `
	INSERT INTO notifications (notif_id, reciever_Id, sender_Id, seen,notif_type, notif_state, content)
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
func (repo *NotifRepository) GetUserNameByUserId(user_id string) (bool, string, *models.ErrorJson) {
	var nickname string
	query := `SELECT * FROM users WHERE userID = ?`
	row := repo.db.QueryRow(query, user_id)
	// fmt.Println("row : ", row)
	err := row.Scan(&nickname)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, "res 11111", nil
		}
		fmt.Printf("Scan error: %v\n", err)

		return false, "res 22222", &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	return true, nickname, nil
}

func (repo *NotifRepository) SelectNotification(data string) (models.Notification, *models.ErrorJson) {

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

func (repo *NotifRepository) SelectAllNotification(userid string) ([]models.Notification, *models.ErrorJson) {
	// all := []models.Notification{}
	// query := `SELECT * FROM notifications WHERE reciever_Id=?`

	// stmt, err := repo.db.Prepare(query)
	// if err != nil {
	// 	return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	// }
	// defer stmt.Close()

	// rows, err := stmt.Query(userid)
	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		fmt.Println("sql.no.rows: no exist data -----")
	// 		return nil, nil
	// 	}
	// 	fmt.Println("sql.no.error: no exist data -----")
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	fmt.Println("inside rows.next -----")
	// 	var notification models.Notification
	// 	err = rows.Scan(&notification)
	// 	if err != nil {
	// 		fmt.Println("error: ------ rows : next")
	// 		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	// 	}
	// 	all = append(all, notification)
	// }

	// fmt.Println("allNotifications: --------", all)

	// return all, nil

	// get data with query param request
	allNotifications := []models.Notification{}

	private := models.NewNotification()
	private.Id = "562b2-b132-4a5d-8863-ed111c6487eb"
	private.Reciever_Id = "562b7f42-b132-4a5d-8863-ed111c6487eb"
	private.Sender_Id = "946257db-65e0-491e-9a09-c33f946f5092"
	// private.Seen = false
	private.Type = "follow-private"
	private.Status = "later"
	private.Content = "SENDER_NAME sent follow request"
	allNotifications = append(allNotifications, *private)

	public := models.NewNotification()
	public.Id = "562b7f42-b2-4a5d-8863-ed1116487eb"
	public.Reciever_Id = "562b7f42-b132-4a5d-8863-ed111c6487eb"
	public.Sender_Id = "946257db-65e0-491e-9a09-c33f946f5092"
	// public.Seen = false
	public.Type = "follow-public"
	public.Status = "later"
	public.Content = "SENDER_NAME follow you"
	allNotifications = append(allNotifications, *public)

	invitation := models.NewNotification()
	invitation.Id = "562b7f42-b-4a5d-8863-ed1c6487eb"
	invitation.Reciever_Id = "562b7f42-b132-4a5d-8863-ed111c6487eb"
	invitation.Sender_Id = "946257db-65e0-491e-9a09-c33f946f5092"
	// invitation.Seen = false
	invitation.Type = "group-invitation"
	invitation.Status = "later"
	invitation.Content = "SENDER_NAME "
	allNotifications = append(allNotifications, *invitation)

	private1 := models.NewNotification()
	private1.Id = "562b7f42-b135d863-ed111c6487e1"
	private1.Reciever_Id = "562b7f42-b132-4a5d-8863-ed111c6487eb"
	private1.Sender_Id = "946257db-65e0-491e-9a09-c33f946f5092"
	// private1.Seen = false
	private1.Type = "follow-private"
	private1.Status = "later"
	private1.Content = "SENDER_NAME sent follow request"
	allNotifications = append(allNotifications, *private1)

	public1 := models.NewNotification()
	public1.Id = "562b7f42-b1a5863-ed111c6487e1"
	public1.Reciever_Id = "562b7f42-b132-4a5d-8863-ed111c6487eb"
	public1.Sender_Id = "946257db-65e0-491e-9a09-c33f946f5092"
	// public1.Seen = false
	public1.Type = "follow-public"
	public1.Status = "later"
	public1.Content = "SENDER_NAME follow you"
	allNotifications = append(allNotifications, *public1)

	invitation1 := models.NewNotification()
	invitation1.Id = "56f42-b135d-8863-ed111c6487e1"
	invitation1.Reciever_Id = "562b7f42-b132-4a5d-8863-ed111c6487eb"
	invitation1.Sender_Id = "946257db-65e0-491e-9a09-c33f946f5092"
	// invitation1.Seen = false
	invitation1.Type = "group-invitation"
	invitation1.Status = "later"
	invitation1.Content = "SENDER_NAME "
	allNotifications = append(allNotifications, *invitation1)

	private2 := models.NewNotification()
	private2.Id = "562b7f42-b132-4a5d63-ed111c6487e2"
	private2.Reciever_Id = "562b7f42-b132-4a5d-8863-ed111c6487eb"
	private2.Sender_Id = "946257db-65e0-491e-9a09-c33f946f5092"
	// private2.Seen = false
	private2.Type = "follow-private"
	private2.Status = "later"
	private2.Content = "SENDER_NAME sent follow request"
	allNotifications = append(allNotifications, *private2)

	public2 := models.NewNotification()
	public2.Id = "562b7f42-b32-4a5d-8863-ed111c6487e2"
	public2.Reciever_Id = "562b7f42-b132-4a5d-8863-ed111c6487eb"
	public2.Sender_Id = "946257db-65e0-491e-9a09-c33f946f5092"
	// public2.Seen = false
	public2.Type = "follow-public"
	public2.Status = "later"
	public2.Content = "SENDER_NAME follow you"
	allNotifications = append(allNotifications, *public2)

	invitation2 := models.NewNotification()
	invitation2.Id = "62b7f42-b132-4a5d-8863-ed111c6487e2"
	invitation2.Reciever_Id = "562b7f42-b132-4a5d-8863-ed111c6487eb"
	invitation2.Sender_Id = "946257db-65e0-491e-9a09-c33f946f5092"
	// invitation2.Seen = false
	invitation2.Type = "group-invitation"
	invitation2.Status = "later"
	invitation2.Content = "SENDER_NAME "
	allNotifications = append(allNotifications, *invitation2)

	private3 := models.NewNotification()
	private3.Id = "b7f42-b132-4a5d-8863-ed111c6487e3"
	private3.Reciever_Id = "562b7f42-b132-4a5d-8863-ed111c6487eb"
	private3.Sender_Id = "946257db-65e0-491e-9a09-c33f946f5092"
	// private3.Seen = false
	private3.Type = "follow-private"
	private3.Status = "later"
	private3.Content = "SENDER_NAME sent follow request"
	allNotifications = append(allNotifications, *private3)

	public3 := models.NewNotification()
	public3.Id = "562b7f-b132-4a5d-8863-ed111c64e3"
	public3.Reciever_Id = "562b7f42-b132-4a5d-8863-ed111c6487eb"
	public3.Sender_Id = "946257db-65e0-491e-9a09-c33f946f5092"
	// public3.Seen = false
	public3.Type = "follow-public"
	public3.Status = "later"
	public3.Content = "SENDER_NAME follow you"
	allNotifications = append(allNotifications, *public3)

	invitation3 := models.NewNotification()
	invitation3.Id = "562b7f42-b13-4a5d-886-ed116487e3"
	invitation3.Reciever_Id = "562b7f42-b132-4a5d-8863-ed111c6487eb"
	invitation3.Sender_Id = "946257db-65e0-491e-9a09-c33f946f5092"
	// invitation3.Seen = false
	invitation3.Type = "group-invitation"
	invitation3.Status = "later"
	invitation3.Content = "SENDER_NAME "
	allNotifications = append(allNotifications, *invitation3)

	return allNotifications, nil
}
