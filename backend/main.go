package main

import (
	"encoding/json"
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

	http.HandleFunc("/api/posts", middleware.CorsMiddleware(handlers.GetPostsHandler))


	mux := http.NewServeMux()
	routes.SetAuthRoutes(mux, db)
	fmt.Println("server is running in : http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func WriteDataBack(w http.ResponseWriter, data any) {
	fmt.Printf("data: %v\n", data)
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(&data)
}
