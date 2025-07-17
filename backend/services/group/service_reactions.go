package group

import "social-network/backend/models"

func (gService *GroupService) AddReaction(reaction *models.GroupReaction, reaction_type int) (*models.GroupReaction, *models.ErrorJson) {
	reaction_created, err := gService.gRepo.AddReaction(reaction, reaction_type)
	if err != nil {
		return nil, err
	}
	return reaction_created, nil
}

func (gService *GroupService) UpdateReaction(reaction *models.GroupReaction, reaction_type int) (*models.GroupReaction, *models.ErrorJson) {
	if reaction_type == 1 {
		reaction_created, err := gService.gRepo.UpdateReactionLike(reaction)
		if err != nil {
			return nil, err
		}
		return reaction_created, nil
	}
	return nil, nil
}

//

func (gService *GroupService) HanldeReaction(reaction *models.GroupReaction, reaction_type int) (*models.GroupReaction, *models.ErrorJson) {
	if errJson := gService.gRepo.GetGroupById(reaction.GroupId); errJson != nil {
		return nil, &models.ErrorJson{Status: errJson.Status, Message: errJson.Message, Error: errJson.Error}
	}
	// always check the membership and also the the group is a valid one
	if errMembership := gService.CheckMembership(reaction.GroupId, reaction.UserId); errMembership != nil {
		return nil, &models.ErrorJson{Status: errMembership.Status, Error: errMembership.Error, Message: errMembership.Message}
	}
	reactionERR := models.GroupReactionErr{}
	if !IsValidType(reaction.EntityType) {
		reactionERR.EntityType = "wrong entity_type!!"
	}

	if IsValidType(reaction.EntityType) {
		switch reaction.EntityType {
		case "post":
			_, exists, _ := gService.gRepo.GetItem("group_posts", "postID", reaction.EntityId)
			if !exists {
				reactionERR.EntityId = "entity_id not found"
			}
		case "comment":
			_, exists, _ := gService.gRepo.GetItem("group_comments", "commentID", reaction.EntityId)
			if !exists {
				reactionERR.EntityId = "entity_id not found"
			}
		}
	}

	if reactionERR != (models.GroupReactionErr{}) {
		return nil, &models.ErrorJson{Status: 400, Message: reactionERR}
	}

	reaction_existed, err := gService.gRepo.HanldeReaction(reaction)
	if err != nil {
		return reaction_existed, &models.ErrorJson{Status: err.Status, Message: err.Message}
	}
	if reaction_existed == nil {
		reaction, errJson := gService.AddReaction(reaction, reaction_type)
		if errJson != nil {
			return nil, errJson
		}
		return reaction, nil
	} else {
		reaction, errJson := gService.UpdateReaction(reaction_existed, reaction_type)
		if errJson != nil {
			return nil, errJson
		}
		return reaction, nil
	}
}

// there is a problem
// we need a query to get if there is a reaction for the comment
// each time we need to check the value stored in the database and then
// based on it we chose either to add -1
// like if the first time add 1 and if the second time
