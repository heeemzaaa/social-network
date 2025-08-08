package middleware

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"social-network/backend/models"
	auth "social-network/backend/services/auth"
	"social-network/backend/services/images"
)

type Middleware struct {
	MiddlewareHanlder http.Handler
	service           *auth.AuthService
}

type LoginRegisterMiddleWare struct {
	MiddlewareHanlder http.Handler
	service           *auth.AuthService
}

type UserInfo struct {
	UserID      string
	Count       int
	LastRequest time.Time
}

type ImageMiddleware struct {
	service *images.ServiceImages
}

type RateLimitMiddleWare struct {
	MiddlewareHanlder http.Handler
	Users             sync.Map
	MaxDuration       time.Duration
	MaxRequests       int
}

type ClientInfo struct {
	Count       int
	LastRequest time.Time
	sync.Mutex
}

func NewRateLimitMiddleWare(handler http.Handler) *RateLimitMiddleWare {
	return &RateLimitMiddleWare{handler, sync.Map{}, time.Duration(time.Minute * 1), 500}
}

func NewMiddleWare(handler http.Handler, service *auth.AuthService) *Middleware {
	return &Middleware{handler, service}
}

func NewLoginMiddleware(handler http.Handler, service *auth.AuthService) *LoginRegisterMiddleWare {
	return &LoginRegisterMiddleWare{handler, service}
}

func NewImagesMiddleware(handler http.Handler, service *images.ServiceImages) *ImageMiddleware {
	return &ImageMiddleware{service: service}
}

func WriteJsonErrors(w http.ResponseWriter, errJson models.ErrorJson) {
	w.WriteHeader(errJson.Status)
	json.NewEncoder(w).Encode(errJson)
}
