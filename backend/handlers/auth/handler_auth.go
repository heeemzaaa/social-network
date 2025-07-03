package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"

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

	autPostPaths := []string{"/api/auth/login", "/api/auth/register", "/api/auth/logout"}

	switch method {
	case http.MethodPost:
		switch path {
		case "/api/auth/login":
			auth.login(w, r)
			return
		case "/api/auth/register":
			auth.register(w, r)
			return
		case "/api/auth/logout":
			auth.logout(w, r)
			return
		default:
			handlers.WriteJsonErrors(w, models.ErrorJson{Status: 404, Error: "Page not found."})
			return
		}
	case http.MethodGet:
		if slices.Contains(autPostPaths, path) {
			handlers.WriteJsonErrors(w, models.ErrorJson{Status: 405, Error: "ERROR!! Method Not Allowed!!"})
			return
		}
		if path != "/api/auth/islogged" {
			handlers.WriteJsonErrors(w, models.ErrorJson{Status: 404, Error: "Page not found."})
			return
		}
		auth.isLoggedIn(w, r)
	default:
		handlers.WriteJsonErrors(w, models.ErrorJson{Status: 405, Error: "ERROR!! Method Not Allowed!!"})
		return
	}
}

func (handler *AuthHandler) isLoggedIn(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside the is logged in handler")
}

func (handler *AuthHandler) login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside the login handler.")
	login := &models.Login{}
	// fmt.Println("login: ", login)
	err := json.NewDecoder(r.Body).Decode(&login)

	if (err == io.EOF || login == &models.Login{}) {
		handlers.WriteJsonErrors(w, models.ErrorJson{
			Status: 400,
			Message: models.Login{
				LoginField: "login field can't be empty",
				Password:   "password field can't be empty",
			},
		})
		return
	}

	if err != nil {
		fmt.Println("error decoding")
		handlers.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: fmt.Sprintf("%v 1", err)})
		return
	}

	// errJson := handler.service.Login(login)
	// if errJson != nil {
	// 	handlers.WriteJsonErrors(w, *errJson)
	// 	return
	// }

	handlers.WriteDataBack(w, "user logged in successfuly")
}

func (handler *AuthHandler) register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside the register handler.")
	user := &models.User{}
	// fmt.Println("login: ", login)
	err := json.NewDecoder(r.Body).Decode(&user)

	if (err == io.EOF || *user == models.User{}) {
		handlers.WriteJsonErrors(w, models.ErrorJson{
			Status: 400,
			Message: models.User{
				FirstName: "login field can't be empty",
				LastName:  "password field can't be empty",
				BirthDate: "login field can't be empty",
				Email:     "password field can't be empty",
				Password:  "password field can't be empty",
			},
		})
		return
	}
	if err != nil {

		fmt.Println("error decoding")
		handlers.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: fmt.Sprintf("%v 1", err)})
		return
	}

	// errJson := handler.service.Login(login)
	// if errJson != nil {
	// 	handlers.WriteJsonErrors(w, *errJson)
	// 	return
	// }

	handlers.WriteDataBack(w, user)
}

func (handler *AuthHandler) logout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside the logout handler")
}
