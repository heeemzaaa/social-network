package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"social-network/backend/models"
	"social-network/backend/utils"
	"strconv"
)

func GetLastNotifications(w http.ResponseWriter, r *http.Request) {
	fmt.Println("START --------")
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "GET" {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Message: "ERROR!! Method Not Allowed!"})
		return
	}

	cookie, err := r.Cookie("session")
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 401, Error: "Unauthorized Access"})
		return
	}

	db, err := sql.Open("sqlite3", "file:./database/forum.db?_foreign_keys=1") // or your DB driver
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Error: "Internal server error", Message: "Failed to connect to database:"})
		return
	}
	defer db.Close()

	session := models.Session{}
	query := `SELECT sessions.userID, sessions.sessionToken , sessions.expiresAt, users.nickname 
	FROM sessions INNER JOIN users ON users.userID = sessions.userID
	WHERE sessionToken = ?`

	row := db.QueryRow(query, cookie.Value).Scan(&session.UserId, &session.Token, &session.ExpDate, &session.Username)
	if row == sql.ErrNoRows {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 401, Message: " Unauthorized Access"})
		return
	}

	// ***** get last notif from data base ***** //

	private := models.NewNotification()
	private.Id = "562b7f42-b132-4a5d-8863-ed111c6487eb"
	private.Reciever_Id = "562b7f42-b132-4a5d-8863-ed111c6487eb"
	private.Sender_Id = "946257db-65e0-491e-9a09-c33f946f5092"
	// private.Seen = false
	private.Type = "follow-private"
	private.Status = "later"
	private.Content = "SENDER_NAME sent follow request"

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(*private)
	fmt.Println("END -------")

}

func GetAllNotification(w http.ResponseWriter, r *http.Request) {
	fmt.Println("START --------")
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "GET" {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Message: "ERROR!! Method Not Allowed!"})
		return
	}

	cookie, err := r.Cookie("session")
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 401, Error: "Unauthorized Access"})
		return
	}

	db, err := sql.Open("sqlite3", "file:./database/forum.db?_foreign_keys=1") // or your DB driver
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Error: "Internal server error", Message: "Failed to connect to database:"})
		return
	}
	defer db.Close()

	session := models.Session{}
	query := `SELECT sessions.userID, sessions.sessionToken , sessions.expiresAt, users.nickname 
	FROM sessions INNER JOIN users ON users.userID = sessions.userID
	WHERE sessionToken = ?`

	row := db.QueryRow(query, cookie.Value).Scan(&session.UserId, &session.Token, &session.ExpDate, &session.Username)
	if row == sql.ErrNoRows {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 401, Message: " Unauthorized Access"})
		return
	}

	fmt.Println("session ---------", session)

	fmt.Println("session.UserId === ", session.UserId)

	value := r.URL.Query().Get("Count")
	fmt.Println(value)
	// // // //  getUserNotifications by userID
	allNotifications := getUserNotificationsByID(session.UserId, db)
	// get data example
	if len(allNotifications) <= 10 {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(allNotifications)
		return
	}

	if value != "" {
		length := len(allNotifications)
		nbr, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println("Invalid query != number")
		}
		fmt.Println(nbr)

		switch {
		case length <= 10:

			return

		case nbr <= 0:
			allNotifications = allNotifications[:10]
			fmt.Println("nbr <= 0")

		case length <= nbr:
			allNotifications = []models.Notifications{}
			fmt.Println("greates than notifications : length <= nbr")

		case length < nbr+10:
			allNotifications = allNotifications[nbr:]
			fmt.Println("length < nbr + 10")

		default:
			allNotifications = allNotifications[nbr : nbr+10]
			fmt.Println("default = [nbr : nbr + 10]")
		}
	}

	// // fmt.Println(allNotifications)
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(allNotifications)
	fmt.Println("END -------")

}

