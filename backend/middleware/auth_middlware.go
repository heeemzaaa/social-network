package middleware

// import (
// 	"net/http"

// 	"social-network/backend/models"

// 	// handler "real-time-forum/backend/handler"
// )

// // could be returning a boolean but to see again
// func (m *Middleware) GetAuthUserEnsureAuth(r *http.Request) (*models.Session, *models.ErrorJson) {
// 	cookie, err := r.Cookie("session")
// 	if err != nil {
// 		return nil, &models.ErrorJson{Status: 401, Message: "ERROR!! Unauthorized Access"}
// 	}
// 	// check if the value of the cookie is correct and if not expired!!!
// 	session, errJson := m.service.GetSessionByTokenEnsureAuth(cookie.Value)
// 	if errJson != nil || session.IsExpired() {
// 		return nil, errJson
// 	}
// 	return session, nil
// }

// func (m *Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("content-Type", "application/json")
// 	_, err := m.GetAuthUserEnsureAuth(r)
// 	if err != nil {
// 		// handler.WriteJsonErrors(w, *err)
// 		return
// 	}
// 	m.MiddlewareHanlder.ServeHTTP(w, r)
// }
