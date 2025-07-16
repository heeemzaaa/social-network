package main

import (
	"fmt"
	"log"
	"net/http"

	database "social-network/backend/database/sqlite"
	"social-network/backend/handlers"
	"social-network/backend/middleware"
	ra "social-network/backend/repositories/auth"
	"social-network/backend/routes"
	"social-network/backend/services/auth"
)

// var db *sql.DB

func main() {
	db, err := database.InitDB("./database/forum.db")
	if err != nil {
		panic(err)
	}
	// Create table if not exists
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		notif_id TEXT NOT NULL UNIQUE,
		reciever_Id TEXT NOT NULL,
		sender_Id TEXT NOT NULL,
		seen TEXT NOT NULL,
		notif_type TEXT NOT NULL,
		notif_state TEXT NOT NULL, 
		content TEXT,

		FOREIGN KEY (reciever_Id) REFERENCES users(userID),
		FOREIGN KEY (sender_Id) REFERENCES users(userID)
	);`

	_, err = db.Database.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	mux := routes.SetRoutes(db.Database)
	mux.HandleFunc("/api/posts", handlers.GetPostsHandler)
	mux.HandleFunc("/api/notification", handlers.GetAllNotification)
	mux.HandleFunc("/api/notif", handlers.HandlerNotif)

	fmt.Println("server is running in : http://localhost:8080")
	http.ListenAndServe(":8080", middleware.NewMiddleWare(middleware.NewCorsMiddlerware(mux), auth.NewAuthServer(ra.NewAuthRepository(db.Database))))
}
