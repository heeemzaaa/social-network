package group

import (
	"social-network/backend/models"
)

func (gService *GroupService) JoinGroup(group *models.Group, userId string) *models.ErrorJson {
	if errJson := gService.grepo.JoinGroup(group, userId); errJson != nil {
		return errJson
	}
	return nil
}

func (gService *GroupService) GetGroupInfo(groupId string) (*models.Group, *models.ErrorJson) {
	group, errjson := gService.grepo.GetGroupDetails(groupId)
	if errjson != nil {
		return nil, &models.ErrorJson{Status: errjson.Status, Message: errjson.Message}
	}
	return group, nil
}
