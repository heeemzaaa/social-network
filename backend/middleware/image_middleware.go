package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"social-network/backend/models"
	"social-network/backend/utils"
)

func (img *ImageMiddleware) AuthImageMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, err := GetUserIDFromContext(r.Context())
		if err != nil {
			utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)})
			return
		}

		creatorID := r.URL.Query().Get("creator")
		postID := r.URL.Query().Get("postID")
		groupID := r.URL.Query().Get("groupID")

		if !strings.HasPrefix(r.URL.Path, "/static/uploads/") {
			utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: "Invalid path!"})
			return
		}

		trimmedPath := strings.TrimPrefix(r.URL.Path, "/static/uploads/")

		switch {
		case strings.HasPrefix(trimmedPath, "posts"),
			strings.HasPrefix(trimmedPath, "postsComments"):
			access, err := img.service.AccessToPostImage(userID.String(), creatorID, postID)
			if err != nil {
				utils.WriteJsonErrors(w, models.ErrorJson{Status: err.Status, Error: err.Error})
				return
			}
			if !access {
				utils.WriteJsonErrors(w, models.ErrorJson{Status: 403, Error: "You don't have access to see this image!"})
				return
			}

		case strings.HasPrefix(trimmedPath, "groups"):
			access, err := img.service.AccessToGroupImage(groupID, userID.String())
			if err != nil {
				utils.WriteJsonErrors(w, models.ErrorJson{Status: err.Status, Error: err.Error})
				return
			}
			if !access {
				utils.WriteJsonErrors(w, models.ErrorJson{Status: 403, Error: "You don't have access to see this image!"})
				return
			}

		case strings.HasPrefix(trimmedPath, "avatars"):
			// nothing to check here , just add it to avoid using it in default 
		default:
			utils.WriteJsonErrors(w, models.ErrorJson{Status: 404, Error: "File not found"})
			return
		}

		http.ServeFile(w, r, "static/uploads/"+trimmedPath)
	})
}
