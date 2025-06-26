package main

import database "social-network/backend/database/sqlite"

func main() {
	_, err := database.InitDB("./database/forum.db")
	if err != nil {
		panic(err)
	}
}
