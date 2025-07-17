package group

import (
	"strings"

	"social-network/backend/models"
	"social-network/backend/utils"
)

// the process of checking if a user is a member of a group of not will be hold at the service of every function

func (gService *GroupService) GetGroupEvents(groupID, userID string, offset int64) ([]models.Event, *models.ErrorJson) {
	if errJson := gService.gRepo.GetGroupById(groupID); errJson != nil {
		return nil, &models.ErrorJson{Status: errJson.Status, Message: errJson.Message, Error: errJson.Error}
	}
	// always check the membership and also the the group is a valid one
	if errMembership := gService.CheckMembership(groupID, userID); errMembership != nil {
		return nil, &models.ErrorJson{Status: errMembership.Status, Error: errMembership.Error, Message: errMembership.Message}
	}
	// if the one trying to access is a member wlla laa

	events, errJson := gService.gRepo.GetGroupEvents(groupID, offset)
	if errJson != nil {
		return nil, &models.ErrorJson{Status: errJson.Status, Message: errJson.Message, Error: errJson.Error}
	}
	return events, nil
}

//  check if the values entered by the user are correct ( waaa tleee3  liya hadshii frassii mumiil)
// as always we need to check if the user is part of the group before adding an event

func (gService *GroupService) AddGroupEvent(event *models.Event) (*models.Event, *models.ErrorJson) {
	if errJson := gService.gRepo.GetGroupById(event.GroupId); errJson != nil {
		return nil, &models.ErrorJson{Status: errJson.Status, Message: errJson.Message, Error: errJson.Error}
	}
	// always check the membership and also the the group is a valid one
	if errMembership := gService.CheckMembership(event.GroupId, event.EventCreatorId); errMembership != nil {
		return nil, &models.ErrorJson{Status: errMembership.Status, Error: errMembership.Error, Message: errMembership.Message}
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
	if err := utils.ValidateDateEvent(event.EventDate); err != nil {
		errValidation.EventDate = err.Error()
	}

	if errValidation != (models.ErrEventGroup{}) {
		return nil, &models.ErrorJson{Status: 400, Message: errValidation}
	}
	event, errJson := gService.gRepo.AddGroupEvent(event)
	if errJson != nil {
		return nil, &models.ErrorJson{Status: errJson.Status, Message: errJson.Message, Error: errJson.Error}
	}
	return event, nil
}
