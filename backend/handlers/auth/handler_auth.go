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

func (authHandler *AuthHandler) register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside the register handler.")
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
	r.ParseMultipartForm()
	fmt.Printf("user: %v\n", user)

	// file, handler, err := r.FormFile("profile_img")
	// if err != nil {
	// 	fmt.Println("Error Retrieving the File")
	// 	fmt.Println(err)
	// 	return
	// }

	// defer file.Close()
	// fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	// fmt.Printf("File Size: %+v\n", handler.Size)
	// fmt.Printf("MIME Header: %+v\n", handler.Header)
	// fmt.Printf("handler: %v\n", handler)

	// fmt.Printf("user data: %v\n", user)
	// err := json.NewDecoder(r.Body).Decode(&user)
	// if err != nil {
	// 	fmt.Println("error while decoding the the register request body: ", err)
	// 	if (err == io.EOF || *user == models.User{}) {
	// 		handlers.WriteJsonErrors(w, models.ErrorJson{
	// 			Status: 400,
	// 			Message: models.User{
	// 				FirstName: "login field can't be empty",
	// 				LastName:  "password field can't be empty",
	// 				BirthDate: "login field can't be empty",
	// 				Email:     "password field can't be empty",
	// 				Password:  "password field can't be empty",
	// 			},
	// 		})
	// 		return
	// 	}
	// }

	// errJson := handler.service.Register(user)
	// if errJson != nil {
	// 	handlers.WriteJsonErrors(w, *errJson)
	// 	return
	// }

	handlers.WriteDataBack(w, "still under dev")
}

func (handler *AuthHandler) logout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside the logout handler")
}
