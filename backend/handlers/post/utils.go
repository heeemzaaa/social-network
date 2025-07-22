package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"social-network/backend/utils"
)

func GetPath(r *http.Request) (string, string, error) {
	splittedPath := strings.Split(strings.Trim(r.URL.Path, "/"), "/")

	if len(splittedPath) < 3 {
		return "", "", fmt.Errorf("the endpoint is not correct")
	}

	postID, err := utils.GetUUIDFromPath(r, "id")
	if err != nil {
		return "", "", fmt.Errorf("%v", err)
	}

	path := ""
	if len(splittedPath) > 4 {
		path = splittedPath[4]
	}
	return postID.String(), path, nil
}
