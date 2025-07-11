package main

import (
	"fmt"
	"net/http"

	database "social-network/backend/database/sqlite"
	"social-network/backend/handlers"
	"social-network/backend/middleware"
	ra "social-network/backend/repositories/auth"
	"social-network/backend/routes"
	"social-network/backend/services/auth"
)

func main() {
	db, err := database.InitDB("./database/forum.db")
	if err != nil {
		panic(err)
	}

	mux := routes.SetRoutes(db.Database)
	mux.HandleFunc("/api/posts", handlers.GetPostsHandler)

	fmt.Println("server is running in : http://localhost:8080")
	http.ListenAndServe(":8080", middleware.NewCorsMiddlerware(middleware.NewMiddleWare(mux, auth.NewAuthServer(ra.NewAuthRepository(db.Database)))))
}
