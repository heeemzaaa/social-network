package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"slices"
	"social-network/backend/models"
	"social-network/backend/utils"

	"github.com/google/uuid"
)

// func WriteJsonErrors(w http.ResponseWriter, errJson models.ErrorJson) {
// 	w.WriteHeader(errJson.Status)
// 	json.NewEncoder(w).Encode(errJson)
// }

// type ErrorJson struct {
// 	Status  int    `json:"status"`
// 	Error   string `json:"error"`
// 	Message any    `json:"errors"`
// }

// func NewErrorJson(status int, err string, message any) *ErrorJson {
// 	return &ErrorJson{
// 		Status:  status,
// 		Error:   err,
// 		Message: message,
// 	}
// }

func HandlerNotif(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var err error
	if r.Method != "POST" {
		fmt.Println("error: method != post ---------")
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Error: "method not allowed - 405", Message: "method not allowed"})
		return
	}

	var Data models.Notif

	err = json.NewDecoder(r.Body).Decode(&Data)
	if err != nil {
		fmt.Println("invalide decode lol", Data)
	}
	fmt.Println("daaaaaaataaaaaaaa=", Data)

	db, err := sql.Open("sqlite3", "file:./database/forum.db?_foreign_keys=1") // or your DB driver
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	switch Data.Type {
	case "follow-private":
		followPrivateProfile(w, Data, db)
	case "follow-public":
		followPublicProfile(w, Data, db)
	case "group-invitation":
		groupInvitationRequest(w, Data, db)
	case "group-join":
		groupJoinRequest(w, Data, db)
	case "group-event":
		groupEventRequest(w, Data, db)
	default:
		fmt.Println("error: bad request = invalid type ---------")
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: "Bad Request - 400", Message: "invalid type"})
		return
	}

}

func CreateNotification(notification *models.Notifications, db *sql.DB) *models.ErrorJson {
	// fmt.Printf("Notification fields: Id=%s, Reciever_Id=%s, Sender_Id=%s, Type=%s, Status=%s, Content=%s\n",
	// notification.Id, notification.Reciever_Id, notification.Sender_Id, notification.Type, notification.Status, notification.Content)

	query := `INSERT INTO notifications (notif_id, reciever_Id, sender_Id, seen, notif_type, notif_state, content) VALUES (?,?,?,?,?,?,?)`

	if _, err := db.Exec(query, notification.Id, notification.Reciever_Id, notification.Sender_Id, "false", notification.Type, notification.Status, notification.Content); err != nil {
		fmt.Println("error 222222222 ------", err)
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	fmt.Println("valid ------")
	return nil
}

func GetUserNameByUserId(user_id string, db *sql.DB) (string, *models.ErrorJson, bool) {

	var nickname string
	query := `SELECT nickname FROM users WHERE userID = ?`
	row := db.QueryRow(query, user_id)
	// fmt.Println("row : ", row)
	err := row.Scan(&nickname)
	if err != nil {
		if err == sql.ErrNoRows {
			return "res 11111", nil, false
		}
		fmt.Printf("Scan error: %v\n", err)

		return "res 22222", &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}, false
	}
	return nickname, nil, true
}

// 1 - follow private profile request
// event : follow profile button : onclick()
func followPrivateProfile(w http.ResponseWriter, data models.Notif, db *sql.DB) {
	fmt.Println("HeeeeeReeee: --------- //// 1111")
	fmt.Println(data.Sender_Id)

	// get user data by sender_id or bad request if sender_id not exist
	ss, _, exist := GetUserNameByUserId(data.Sender_Id, db)
	if !exist {
		// bad requet
		fmt.Println("error : bad request")
	}
	fmt.Println("sender name = ----", ss)
	sender := ss

	// get user data by reciever_id or bad request if reciever_id not exist
	rr, _, exist := GetUserNameByUserId(data.Reciever_Id, db)
	if !exist {
		// bad requet
		fmt.Println("error : bad request")
	}
	fmt.Println("reciever name = ----", rr)
	reciever := rr // get user_id from users where user_profile_id = data.reciever_id

	notification := models.NewNotification()

	// generate new notification id
	notification.Id = uuid.New().String()

	notification.Sender_Id = data.Sender_Id     // valid id
	notification.Reciever_Id = data.Reciever_Id // valid id

	// notification.Seen = "false"
	notification.Type = data.Type
	notification.Status = "later"
	notification.Content = fmt.Sprintf("%v sent follow request %v", sender, reciever)

	// insert into notifications new row notification data
	// CreateNotification(notification, db)

	fmt.Println(notification)
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(notification)
}

// 2 - follow public profile request
// event : follow profile button : onclick()
func followPublicProfile(w http.ResponseWriter, data models.Notif, db *sql.DB) {
	// get user data by sender_id or bad request if sender_id not exist
	sName, _, exist := GetUserNameByUserId(data.Sender_Id, db)
	if !exist {
		// bad requet
		fmt.Println("error : bad request")
	}
	fmt.Println("sender name = ----", sName)
	sender := sName

	// get user data by reciever_id or bad request if reciever_id not exist
	rName, _, exist := GetUserNameByUserId(data.Reciever_Id, db)
	if !exist {
		// bad requet
		fmt.Println("error : bad request")
	}
	fmt.Println("reciever name = ----", rName)
	reciever := rName // get user_id from users where user_profile_id = data.reciever_id

	notification := models.NewNotification()

	// generate new notification id
	notification.Id = uuid.New().String()

	notification.Sender_Id = data.Sender_Id
	notification.Reciever_Id = data.Reciever_Id

	// notification.Seen =    false
	notification.Type = data.Type
	notification.Status = "direct"
	notification.Content = sender + " sent follow request " + reciever

	// insert into notifications new row notification data
	// CreateNotification(notification, db)

	fmt.Println(notification)
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(notification)
}

