package group

import "social-network/backend/models"

func (service *GroupService) GetEventDetails(eventId, userId, groupId string) (*models.Event, *models.ErrorJson) {
	event, errJson := service.gRepo.GetEventDetails(eventId, userId, groupId)
	if errJson != nil {
		return nil, &models.ErrorJson{Status: errJson.Status, Error: errJson.Error, Message: errJson.Message}
	}
	return event, nil
}

// SECTION DYAL GOING AND NOT GOING
func (service *GroupService) AddAction(userId, groupId, eventId string, action_chosen int) (*models.UserEventAction, *models.ErrorJson) {
	action, err := service.gRepo.AddAction(userId, groupId, eventId, action_chosen)
	if err != nil {
		return nil, &models.ErrorJson{Status: err.Status, Error: err.Error, Message: err.Message}
	}
	return action, nil
}

func (gService *GroupService) UpdateAction(userId, groupId, eventId string, action_chosen int) (*models.UserEventAction, *models.ErrorJson) {
	switch action_chosen {
	case 1:
		return gService.gRepo.UpdateToGoing(userId, groupId, eventId)
	case -1:
		return gService.gRepo.UpdateToNotGoing(userId, groupId, eventId)
	default:
		return nil, &models.ErrorJson{Status: 400, Error: "error!! wrong type of action only -1 and 1 are allowed!"}
	}
}

func (gService *GroupService) HandleActionChosen(userId, groupId, eventId string, action int) (*models.UserEventAction, *models.ErrorJson) {
	if errJson := gService.gRepo.GetGroupById(groupId); errJson != nil {
		return nil, &models.ErrorJson{Status: errJson.Status, Message: errJson.Message, Error: errJson.Error}
	}
	// always check the membership and also the the group is a valid one
	if errMembership := gService.CheckMembership(groupId, userId); errMembership != nil {
		return nil, &models.ErrorJson{Status: errMembership.Status, Error: errMembership.Error, Message: errMembership.Message}
	}

	// now let's check if the event where we're trying to add
	_, eventExists, _ := gService.gRepo.GetItem("group_events", "eventID", eventId)
	if !eventExists {
		return nil, &models.ErrorJson{
			Status: 404,
			Error:  "event not found!",
		}
	}

	if action != -1 && action != 1 {
		return nil, &models.ErrorJson{Status: 400, Error: "bad type of action"}
	}

	action_existed, errJson := gService.gRepo.HanldeAction(eventId, userId, userId)
	if errJson != nil {
		return nil, &models.ErrorJson{Status: errJson.Status, Error: errJson.Error, Message: errJson.Message}
	}
	// so there was a reaction and we need to edit it
	if action_existed == nil {
		action_created, errJson := gService.AddAction(userId, groupId, eventId, action)
		if errJson != nil {
			return nil, &models.ErrorJson{Status: errJson.Status, Error: errJson.Error, Message: errJson.Message}
		}
		return action_created, nil
	}

	action_created, errJson := gService.UpdateAction(userId, groupId, eventId, action)
	if errJson != nil {
		return nil, &models.ErrorJson{Status: errJson.Status, Error: errJson.Error, Message: errJson.Message}
	}
	return action_created, nil

	//  now let's check if there is a reaction or not and then
	// update it if necessary
}

// func (gService *GroupService) HanldeReaction(reaction *models.GroupReaction, reaction_type int) (*models.GroupReaction, *models.ErrorJson) {
// 	if errJson := gService.gRepo.GetGroupById(reaction.GroupId); errJson != nil {
// 		return nil, &models.ErrorJson{Status: errJson.Status, Message: errJson.Message, Error: errJson.Error}
// 	}
// 	// always check the membership and also the the group is a valid one
// 	if errMembership := gService.CheckMembership(reaction.GroupId, reaction.UserId); errMembership != nil {
// 		return nil, &models.ErrorJson{Status: errMembership.Status, Error: errMembership.Error, Message: errMembership.Message}
// 	}
// 	reactionERR := models.GroupReactionErr{}
// 	if !IsValidType(reaction.EntityType) {
// 		reactionERR.EntityType = "wrong entity_type!!"
// 	}

// 	if IsValidType(reaction.EntityType) {
// 		switch reaction.EntityType {
// 		case "post":
// 			_, exists, _ := gService.gRepo.GetItem("group_posts", "postID", reaction.EntityId)
// 			if !exists {
// 				reactionERR.EntityId = "entity_id not found"
// 			}
// 		case "comment":
// 			_, exists, _ := gService.gRepo.GetItem("group_comments", "commentID", reaction.EntityId)
// 			if !exists {
// 				reactionERR.EntityId = "entity_id not found"
// 			}
// 		}
// 	}

// 	if reactionERR != (models.GroupReactionErr{}) {
// 		return nil, &models.ErrorJson{Status: 400, Message: reactionERR}
// 	}

// 	reaction_existed, err := gService.gRepo.HanldeReaction(reaction)
// 	if err != nil {
// 		return reaction_existed, &models.ErrorJson{Status: err.Status, Message: err.Message}
// 	}
// 	if reaction_existed == nil {
// 		reaction, errJson := gService.AddReaction(reaction, reaction_type)
// 		if errJson != nil {
// 			return nil, errJson
// 		}
// 		return reaction, nil
// 	} else {
// 		reaction, errJson := gService.UpdateReaction(reaction_existed, reaction_type)
// 		if errJson != nil {
// 			return nil, errJson
// 		}
// 		return reaction, nil
// 	}
// }
