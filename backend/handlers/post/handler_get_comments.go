package handlers

import (
	"fmt"
	"net/http"

	"social-network/backend/models"
	"social-network/backend/utils"
)

func (h *PostHandler) GetComments(w http.ResponseWriter, r *http.Request) {
	fmt.Println("**********************************")
	postID, err := utils.GetUUIDFromPath(r, "id")
	fmt.Println(postID.String())
	if err != nil {
		fmt.Println(err)
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: "invalid path"})
		return
	}

	comments, errComments := h.service.GetComments(postID.String())
	if errComments != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errComments.Status, Error: errComments.Error})
		return
	}

	utils.WriteDataBack(w, comments)
}
