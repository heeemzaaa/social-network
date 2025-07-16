package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"social-network/backend/models"
)

// type Notifications struct {
// 	Id          string
// 	Recieved_Id string // or name
// 	Sender_name string
// 	Seen        bool
// 	Type        string // follow request // invite group // join group to admin // event notif for members group // public request . bonus
// 	Status      string // accepted or rejected or "" for public request
// 	Content     string
// 	// Created_at time.Time
// }

// get notifications by user id

func HandleNotification(w http.ResponseWriter, r *http.Request) {
	// follow request type

}

func SingleNotification(w http.ResponseWriter, r *http.Request) {
	// method get
	fmt.Println("request heeeeeeereeeee ----------")

	data := models.Notifications{
		Id:          "RTR467T",
		Reciever_Id: "222222", // or name,
		Sender_Id:   "amine",
		Seen:        false,
		Type:        "follow request", // follow request // invite group // join group to admin // event notif for members group // public request . bonus,
		Status:      "empty",          // accepted or rejected or "" for public request,
		Content:     "content",
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(data)
}
