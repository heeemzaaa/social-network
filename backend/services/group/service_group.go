package group

import (
	"strings"

	"social-network/backend/models"
)





func (gService *GroupService) JoinGroup(group *models.Group, userId string) *models.ErrorJson {
	if strings.TrimSpace(group.GroupId) == "" {
		return &models.ErrorJson{Status: 400, Message: models.ErrJoinGroup{
			GroupId: "missing group_id field!!",
		}}
	}
	if errJson := gService.gRepo.GetGroupById(group.GroupId); errJson != nil {
		return &models.ErrorJson{Status: errJson.Status, Message: errJson.Message, Error: errJson.Error}
	}
	if errJson := gService.gRepo.JoinGroup(group, userId); errJson != nil {
		return &models.ErrorJson{Status: errJson.Status, Message: errJson.Message, Error: errJson.Error}
	}
	return nil
}

func (gService *GroupService) GetGroupInfo(groupId string) (*models.Group, *models.ErrorJson) {
	if errJson := gService.gRepo.GetGroupById(groupId); errJson != nil {
		return nil, &models.ErrorJson{Status: errJson.Status, Message: errJson.Message, Error: errJson.Error}
	}

	group, errjson := gService.gRepo.GetGroupDetails(groupId)
	if errjson != nil {
		return nil, &models.ErrorJson{Status: errjson.Status, Message: errjson.Message, Error: errjson.Error}
	}
	return group, nil
}
