package main

import (
	"encoding/json"
	"net/http"

	database "social-network/backend/database/sqlite"
	"social-network/backend/models"
)

func main() {
	_, err := database.InitDB("./database/forum.db")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/api/test", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		test := &models.Test{Message: "this is a test"}
		// this is for decoding the request body
		// err := json.NewDecoder(r.Body).Decode(&post)
		WriteDataBack(w, test)
	})

	http.ListenAndServe(":8080", nil)
}

func WriteDataBack(w http.ResponseWriter, data any) {
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(&data)
}
