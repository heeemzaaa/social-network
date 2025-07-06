package main

import (
	"fmt"
	"net/http"

	database "social-network/backend/database/sqlite"
	"social-network/backend/handlers"
	"social-network/backend/middleware"
	"social-network/backend/routes"
)

func main() {
	db, err := database.InitDB("./database/forum.db")
	if err != nil {
		panic(err)
	}

	// just for testing
	http.HandleFunc("/api/posts", middleware.CorsMiddleware(handlers.GetPostsHandler))
	mux := routes.SetRoutes(db.Database)
	fmt.Println("server is running in : http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
