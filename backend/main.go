package main

import (
	"net/http"
	database "social-network/backend/database/sqlite"
	"social-network/backend/handlers/profile"
	"social-network/backend/routes"
)

func main() {
	_, err := database.InitDB("./database/forum.db")
	if err != nil {
		panic(err)
	}

	mux := routes.SetProfileRoutes(&profile.ProfileHandler{})

	http.ListenAndServe(":8080", mux)
}
