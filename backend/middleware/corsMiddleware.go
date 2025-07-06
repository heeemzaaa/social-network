package middleware

import (
	"fmt"
	"net/http"
)

func CorsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	fmt.Println("here")
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000/")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		// if r.Method == http.MethodOptions {
		// 	w.WriteHeader(http.StatusOK)
		// 	return
		// }
		next(w, r)
	}
}
