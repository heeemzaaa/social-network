package group

import (
	"social-network/backend/models"
)

func (gService *GroupService) GetEventDetails(eventId, userId, groupId string) (*models.Event, *models.ErrorJson) {
	if errJson := gService.gRepo.GetGroupById(groupId); errJson != nil {
		return nil, &models.ErrorJson{Status: errJson.Status, Message: errJson.Message, Error: errJson.Error}
	}
	// always check the membership and also the the group is a valid one
	if errMembership := gService.CheckMembership(groupId, userId); errMembership != nil {
		return nil, &models.ErrorJson{Status: errMembership.Status, Error: errMembership.Error, Message: errMembership.Message}
	}
	event, errJson := gService.gRepo.GetEventDetails(eventId, groupId, userId)
	if errJson != nil {
		return nil, &models.ErrorJson{Status: errJson.Status, Error: errJson.Error, Message: errJson.Message}
	}
	return event, nil
}

// SECTION DYAL GOING AND NOT GOING
// HERE we'll only check the membership once 

func (gService *GroupService) HandleActionChosen(actionChosen *models.UserEventAction) (*models.UserEventAction, *models.ErrorJson) {
	if errJson := gService.gRepo.GetGroupById(actionChosen.GroupId); errJson != nil {
		return nil, &models.ErrorJson{Status: errJson.Status, Message: errJson.Message, Error: errJson.Error}
	}
	// always check the membership and also the the group is a valid one
	if errMembership := gService.CheckMembership(actionChosen.GroupId, actionChosen.UserId); errMembership != nil {
		return nil, &models.ErrorJson{Status: errMembership.Status, Error: errMembership.Error, Message: errMembership.Message}
	}

	// now let's check if the event where we're trying to add
	_, eventExists, _ := gService.gRepo.GetItem("group_events", "eventID", actionChosen.EventId)
	if !eventExists {
		return nil, &models.ErrorJson{
			Status: 404,
			Error:  "event not found!",
		}
	}

	if actionChosen.Action != -1 && actionChosen.Action != 1 {
		return nil, &models.ErrorJson{Status: 400, Error: "bad type of action"}
	}

	action_existed, errJson := gService.gRepo.HanldeAction(actionChosen)
	if errJson != nil {
		return nil, &models.ErrorJson{Status: errJson.Status, Error: errJson.Error, Message: errJson.Message}
	}

	// so there was a reaction and we need to edit it
	if action_existed == nil {
		action_created, errJson := gService.AddAction(actionChosen)
		if errJson != nil {
			return nil, &models.ErrorJson{Status: errJson.Status, Error: errJson.Error, Message: errJson.Message}
		}
		return action_created, nil
	}
                    
	action_created, errJson := gService.UpdateAction(actionChosen)
	if errJson != nil {
		return nil, &models.ErrorJson{Status: errJson.Status, Error: errJson.Error, Message: errJson.Message}
	}
	return action_created, nil

	//  now let's check if there is a reaction or not and then
	// update it if necessary
}

func (gService *GroupService) AddAction(actionChosen *models.UserEventAction) (*models.UserEventAction, *models.ErrorJson) {
	action, err := gService.gRepo.AddAction(actionChosen)
	if err != nil {
		return nil, &models.ErrorJson{Status: err.Status, Error: err.Error, Message: err.Message}
	}
	return action, nil
}

func (gService *GroupService) UpdateAction(actionChosen *models.UserEventAction) (*models.UserEventAction, *models.ErrorJson) {
	switch actionChosen.Action {
	case 1:
		return gService.gRepo.UpdateToGoing(actionChosen)
	case -1:
		return gService.gRepo.UpdateToNotGoing(actionChosen)
	default:
		return nil, &models.ErrorJson{Status: 400, Error: "error!! wrong type of action only -1 and 1 are allowed!"}
	}
}
