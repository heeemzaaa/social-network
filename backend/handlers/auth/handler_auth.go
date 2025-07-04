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
			if path == "/api/auth/islogged" {
				handlers.WriteJsonErrors(w, models.ErrorJson{Status: 405, Error: "Method Not Allowed."})
			} else {
				handlers.WriteJsonErrors(w, models.ErrorJson{Status: 404, Error: "Page not found."})
			}
			return
		}
	case http.MethodGet:
		if slices.Contains(autPostPaths, path) {
			handlers.WriteJsonErrors(w, models.ErrorJson{Status: 405, Error: "Method Not Allowed."})
			return
		}
		if path != "/api/auth/islogged" {
			fmt.Println("heeeee")
			handlers.WriteJsonErrors(w, models.ErrorJson{Status: 404, Error: "Page not found."})
			return
		}
		auth.isLoggedIn(w, r)
	default:
		handlers.WriteJsonErrors(w, models.ErrorJson{Status: 405, Error: "Method Not Allowed."})
		return
	}
}

func (handler *AuthHandler) isLoggedIn(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside the is logged in handler")

	handlers.WriteDataBack(w, "user is logged in")
}

func (handler *AuthHandler) login(w http.ResponseWriter, r *http.Request) {
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

func (authHandler *AuthHandler) register(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("inside the register handler.")
	user := &models.User{}

	data := r.FormValue("data")
	err := json.Unmarshal([]byte(data), &user)
	if err != nil {
		fmt.Println("Error while decoding the the register request body: ", err)
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
	}

	file, handler, _ := r.FormFile("profile_img")
	user.ProfileImage = handler.Filename
	user.ProfileImgSize = handler.Size

	defer file.Close()

	errJson := authHandler.service.Register(user, file)
	if errJson != nil {
		handlers.WriteJsonErrors(w, *errJson)
		return
	}

	userData, errJson := authHandler.service.GetUser(&models.Login{LoginField: user.Email})
	fmt.Printf("userData: %v\n", userData)

	if errJson != nil {
		handlers.WriteJsonErrors(w, *errJson)
		return
	}
	// fmt.Printf("userData: %v\n", userData)
	// var login = models.Login{LoginField: user.Nickname}
	session, err_ := authHandler.service.SetUserSession(userData)
	if err_ != nil {
		handlers.WriteJsonErrors(w, *err_)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "session",
		Value:   session.Token,
		Expires: session.ExpDate,
		Path:    "/",
	})
	// we don't need to write back the data for the repsonse ( sentitive data ;)
	handlers.WriteDataBack(w, "User registered seccussfully")
}

func (handler *AuthHandler) logout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside the logout handler")
}