// 3 - group invitation request
// event : select invitation button : onclick()
func groupInvitationRequest(w http.ResponseWriter, data models.Notif, db *sql.DB) {

	// get user data by sender_id or bad request if sender_id not exist //// **** in the same group **** ////
	sName, _, exist := GetUserNameByUserId(data.Sender_Id, db)
	if !exist {
		fmt.Println("error : bad request")
	}
	fmt.Println("sender name = ----", sName)
	sender := sName

	// get user data by reciever_id or bad request if reciever_id exist //// **** in the same group **** ////
	rName, _, exist := GetUserNameByUserId(data.Reciever_Id, db)
	if !exist {
		fmt.Println("error : bad request")
	}
	fmt.Println("reciever name = ----", rName)
	// reciever := rName // get user_id from users where user_profile_id = data.reciever_id

	// check if group exist by id or name = data.Content:
	// gName, _, exist := GetUserNameByUserId(data.Content, db)
	// if !exist {
	// 	fmt.Println("error : bad request")
	// }
	// fmt.Println("sender name = ----", gName)
	club := "club test" // gName

	notification := models.NewNotification()

	// generate new notification id
	notification.Id = uuid.New().String()

	notification.Reciever_Id = data.Reciever_Id
	notification.Sender_Id = data.Sender_Id

	// notification.Seen =    false
	notification.Type = data.Type
	notification.Status = "later"
	notification.Content = sender + " invited to join group " + club // check if group_name exist

	// insert into notifications new row notification data
	// CreateNotification(notification, db)

	fmt.Println(notification)
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(notification)
}

// 4 - group join request [admin] //// ------ get reciever id from admin_group
// event : group join button : onclick()
func groupJoinRequest(w http.ResponseWriter, data models.Notif, db *sql.DB) {

	// get user data by sender_id or bad request if sender_id not exist //// **** in the same group **** ////
	sName, _, exist := GetUserNameByUserId(data.Sender_Id, db)
	if !exist {
		fmt.Println("error : bad request")
	}
	fmt.Println("sender name = ----", sName)
	sender := sName

	// get reciever_id = "group_admin_user_id" // get group_admin_user_id from groups where group_name or id = data.content

	// check if group exist by ID or G_NAME = data.Content:
	// group, _, exist := GetGroupById(data.Content, db)
	// if !exist {
	// 	fmt.Println("error : bad request")
	// }
	// fmt.Println("sender name = ----", gName)
	club := "club test" // group.groupCreatorID

	notification := models.NewNotification()

	// generate new notification id
	notification.Id = uuid.New().String()

	notification.Reciever_Id = data.Reciever_Id
	notification.Sender_Id = data.Sender_Id

	// notification.Seen =    false
	notification.Type = data.Type
	notification.Status = "later"
	notification.Content = sender + " want join " + club

	// insert into notifications new row notification data
	// CreateNotification(notification, db)

	fmt.Println(notification)
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(notification)
}

// 5 - group event created [group-members]
// event : create group event button : onclick()
func groupEventRequest(w http.ResponseWriter, data models.Notif, db *sql.DB) {

	// get user data by sender_id or bad request if sender_id not exist //// **** in the same group **** ////
	sName, _, exist := GetUserNameByUserId(data.Sender_Id, db)
	if !exist {
		fmt.Println("error : bad request")
	}
	fmt.Println("sender name = ----", sName)
	sender := sName

	// check if group_name exist
	// check if group exist by ID or G_NAME = data.Content:
	// gName, _, exist := GetUserNameByUserId(data.Content, db)
	// if !exist {
	// 	fmt.Println("error : bad request")
	// }
	// fmt.Println("sender name = ----", gName)
	club := "club test" // gName

	// get user_ids from groupMembers table for row.Next()

	groupMembers := []string{"b3cc45f5-b1bb-4c65-871c-75dfe8b5692f", "9f57a431-adf6-42b2-94f3-89c4496d8932", "562b7f42-b132-4a5d-8863-ed111c6487eb", "946257db-65e0-491e-9a09-c33f946f5092"} // get group_admin_user_id from groups where group_name = data.content [group_name]
	var result []models.Notifications

	// check if the sender in the group
	if !slices.Contains(groupMembers, data.Sender_Id) {
		fmt.Println("error : bad request")
	}

	for _, member_Id := range groupMembers {

		if member_Id == data.Sender_Id {
			continue
		}

		notification := models.NewNotification()

		// generate new notification id
		notification.Id = uuid.New().String()

		notification.Reciever_Id = member_Id
		notification.Sender_Id = data.Sender_Id

		// notification.Seen =    false
		notification.Type = data.Type
		notification.Status = "later"
		notification.Content = sender + " create event at club " + club

		// insert into notifications new row result data
		// CreateNotification(notification, db)

		fmt.Println(notification)
		result = append(result, *notification)
	}

	fmt.Println(result)
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(result)
}
