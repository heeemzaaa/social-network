package group

import (
	"errors"

	"social-network/backend/models"
	"social-network/backend/utils"
)

func (service *GroupService) ValidateGroupTitle(title string) error {
	_, has_title, _ := service.gRepo.GetItem("groups", "title", title)
	if has_title {
		return errors.New("title already in use")
	}

	if err := utils.ValidateTitle(title); err != nil {
		return err
	}
	return nil
}

func (service *GroupService) CheckMembership(groupID, userID string) *models.ErrorJson {
	isMember, errJson := service.IsMemberGroup(groupID, userID)
	if errJson != nil {
		return &models.ErrorJson{Status: errJson.Status, Message: errJson.Message}
	}
	if !isMember {
		return &models.ErrorJson{Status: 403, Message: "ERROR!! Acces Forbidden!"}
	}
	return nil
}

func (service *GroupService) GroupExists(groupID string) *models.ErrorJson {
	if err := service.gRepo.GetGroupById(groupID); err != nil {
		return &models.ErrorJson{Status: err.Status, Error: err.Error, Message: err.Message}
	}
	return nil
}

func IsValidType(typeEntity string) bool {
	return typeEntity == "post" || typeEntity == "comment"
}
