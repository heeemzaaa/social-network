package auth

import (
	"net/http"
	"time"

	"social-network/backend/models"
	"social-network/backend/services/auth"
	"social-network/backend/utils"
)

type Logout AuthHandler

func NewLogoutHandler(service *auth.AuthService) *Logout {
	return &Logout{service: service}
}

// let's implement the logout
func (logout *Logout) Logout(w http.ResponseWriter, r *http.Request) {
	// delete from the database before
	cookie, _ := r.Cookie("session")
	session, errJson := logout.service.GetSessionByTokenEnsureAuth(cookie.Value)
	if errJson != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errJson.Status, Message: errJson.Message})
		return
	}
	if err := logout.service.DeleteSession(session); err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{})
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "session",
		Value:   "",
		MaxAge:  -1,
		Expires: time.Unix(0, 0),
		Path:    "/",
	})

	w.WriteHeader(http.StatusNoContent)
}

func (logout *Logout) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch {
	case r.Method != http.MethodPost && r.URL.Path == "/api/auth/logout":
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Message: "ERROR!! Method Not Allowed!"})
		return
	case r.Method == http.MethodPost && r.URL.Path == "/api/auth/logout":
		logout.Logout(w, r)
		return
	default:
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 404, Message: "ERROR!! Page Not Found!"})
		return

	}
}
