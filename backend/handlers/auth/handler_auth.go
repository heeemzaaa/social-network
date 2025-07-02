package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"social-network/backend/handlers"
	"social-network/backend/models"
	"social-network/backend/services/auth"
)

type AuthHandler struct {
	service *auth.AuthService
}

func NewAuthHandler(service *auth.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (auth *AuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	method := r.Method
	path := r.URL.Path
	if method == http.MethodPost {
		switch path {
		case "/auth/login":
			auth.login(w, r)
		case "/auth/register":
			auth.register(w,r)
		case "/auth/logout":
			auth.lo
		default:
				handlers.WriteJsonErrors(w, models.ErrorJson{Status: 404, Message: "Page not found."})
			return
		} 
	} else if method == http.MethodGet {
		if path != "/auth/islogged" {
			handlers.WriteJsonErrors(w, models.ErrorJson{Status: 404, Message: "Page not found."})
			return
		}
	} else {
		handlers.WriteJsonErrors(w, models.ErrorJson{Status: 405, Message: "ERROR!! Method Not Allowed!!"})
		return
	}
}

func (handler *AuthHandler) isLoggedIn(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside the is logged in handler")

}

func (handler *AuthHandler) login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside the login handler.")

	login := &models.Login{}
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		if err == io.EOF {
			handlers.WriteJsonErrors(w, models.ErrorJson{
				Status: 400,
				Message: models.Login{
					LoginField: "empty login field!!",
					Password:   "empty password field!!",
				},
			})
			return
		}
		handlers.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: fmt.Sprintf("%v 1", err)})
		return
	}

	errJson := handler.service.Login(login)
	if errJson != nil {
		handlers.WriteJsonErrors(w, *errJson)
		return
	}

	handlers.WriteDataBack(w, "user logged in successfuly")
}

func (handler *AuthHandler) register(w http.ResponseWriter, r *http.Request) {
}

func (handler *AuthHandler) logout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside the logout handler")
}