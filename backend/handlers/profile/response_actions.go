package profile

import (
	"encoding/json"
	"net/http"

	"social-network/backend/models"
	ps "social-network/backend/services/profile"
	"social-network/backend/utils"
)

type ResponseHandler struct {
	service *ps.ProfileService
}

func NewResponseHandler(service *ps.ProfileService) *ResponseHandler {
	return &ResponseHandler{service: service}
}

// POST api/profile/id/response/accepted
func (rh *ResponseHandler) AcceptedRequest(w http.ResponseWriter, r *http.Request, profileID string) {
	type RequestBody struct {
		Requestor string `json:"requestor_id"`
	}

	var request RequestBody
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: "Invalid data !"})
		return
	}

	errRequest := rh.service.AcceptedRequest(profileID, request.Requestor)
	if errRequest != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errRequest.Status, Message: errRequest.Message})
		return
	}
	utils.WriteDataBack(w, "done")
}

// POST api/profile/id/response/rejected
func (rh *ResponseHandler) RejectedRequest(w http.ResponseWriter, r *http.Request, profileID string) {
	type RequestBody struct {
		Requestor string `json:"requestor_id"`
	}

	var request RequestBody
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: "Invalid data !"})
		return
	}

	errRequest := rh.service.RejectedRequest(profileID, request.Requestor)
	if errRequest != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errRequest.Status, Message: errRequest.Message})
		return
	}
	utils.WriteDataBack(w, "done")
}

func (rh *ResponseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Message: "Method not allowed !"})
		return
	}

	profileID, path, err := GetPath(r)
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: "Invalid path"})
	}

	switch path {
	case "accepted":
		rh.AcceptedRequest(w, r, profileID)
	case "rejected":
		rh.RejectedRequest(w, r, profileID)
	default:
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 404, Message: "Page not found !"})
	}
}
