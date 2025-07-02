package middleware

// import (
// 	"encoding/json"
// 	"net/http"
// 	"sync"
// 	"time"
// )

// type Middleware struct {
// 	MiddlewareHanlder http.Handler
// 	service           *service.AppService
// }

// type LoginRegisterMiddleWare struct {
// 	MiddlewareHanlder http.Handler
// 	service           *service.AppService
// }

// type UserInfo struct {
// 	UserID      int
// 	Count       int
// 	LastRequest time.Time
// }

// type RateLimitMiddleWare struct {
// 	MiddlewareHanlder http.Handler
// 	Users             sync.Map
// 	MaxDuration       time.Duration
// 	MaxRequests       int
// }

// type ClientInfo struct {
// 	Count       int
// 	LastRequest time.Time
// 	sync.Mutex
// }

// func NewRateLimitMiddleWare(handler http.Handler) *RateLimitMiddleWare {
// 	return &RateLimitMiddleWare{handler, sync.Map{}, time.Duration(time.Minute * 1), 500}
// }

// func NewMiddleWare(handler http.Handler, service *service.AppService) *Middleware {
// 	return &Middleware{handler, service}
// }

// func NewLoginMiddleware(handler http.Handler, service *service.AppService) *LoginRegisterMiddleWare {
// 	return &LoginRegisterMiddleWare{handler, service}
// }

// func WriteJsonErrors(w http.ResponseWriter, errJson models.ErrorJson) {
// 	w.WriteHeader(errJson.Status)
// 	json.NewEncoder(w).Encode(errJson)
// }
