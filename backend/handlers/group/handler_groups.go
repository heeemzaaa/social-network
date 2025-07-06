package group

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"social-network/backend/models"
	gservice "social-network/backend/services/group"
	"social-network/backend/utils"

	"github.com/google/uuid"
)

type GroupHanlder struct {
	gservice *gservice.GroupService
}

func NewGroupHandler(service *gservice.GroupService) *GroupHanlder {
	return &GroupHanlder{gservice: service}
}

func (Ghandler *GroupHanlder) GetGroups(w http.ResponseWriter, r *http.Request) {
	userIDVal := r.Context().Value("userID")
	userID, ok := userIDVal.(uuid.UUID)
	if !ok {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Message: "Incorrect type of userID value!"})
		return
	}
	filter := r.URL.Query().Get("filter")
	offset, errOffset := strconv.ParseInt(r.URL.Query().Get("offset"), 10, 64)
	if errOffset != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: "Incorrect Offset Format!"})
		return
	}
	if !utils.IsValidFilter(filter) {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: "Incorrect filter by field!!"})
		return
	}
	groups, errJson := Ghandler.gservice.GetGroups(filter, offset, userID.String())
	if errJson != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errJson.Status, Message: errJson.Message})
		return
	}
	if err := json.NewEncoder(w).Encode(groups); err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)})
		return
	}
}

func (Ghandler *GroupHanlder) CreateGroup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hnaaaa")
	// userIDVal := r.Context().Value("userID")
	// userID, ok := userIDVal.(uuid.UUID)
	// if !ok {
	// 	utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Message: "Incorrect type of userID value!"})
	// 	return
	// }
	var group_to_create *models.Group

	data := r.FormValue("data")
	if err := json.Unmarshal([]byte(data), &group_to_create); err != nil {
		if err == io.EOF {
			utils.WriteJsonErrors(w, models.ErrorJson{
				Status: 400,
				Message: models.ErrGroup{
					Title:       "empty title field!",
					Description: "empty description field!",
				},
			})
			return
		}

		utils.WriteJsonErrors(w, models.ErrorJson{
			Status:  400,
			Message: "an error occured while trying to decode the json!",
		})
		return
	}

	groupCreatorId, errToUUID := uuid.FromBytes([]byte("25f9ba66-80aa-42c1-960d-c22e7757d91a"))
	if errToUUID != nil {
		fmt.Println("err", errToUUID)
	}
	// handle the image encoding in the phase that comes before the adding process
	errUploadImg := HanldeUploadImage(w, r, "group")
	if errUploadImg != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errUploadImg.Status, Message: errUploadImg.Message})
		return
	}

	group_to_create.GroupCreatorId = groupCreatorId
	group, errJson := Ghandler.gservice.AddGroup(group_to_create)
	if errJson != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errJson.Status, Message: errJson.Message})
		return
	}
	utils.WriteDataBack(w, group)
}

func HanldeUploadImage(w http.ResponseWriter, r *http.Request, fileName string) *models.ErrorJson {
	file, header, err := r.FormFile(fileName)
	if err != nil {
		if err == http.ErrMissingFile || err == io.EOF {
			return &models.ErrorJson{Status: 400, Message: "Error!! Missing file"}
		}
	}
	defer file.Close()

	mimeType := header.Header.Get("Content-Type")
	if !utils.IsValidImageType(mimeType) {
		return &models.ErrorJson{Status: 400, Message: "Error!! Only PNG, JPEG and GIF images are allowed"}
	}
	// fmt.Println("header", header)
	
	return nil
}

func (Ghandler *GroupHanlder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appplication/json")
	fmt.Println("inside the serveHttp func")
	switch r.Method {
	case http.MethodGet:
		Ghandler.GetGroups(w, r)
		return
	case http.MethodPost:
		Ghandler.CreateGroup(w, r)
		return
	default:
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Message: "ERROR!! Method Not Allowed!"})
		return
	}
}
