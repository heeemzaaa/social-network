package main

import (
	"fmt"
	"net/http"

	database "social-network/backend/database/sqlite"
	middelware "social-network/backend/middleware"
	"social-network/backend/routes"
)

func main() {
	db, err := database.InitDB("./database/forum.db")
	if err != nil {
		panic(err)
	}

	mux := routes.SetRoutes(db.Database)

	fmt.Println("server is running in : http://localhost:8080")
	http.ListenAndServe(":8080", middelware.NewCorsMiddlerware(middelware.NewRateLimitMiddleWare(mux)))
}
