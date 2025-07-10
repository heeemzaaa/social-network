package group

import (
	"strings"

	"social-network/backend/models"
	"social-network/backend/utils"
)

// the process of checking if a user is a member of a group of not will be hold at the service of every function

func (service *GroupService) GetGroupEvents(groupID, userID string, offset int64) ([]models.Event, *models.ErrorJson) {
	// check if the user is a member or not

	exists, errJson := service.IsMemberGroup(groupID, userID)
	if errJson != nil {
		return nil, &models.ErrorJson{Status: errJson.Status, Message: errJson.Message}
	}
	if !exists {
		return nil, &models.ErrorJson{Status: 403, Message: "ERROR!! Acces Forbidden!"}
	}

	events, errJson := service.grepo.GetGroupEvents(groupID, offset)
	if errJson != nil {
		return nil, &models.ErrorJson{Status: errJson.Status, Message: errJson.Message}
	}
	return events, nil
}

//  check if the values entered by the user are correct ( waaa tleee3  liya hadshii frassii mumiil)
// as always we need to check if the user is part of the group before adding an event

func (service *GroupService) AddGroupEvent(event *models.Event, groupID, userID string) (*models.Event, *models.ErrorJson) {
	exists, errJson := service.IsMemberGroup(groupID, userID)
	if errJson != nil {
		return nil, &models.ErrorJson{Status: errJson.Status, Message: errJson.Message}
	}
	if !exists {
		return nil, &models.ErrorJson{Status: 403, Message: "ERROR!! Acces Forbidden!"}
	}
	// here we'll be checking if the input is valid
	errValidation := models.ErrEventGroup{}
	trimmedTitle := strings.TrimSpace(event.Title)
	trimmedDesc := strings.TrimSpace(event.Description)
	if err := utils.ValidateTitle(trimmedTitle); err != nil {
		errValidation.Title = err.Error()
	}
	if err := utils.ValidateDesc(trimmedDesc); err != nil {
		errValidation.Description = err.Error()
	}
    // check the date of the event (but how ???)
	if event.EventDate.IsZero() {
		
	}

	if errValidation != (models.ErrEventGroup{}) {
		return nil, &models.ErrorJson{Status: 400, Message: errValidation}
	}

	return service.grepo.AddGroupEvent(event)
}
