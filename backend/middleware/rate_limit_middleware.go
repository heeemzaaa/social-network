package middleware

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"social-network/backend/models"
)

func (rl *RateLimitMiddleWare) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		http.Error(w, "Invalid IP address", http.StatusInternalServerError)
		return
	}

	val, ok := rl.Users.Load(ip)
	if ok {
		clientInfo, ok := val.(*ClientInfo)
		if !ok {
			http.Error(w, "ERROR!! Internal server error", http.StatusInternalServerError)
			return
		}

		clientInfo.Lock()
		defer clientInfo.Unlock()

		if time.Since(clientInfo.LastRequest) > rl.MaxDuration {
			clientInfo.Count = 1
			clientInfo.LastRequest = time.Now()
		} else {
			if clientInfo.Count >= rl.MaxRequests {
				fmt.Println("", clientInfo.Count)
				WriteJsonErrors(w, models.ErrorJson{
					Status:  http.StatusTooManyRequests,
					Message: "ERROR!! Too many Requests",
				})
				return
			}
			clientInfo.Count++
		}
	} else {
		rl.Users.Store(ip, &ClientInfo{
			Count:       1,
			LastRequest: time.Now(),
		})
	}

	rl.MiddlewareHanlder.ServeHTTP(w, r)
}