func getUserNotificationsByID(_ string, _ *sql.DB) []models.Notifications {
	fmt.Println("geeet ----- get user notifications by id ***")

	// data := []models.Notifications{}
	// query := `SELECT * FROM notifications WHERE reciever_Id=?`

	// rows, err := db.Query(query, id)
	// if err != nil {
	// 	models.NewErrorJson(500, "Internal Server error", nil)
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	var row models.Notifications
	// 	err := rows.Scan(&row)
	// 	if err != nil {
	// 		fmt.Println("error: ------ rows : next")
	// 	}
	// 	data = append(data, row)
	// }

	// fmt.Println("data: --------", data)

	// if len(data) != 0 {
	// 	return data
	// }

	// get data with query param request
	allNotifications := []models.Notifications{}

	private := models.NewNotification()
	private.Id = "562b7f42-b132-4a5d-8863-ed111c6487eb"
	private.Reciever_Id = "562b7f42-b132-4a5d-8863-ed111c6487eb"
	private.Sender_Id = "946257db-65e0-491e-9a09-c33f946f5092"
	// private.Seen = false
	private.Type = "follow-private"
	private.Status = "later"
	private.Content = "SENDER_NAME sent follow request"
	allNotifications = append(allNotifications, *private)

	public := models.NewNotification()
	public.Id = "562b7f42-b132-4a5d-8863-ed1116487eb"
	public.Reciever_Id = "562b7f42-b132-4a5d-8863-ed111c6487eb"
	public.Sender_Id = "946257db-65e0-491e-9a09-c33f946f5092"
	// public.Seen = false
	public.Type = "follow-public"
	public.Status = "later"
	public.Content = "SENDER_NAME follow you"
	allNotifications = append(allNotifications, *public)

	invitation := models.NewNotification()
	invitation.Id = "562b7f42-b132-4a5d-8863-ed1c6487eb"
	invitation.Reciever_Id = "562b7f42-b132-4a5d-8863-ed111c6487eb"
	invitation.Sender_Id = "946257db-65e0-491e-9a09-c33f946f5092"
	// invitation.Seen = false
	invitation.Type = "group-invitation"
	invitation.Status = "later"
	invitation.Content = "SENDER_NAME "
	allNotifications = append(allNotifications, *invitation)

	private1 := models.NewNotification()
	private1.Id = "562b7f42-b132-4a5d863-ed111c6487e1"
	private1.Reciever_Id = "562b7f42-b132-4a5d-8863-ed111c6487eb"
	private1.Sender_Id = "946257db-65e0-491e-9a09-c33f946f5092"
	// private1.Seen = false
	private1.Type = "follow-private"
	private1.Status = "later"
	private1.Content = "SENDER_NAME sent follow request"
	allNotifications = append(allNotifications, *private1)

	public1 := models.NewNotification()
	public1.Id = "562b7f42-b1a5d-8863-ed111c6487e1"
	public1.Reciever_Id = "562b7f42-b132-4a5d-8863-ed111c6487eb"
	public1.Sender_Id = "946257db-65e0-491e-9a09-c33f946f5092"
	// public1.Seen = false
	public1.Type = "follow-public"
	public1.Status = "later"
	public1.Content = "SENDER_NAME follow you"
	allNotifications = append(allNotifications, *public1)

	invitation1 := models.NewNotification()
	invitation1.Id = "56f42-b132-4a5d-8863-ed111c6487e1"
	invitation1.Reciever_Id = "562b7f42-b132-4a5d-8863-ed111c6487eb"
	invitation1.Sender_Id = "946257db-65e0-491e-9a09-c33f946f5092"
	// invitation1.Seen = false
	invitation1.Type = "group-invitation"
	invitation1.Status = "later"
	invitation1.Content = "SENDER_NAME "
	allNotifications = append(allNotifications, *invitation1)

	private2 := models.NewNotification()
	private2.Id = "562b7f42-b132-4a5d8863-ed111c6487e2"
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

	return allNotifications
}
