package middleware

import (
	"net/http"
)

type CorsMiddleware struct {
	handler http.Handler
}

func NewCorsMiddlerware(handler http.Handler) *CorsMiddleware {
	return &CorsMiddleware{handler}
}

func (m *CorsMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// switch origin := r.Header.Get("Origin"); origin {
	// case "http://localhost:3000", "http://frontend:3000":
	// 	w.Header().Set("Access-Control-Allow-Origin", origin)
	// }
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, OPTIONS, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	m.handler.ServeHTTP(w, r)
}
