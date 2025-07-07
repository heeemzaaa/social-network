package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"
	"time"

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
	islogged := &models.IsLoggedIn{}
	if r.Method != http.MethodGet {
		handlers.WriteJsonErrors(w, *models.NewErrorJson(405, "Method Not Allowed", nil))
		return
	}

	cookie, err := r.Cookie("session")
	if err != nil {
		islogged.IsLoggedIn = false
		handlers.WriteDataBack(w, islogged)
		return
	}

	islogged, errJson := handler.service.IsLoggedInUser(cookie.Value)

	if errJson != nil {
		islogged.IsLoggedIn = false
		handlers.WriteDataBack(w, islogged)
		return
	}
	islogged.IsLoggedIn = true
	handlers.WriteDataBack(w, islogged)
}

func (handler *AuthHandler) login(w http.ResponseWriter, r *http.Request) {
	login := &models.Login{}
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		if err == io.EOF {
			// case if the body sent if empty
			handlers.WriteJsonErrors(w, models.ErrorJson{
				Status: 400,
				Message: models.Login{
					LoginField: "field is required",
					Password:   "field is required",
				},
			})
			return
		}
		handlers.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: fmt.Sprintf("%v 1", err)})
		return
	}

	user, errJson := handler.service.Login(login)
	if errJson != nil {
		handlers.WriteJsonErrors(w, *errJson)
		return
	}

	session, errJSON := handler.service.CreateOrUpdateSession(user)
	if errJSON != nil {
		handlers.WriteJsonErrors(w, models.ErrorJson{
			Status:  errJSON.Status,
			Message: errJSON.Message,
		})
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "session",
		Value:   session.Token,
		Expires: session.ExpDate,
		Path:    "/",
	})

	handlers.WriteDataBack(w, "user logged in successfuly")
}

func (authHandler *AuthHandler) register(w http.ResponseWriter, r *http.Request) {
	user := models.NewUser()

	data := r.FormValue("data")
	err := json.Unmarshal([]byte(data), &user)
	if err != nil {
		fmt.Println("Error while decoding the the register request body: ", err)
		if err == io.EOF || *user == (models.User{}) {
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

	file, handler, err := r.FormFile("profile_img")
	if err != nil {
		// No file uploaded, set defaults for optional image
		user.ProfileImage = ""
		user.ProfileImgSize = 0
	} else {
		user.ProfileImage = handler.Filename
		user.ProfileImgSize = handler.Size
		defer file.Close()
	}

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

	handlers.WriteDataBack(w, "User registered seccussfully")
}

func (handler *AuthHandler) logout(w http.ResponseWriter, r *http.Request) {
	// delete from the database before
	cookie, _ := r.Cookie("session")
	session, errJson := handler.service.GetSessionByTokenEnsureAuth(cookie.Value)
	if errJson != nil {
		handlers.WriteJsonErrors(w, *models.NewErrorJson(errJson.Status, "", errJson.Message))
		return
	}
	if err := handler.service.Logout(session); err != nil {
		handlers.WriteJsonErrors(w, *models.NewErrorJson(err.Status, "", err.Message))
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
